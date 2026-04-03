package interfaces

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Zyprush18/badmintonzz/internal/services/application"
	"github.com/Zyprush18/badmintonzz/internal/services/domain"
	"github.com/Zyprush18/badmintonzz/internal/services/interfaces/request"
	"github.com/Zyprush18/badmintonzz/internal/shared/errs"
	"github.com/Zyprush18/badmintonzz/internal/shared/validation"
	"github.com/gin-gonic/gin"
)

type HandlerServices struct {
	app application.ApplicationServices
}

func NewHandlerService(a application.ApplicationServices) *HandlerServices {
	return &HandlerServices{
		app: a,
	}
}

func (a *HandlerServices) Index(c *gin.Context)  {
	data, err := a.app.QueriesServices().GetServices(c.Request.Context())
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
		"message": "success",
		"data": data,
	})

}


func (a *HandlerServices) Show(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errs.InvalidRequest})
		return
	}

	service, err := a.app.QueriesServices().GetServiceByID(c.Request.Context(),serviceID)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}

		if errors.Is(err, errs.NotFoundRow) {
			c.JSON(http.StatusNotFound, gin.H{"message": domain.NotFoundService})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": service,
	})
}


func (a *HandlerServices) Create(c *gin.Context) {
	service := new(request.Services)
	if err := c.ShouldBindJSON(service); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": errs.InvalidRequest})
		return
	}

	if err:= validation.ValidateCheckFields(c.Request.Context(), service);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errs.InvalidValidation,
			"errors": err.Error(),
		})

		return
	}

	if err := a.app.CommandsServices().CreateService(c.Request.Context(), service); err != nil {
		log.Println(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success create new service",
	})
}


func (a *HandlerServices) Update(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errs.InvalidRequest})
		return
	}

	service := new(request.Services)
	if err := c.ShouldBindJSON(service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errs.InvalidRequest})
		return
	}

	if err:= validation.ValidateCheckFields(c.Request.Context(), service);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errs.InvalidValidation,
			"errors": err.Error(),
		})
		return
	}

	if err := a.app.CommandsServices().UpdateService(c.Request.Context(), serviceID, service); err != nil {
		log.Println(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}

		if errors.Is(err, errs.NotFoundRow) {
			c.JSON(http.StatusNotFound, gin.H{"message": domain.NotFoundService})
			return
		}


		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success update service",
	})
}


func (a *HandlerServices) Delete(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errs.InvalidRequest})
		return
	}

	if err := a.app.CommandsServices().DeleteService(c.Request.Context(), serviceID); err != nil {
		log.Println(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}
		
		if errors.Is(err, errs.NotFoundRow) {
			c.JSON(http.StatusNotFound, gin.H{"message": domain.NotFoundService})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success delete service",
	})
}