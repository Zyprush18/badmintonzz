package routes

import (
	"github.com/Zyprush18/badmintonzz/internal/config"
	"github.com/Zyprush18/badmintonzz/internal/payments/application"
	"github.com/Zyprush18/badmintonzz/internal/payments/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/payments/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutePayments(r *gin.RouterGroup, db *sqlx.DB, m config.MidtransCfg) {
	repo_payment := infrastructure.NewRepositoryPayment(db)
	svc_payment := application.NewApplicationPayment(repo_payment, m)
	hndle_payment := interfaces.NewHandlePayment(svc_payment)


	r.GET("/", hndle_payment.Index)
	r.GET("/:id", hndle_payment.Show)
	r.GET("/callback", hndle_payment.CallBacksMidtransAddPayment)
}