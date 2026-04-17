package commands

import (
	"crypto/rand"
	"os"
	"strconv"

	"github.com/Zyprush18/badmintonzz/internal/booking/interfaces/request"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)


type midtransStruct struct {
	env string
	requestData *request.BookingRequest
}

func NewMidtrans(req *request.BookingRequest) *midtransStruct {
	return &midtransStruct{
		env: os.Getenv("MIDTRANS_SERVER_KEY"),
		requestData: req,
	}
}


func (m *midtransStruct) SnapRequest() (*snap.Response, error) {
	s := snap.Client{}
	s.New(m.env, midtrans.Sandbox)


	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: "badmintonzz-" + strconv.Itoa(m.requestData.Service_id) + "-" + rand.Text(),
			GrossAmt: int64(m.requestData.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		Callbacks: &snap.Callbacks{
			Finish: "https://yourdomain.com/finish",
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


	response, err := s.CreateTransaction(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}