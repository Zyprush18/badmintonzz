package interfaces

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Zyprush18/badmintonzz/internal/users/application/queries"
	"github.com/Zyprush18/badmintonzz/internal/users/domain"
	"github.com/gin-gonic/gin"
)

type HandlerUsers struct {
	svc queries.ServiceUsers
}


func NewHandlerUsers(s queries.ServiceUsers) *HandlerUsers {
	return &HandlerUsers{svc: s}
}


func (s *HandlerUsers) Index(c *gin.Context) {
	users, err := s.svc.GetUsers(c.Request.Context())
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": domain.ServerError})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    users,

	})
}


func (s *HandlerUsers) Show(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": domain.InvalidID})
		return
	}

	user, err := s.svc.GetUser(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, domain.NotFoundRow) {
			c.JSON(http.StatusNotFound, gin.H{"message": domain.NotFoundUser})
			return
		}

		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": domain.ServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    user,
	})
}