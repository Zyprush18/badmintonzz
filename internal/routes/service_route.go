package routes

import (
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

	services.GET("/", svc_handl.Index)
}