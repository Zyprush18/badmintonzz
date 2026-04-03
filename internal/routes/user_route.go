package routes

import (
	"github.com/Zyprush18/badmintonzz/internal/users/application"
	"github.com/Zyprush18/badmintonzz/internal/users/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/users/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRouteUsers(users *gin.RouterGroup, db *sqlx.DB)  {
	user_repo := infrastructure.NewRepoUsers(db)
	user_svc := application.NewServiceUsers(user_repo)
	user_hndl := interfaces.NewHandlerUsers(user_svc)

	users.GET("/", user_hndl.Index)
	users.POST("/", user_hndl.Create)
	users.GET("/:id", user_hndl.Show)
	users.PUT("/:id", user_hndl.Update)
	users.DELETE("/:id", user_hndl.Delete)

}