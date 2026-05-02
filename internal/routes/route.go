package routes

import (
	"github.com/Zyprush18/badmintonzz/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRoute(r *gin.RouterGroup, db *sqlx.DB)  {
	cfg_midtrans := config.NewMidtrans()

	RegisterRouteUsers(r, db)

	services := r.Group("/services")
	RegisterRouteServices(services, db)

	booking := r.Group("/bookings")
	RegisterRouteBooking(booking, db, cfg_midtrans)

	payments := r.Group("/payments")
	RegisterRoutePayments(payments, db, cfg_midtrans)
}