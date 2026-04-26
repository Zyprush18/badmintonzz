package config

import (
	"os"
	"strconv"

	"github.com/Zyprush18/badmintonzz/internal/booking/interfaces/request"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransCfg interface {
	SnapRequest(requestData *request.BookingRequest) (*snap.Response, error)
}

type midtransStruct struct {
	snap snap.Client
	core coreapi.Client
}

func NewMidtrans() MidtransCfg {
	env := os.Getenv("MIDTRANS_SERVER_KEY")
	s := snap.Client{}
	c := coreapi.Client{}
	s.New(env, midtrans.Sandbox)
	c.New(env, midtrans.Sandbox)
	return &midtransStruct{
		snap: s,
		core: c,
	}
}


func (m *midtransStruct) SnapRequest(requestData *request.BookingRequest) (*snap.Response, error) {
	amount := int64(requestData.Price) * int64(requestData.Hour)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: requestData.Order_Id,
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
				ID: strconv.Itoa(requestData.Service_id),
				Name: requestData.Name_svc,
				Price: int64(requestData.Price),
				Qty: int32(requestData.Hour),
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