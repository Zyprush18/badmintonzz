package routes

import (
	"github.com/Zyprush18/badmintonzz/internal/middleware"
	"github.com/Zyprush18/badmintonzz/internal/services/application"
	"github.com/Zyprush18/badmintonzz/internal/services/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/services/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRouteServices(services *gin.RouterGroup, db *sqlx.DB)  {
	svc_repo := infrastructure.NewRepoServices(db)
	svc_app := application.NewApplicationServices(svc_repo)
	svc_handl := interfaces.NewHandlerService(svc_app)

	services.GET("/",svc_handl.Index)
	services.GET("/:id",svc_handl.Show)

	
	services.Use(middleware.CheckAuthToken())
	services.POST("/", middleware.CheckRole("admin"), svc_handl.Create)
	services.PUT("/:id", middleware.CheckRole("admin"),svc_handl.Update)
	services.DELETE("/:id", middleware.CheckRole("admin"),svc_handl.Delete)
}