package routes

import (
	"github.com/Zyprush18/badmintonzz/internal/middleware"
	"github.com/Zyprush18/badmintonzz/internal/users/application"
	"github.com/Zyprush18/badmintonzz/internal/users/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/users/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRouteUsers(r *gin.RouterGroup, db *sqlx.DB) {
	user_repo := infrastructure.NewRepoUsers(db)
	user_svc := application.NewServiceUsers(user_repo)
	user_hndl := interfaces.NewHandlerUsers(user_svc)

	r.GET("/profile", middleware.CheckAuthToken(),middleware.CheckRole("user", "admin"), user_hndl.GetProfile)
	auth := r.Group("/auth")
	{
		auth.POST("/register", user_hndl.AuthRegister)
		auth.POST("/login", user_hndl.AuthLogin)
	}


	users := r.Group("/users")
	users.Use(middleware.CheckAuthToken(),middleware.CheckRole("admin"))
	{
		users.GET("/", user_hndl.Index)
		users.POST("/",user_hndl.Create)
		users.GET("/:id", user_hndl.Show)
		users.PUT("/:id", user_hndl.Update)
		users.DELETE("/:id", user_hndl.Delete)

	}
}
