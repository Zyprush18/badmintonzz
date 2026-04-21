package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRoute(r *gin.RouterGroup, db *sqlx.DB)  {
	users := r.Group("/users")
	RegisterRouteUsers(users, db)

	services := r.Group("/services")
	RegisterRouteServices(services, db)

	booking := r.Group("/bookings")
	RegisterRouteBooking(booking, db)

	payments := r.Group("/payments")
	RegisterRoutePayments(payments, db)
}