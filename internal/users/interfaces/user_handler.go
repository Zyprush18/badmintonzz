package interfaces

import (
	"log"
	"net/http"

	"github.com/Zyprush18/badmintonzz/internal/users/application/queries"
	"github.com/gin-gonic/gin"
)

type HandlerUsers struct {
	svc queries.ServiceUsers
}


func NewHandlerUsers(s queries.ServiceUsers) *HandlerUsers {
	return &HandlerUsers{svc: s}
}


func (s *HandlerUsers) Index(r *gin.Context) {
	users, err := s.svc.GetUsers()
	if err != nil {
		log.Println(err.Error())
		r.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    users,

	})
}