package interfaces

import (
	"log"
	"net/http"

	"github.com/Zyprush18/badmintonzz/internal/services/application"
	"github.com/Zyprush18/badmintonzz/internal/shared/errs"
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
	data, err := a.app.QueriesServices().GetServices()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": data,
	})

}