package interfaces

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Zyprush18/badmintonzz/internal/shared/cntx"
	"github.com/Zyprush18/badmintonzz/internal/shared/errs"
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


func (s *HandlerUsers) AuthRegister(c *gin.Context) {
	ctx, cancel := cntx.TimeoutLongContext(c.Request.Context())
	defer cancel()

	req := new(request.UserAuthRegisterRequest)

	if err:= c.ShouldBindJSON(req);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errs.InvalidRequest,
		})

		return
	}

	if err:= validation.ValidateCheckFields(ctx, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errs.InvalidValidation,
			"error": err.Error(),
		})
		return
	}


	if err := s.svc.CommandsUsers().AuthRegister(ctx, req); err != nil {
		log.Println(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}

		if domain.CheckDuplicate(err) {
			c.JSON(http.StatusConflict, gin.H{"message": domain.DuplicateUser})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})

}


func (s *HandlerUsers) AuthLogin(c *gin.Context) {
	ctx, cancel := cntx.TimeoutLongContext(c.Request.Context())
	defer cancel()

	req := new(request.UserAuthLoginRequest)

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errs.InvalidRequest,
		})
		return
	}

	if err:= validation.ValidateCheckFields(ctx, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errs.InvalidValidation,
			"error": err.Error(),
		})
		return
	}

	token, err := s.svc.CommandsUsers().AuthLogin(ctx, req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}

		if err.Error() == "invalid credentials" || errors.Is(err, errs.NotFoundRow) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "email or password invalid"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

func (s *HandlerUsers) GetProfile(c *gin.Context) {
	id := c.GetString("user_id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": domain.InvalidID})
		return
	}

	ctx, cancel := cntx.TimeoutShortContext(c.Request.Context())
	defer cancel()

	user, err := s.svc.QueriesUsers().GetUser(ctx, userID)
	if err != nil {
		log.Println(err.Error())

		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}

		if errors.Is(err, errs.NotFoundRow) {
			c.JSON(http.StatusNotFound, gin.H{"message": domain.NotFoundUser})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    user,
	})
}


func (s *HandlerUsers) Index(c *gin.Context) {
	ctx, cancel := cntx.TimeoutLongContext(c.Request.Context())
	defer cancel()

	users, err := s.svc.QueriesUsers().GetUsers(ctx)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}


		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
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

	ctx, cancel := cntx.TimeoutShortContext(c.Request.Context())
	defer cancel()

	user, err := s.svc.QueriesUsers().GetUser(ctx, userID)
	if err != nil {
		log.Println(err.Error())

		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}

		if errors.Is(err, errs.NotFoundRow) {
			c.JSON(http.StatusNotFound, gin.H{"message": domain.NotFoundUser})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
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
		c.JSON(http.StatusBadRequest, gin.H{"message": errs.InvalidRequest})
		return
	}


	ctx, cancel := cntx.TimeoutLongContext(c.Request.Context())
	defer cancel()


	if err:= validation.ValidateCheckFields(ctx, user);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errs.InvalidValidation,
			"error": err.Error(),
		})
		return
	}


	if err := s.svc.CommandsUsers().CreateUsers(ctx, user); err != nil {
		log.Println(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}

		if domain.CheckDuplicate(err) {
			c.JSON(http.StatusConflict, gin.H{"message": domain.DuplicateUser})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}


func (s *HandlerUsers) Update(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": domain.InvalidID})
		return
	}

	user := new(request.UserRequest)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errs.InvalidRequest})
		return
	}


	ctx, cancel := cntx.TimeoutLongContext(c.Request.Context())
	defer cancel()

	if err:= validation.ValidateCheckFields(ctx, user);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errs.InvalidValidation,
			"error": err.Error(),
		})
		return
	}

	if err := s.svc.CommandsUsers().UpdateUsers(ctx, userID, user); err != nil {
		log.Println(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}

		if errors.Is(err, errs.NotFoundRow) {
			c.JSON(http.StatusNotFound, gin.H{"message": domain.NotFoundUser})
			return
		}

		if domain.CheckDuplicate(err) {
			c.JSON(http.StatusConflict, gin.H{"message": domain.DuplicateUser})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}


func (s *HandlerUsers) Delete(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": domain.InvalidID})
		return
	}

	ctx, cancel := cntx.TimeoutShortContext(c.Request.Context())
	defer cancel()

	if err := s.svc.CommandsUsers().DeleteUsers(ctx, userID); err != nil {
		log.Println(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}

		if errors.Is(err, errs.NotFoundRow) {
			c.JSON(http.StatusNotFound, gin.H{"message": domain.NotFoundUser})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}