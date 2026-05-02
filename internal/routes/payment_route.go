package routes

import (
	"github.com/Zyprush18/badmintonzz/internal/config"
	"github.com/Zyprush18/badmintonzz/internal/middleware"
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

	r.GET("/callback", hndle_payment.CallBacksMidtransAddPayment)

	r.Use(middleware.CheckAuthToken())
	r.GET("/", middleware.CheckRole("admin"),hndle_payment.Index)
	r.GET("/:id", middleware.CheckRole("admin"),hndle_payment.Show)
}