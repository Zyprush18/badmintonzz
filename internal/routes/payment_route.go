package routes

import (
	"github.com/Zyprush18/badmintonzz/internal/payments/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutePayments(r *gin.RouterGroup, db *sqlx.DB) {
	hndle_payment := interfaces.NewHandlePayment()

	r.POST("/callback", hndle_payment.CallBacksMidtrans)
}