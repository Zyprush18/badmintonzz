package interfaces

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Zyprush18/badmintonzz/internal/schedule/application"
	"github.com/Zyprush18/badmintonzz/internal/schedule/domain"
	"github.com/Zyprush18/badmintonzz/internal/shared/cntx"
	"github.com/Zyprush18/badmintonzz/internal/shared/errs"
	"github.com/gin-gonic/gin"
)

type HandlerSchedule struct {
	svc application.ScheduleApp
}

func NewHandlerSchedule(s application.ScheduleApp) HandlerSchedule  {
	return HandlerSchedule{
		svc: s,
	}
}


func (s *HandlerSchedule) Index(c *gin.Context) {
	ctx, cancel := cntx.TimeoutLongContext(c.Request.Context())
	defer cancel()

	data, err := s.svc.QueriesSchedules().GetSchedules(ctx)
	if err != nil {
		log.Println(err.Error())

		if errors.Is(err, errs.ContextTimeout) {
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


func (s *HandlerSchedule) Show(c *gin.Context) {
	ctx, cancel := cntx.TimeoutShortContext(c.Request.Context())
	defer cancel()

	id_params := c.Param("id")
	if id_params == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": domain.InvalidId})
		return
	}


	id, err := strconv.Atoi(id_params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": domain.InvalidId})
		return
	}


	data, err := s.svc.QueriesSchedules().GetSchedule(ctx, id)
	if err != nil {
		log.Println(err.Error())

		if errors.Is(err, errs.NotFoundRow) {
			c.JSON(http.StatusNotFound, gin.H{"message": domain.NotFoundRow})
			return
		}


		if errors.Is(err, errs.ContextTimeout) {
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