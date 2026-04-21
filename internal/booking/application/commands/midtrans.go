package commands

import (
	"crypto/rand"
	"os"
	"strconv"

	"github.com/Zyprush18/badmintonzz/internal/booking/interfaces/request"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)


type midtransStruct struct {
	snap snap.Client
	core coreapi.Client
	requestData *request.BookingRequest
}

func NewMidtrans(req *request.BookingRequest) *midtransStruct {
	env := os.Getenv("MIDTRANS_SERVER_KEY")
	s := snap.Client{}
	c := coreapi.Client{}
	s.New(env, midtrans.Sandbox)
	c.New(env, midtrans.Sandbox)
	return &midtransStruct{
		snap: s,
		core: c,
		requestData: req,
	}
}


func (m *midtransStruct) SnapRequest() (*snap.Response, error) {
	amount := int64(m.requestData.Amount) * int64(m.requestData.Hour)

	order_id := "badmintonzz-" + strconv.Itoa(m.requestData.Service_id) + "-" + rand.Text()

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: order_id,
			GrossAmt: amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		Callbacks: &snap.Callbacks{
			Finish: os.Getenv("APP_HOST") + "/payments/callback",
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID: strconv.Itoa(m.requestData.Service_id),
				Name: m.requestData.Name_svc,
				Price: int64(m.requestData.Amount),
				Qty: int32(m.requestData.Hour),
			},
		},
	}


	response, err := m.snap.CreateTransaction(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (m *midtransStruct) GetData(orderID string) (*coreapi.TransactionStatusResponse, error) {
	response, err := m.core.CheckTransaction(orderID)
	if err != nil {
		return nil, err
	}
	
	return response, nil
}