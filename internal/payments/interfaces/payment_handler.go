package interfaces

import (
	"errors"
	"log"
	"net/http"

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

func (s *handlePayments) CallBacksMidtrans(c *gin.Context) {
	

	c.JSON(http.StatusOK, gin.H{
		"message": "masuk ke callback",
	})
}