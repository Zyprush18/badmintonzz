package interfaces

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Zyprush18/badmintonzz/internal/shared/validation"
	"github.com/Zyprush18/badmintonzz/internal/users/application"
	"github.com/Zyprush18/badmintonzz/internal/users/domain"
	"github.com/Zyprush18/badmintonzz/internal/users/interfaces/request"
	"github.com/gin-gonic/gin"
)

type HandlerUsers struct {
	svc application.ServicesUsers
}


func NewHandlerUsers(s application.ServicesUsers) *HandlerUsers {
	return &HandlerUsers{svc: s}
}


func (s *HandlerUsers) Index(c *gin.Context) {
	users, err := s.svc.QueriesUsers().GetUsers(c.Request.Context())
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

	user, err := s.svc.QueriesUsers().GetUser(c.Request.Context(), userID)
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


func (s *HandlerUsers) Create(c *gin.Context) {
	user := new(request.UserRequest)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": domain.InvalidRequest})
		return
	}

	if err:= validation.ValidateCheckFields(c.Request.Context(), user);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": domain.InvalidValidation,
			"error": err.Error(),
		})
		return
	}

	if err := s.svc.CommandsUsers().CreateUsers(c.Request.Context(), user); err != nil {
		log.Println(err.Error())
		if domain.CheckDuplicate(err) {
			c.JSON(http.StatusConflict, gin.H{"message": domain.DuplicateUser})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": domain.ServerError})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}