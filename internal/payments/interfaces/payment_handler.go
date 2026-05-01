package interfaces

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Zyprush18/badmintonzz/internal/payments/application"
	"github.com/Zyprush18/badmintonzz/internal/shared/cntx"
	"github.com/Zyprush18/badmintonzz/internal/shared/errs"
	"github.com/gin-gonic/gin"
)



type handlePayments struct {
	svc application.ApplicationPayment
}

func NewHandlePayment(s application.ApplicationPayment) *handlePayments {
	return &handlePayments{
		svc: s,
	}
}

func (s *handlePayments) Index(c *gin.Context) {
	ctx, cancel := cntx.TimeoutLongContext(c.Request.Context())
	defer cancel()

	data, err := s.svc.QueriesPayment().GetAll(ctx)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, errs.ContextTimeout) {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"message": errs.RequestTimeout,
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errs.ServerError,
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message":"Success",
		"data": data,
	})
}


func (s *handlePayments) Show(c *gin.Context)  {
	ctx, cancel := cntx.TimeoutLongContext(c.Request.Context())
	defer cancel()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid payment ID",
		})
		return
	}

	data, err := s.svc.QueriesPayment().GetPaymentById(ctx, id)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, errs.ContextTimeout) {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"message": errs.RequestTimeout,
			})
			return
		}

		if errors.Is(err, errs.NotFoundRow) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Payment not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errs.ServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": data,
	})
}

func (s *handlePayments) CallBacksMidtransAddPayment(c *gin.Context) {
	ctx, cancel := cntx.TimeoutLongContext(c.Request.Context())
	defer cancel()

	order_id := c.Query("order_id")
	transaction_status := c.Query("transaction_status")

	if err:= s.svc.CommandsPayment().UpdatePayment(ctx, order_id, transaction_status);err != nil {
		log.Println(err.Error())
		if errors.Is(err, errs.ContextTimeout) {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"message": errs.RequestTimeout,
			})
			return
		}

		if errors.Is(err, errs.NotFoundRow) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Payment not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errs.ServerError,
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "Success Added Booking",
	})
}