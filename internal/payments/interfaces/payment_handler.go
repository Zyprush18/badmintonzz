package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



type handlePayments struct {
	
}

func NewHandlePayment() *handlePayments {
	return &handlePayments{}
}

func (s *handlePayments) CallBacksMidtrans(c *gin.Context) {
	

	c.JSON(http.StatusOK, gin.H{
		"message": "masuk ke callback",
	})
}