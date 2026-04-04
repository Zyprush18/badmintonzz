package routes

import (
	"github.com/Zyprush18/badmintonzz/internal/schedule/application"
	"github.com/Zyprush18/badmintonzz/internal/schedule/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/schedule/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterSchedule(r *gin.RouterGroup, db *sqlx.DB)  {
	repo_schedule := infrastructure.NewRepoSchedule(db)
	svc_schedule := application.NewApplicationSchedules(repo_schedule)
	hndl_schedule := interfaces.NewHandlerSchedule(svc_schedule)


	r.GET("/", hndl_schedule.Index)
	r.GET("/:id", hndl_schedule.Show)

}