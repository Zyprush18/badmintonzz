package routes

import (
	"github.com/Zyprush18/badmintonzz/internal/users/application/queries"
	"github.com/Zyprush18/badmintonzz/internal/users/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/users/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRouteUser(r *gin.RouterGroup, db *sqlx.DB)  {
	user_repo := infrastructure.NewRepoUsers(db)
	user_svc := queries.NewServiceUsers(user_repo)
	user_hndl := interfaces.NewHandlerUsers(user_svc)

	r.GET("/", user_hndl.Index)

}