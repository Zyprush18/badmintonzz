package http

import (
	"log"

	// "github.com/Zyprush18/badmintonzz/internal/domain/users"
	"github.com/Zyprush18/badmintonzz/internal/infrastucture/db"
	"github.com/gin-gonic/gin"
)


func RegisteRoute(r *gin.RouterGroup)  {
	_, err := db.Connect_DB()
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Failed to connect to database")
	}

	// user_repo := users.NewRepoUsers(database)


}