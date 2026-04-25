package commands

import (
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
	amount := int64(m.requestData.Price) * int64(m.requestData.Hour)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: m.requestData.Order_Id,
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
				Price: int64(m.requestData.Price),
				Qty: int32(m.requestData.Hour),
				Category: "Booking",
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