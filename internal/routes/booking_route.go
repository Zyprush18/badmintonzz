package routes

import (
	"github.com/Zyprush18/badmintonzz/internal/booking/application"
	"github.com/Zyprush18/badmintonzz/internal/booking/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/booking/interfaces"
	"github.com/Zyprush18/badmintonzz/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)


func RegisterRouteBooking(r *gin.RouterGroup, db *sqlx.DB)  {
	repo_booking := infrastructure.NewInfrastructureBooking(db)
	svc_booking := application.NewApplicationBooking(repo_booking)
	handl_booking := interfaces.NewHandlerBooking(svc_booking)

	r.Use(middleware.CheckRole("admin"))

	r.GET("/", handl_booking.Index)
	r.GET("/:id", handl_booking.Show)
}