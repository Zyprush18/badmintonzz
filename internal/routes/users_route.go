package routes

import (
	"github.com/Zyprush18/badmintonzz/internal/users/application"
	"github.com/Zyprush18/badmintonzz/internal/users/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/users/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRouteUser(r *gin.RouterGroup, db *sqlx.DB)  {
	user_repo := infrastructure.NewRepoUsers(db)
	user_svc := application.NewServiceUsers(user_repo)
	user_hndl := interfaces.NewHandlerUsers(user_svc)

	r.GET("/", user_hndl.Index)
	r.POST("/", user_hndl.Create)
	r.GET("/:id", user_hndl.Show)
	r.PUT("/:id", user_hndl.Update)
	r.DELETE("/:id", user_hndl.Delete)

}