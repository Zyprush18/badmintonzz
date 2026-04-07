package interfaces

import (
	"errors"
	"log"
	"net/http"

	"github.com/Zyprush18/badmintonzz/internal/booking/application"
	"github.com/Zyprush18/badmintonzz/internal/shared/cntx"
	"github.com/Zyprush18/badmintonzz/internal/shared/errs"
	"github.com/gin-gonic/gin"
)

type handlerBooking struct {
	svc application.BookingApplication
}

func NewHandlerBooking(s application.BookingApplication) *handlerBooking {
	return &handlerBooking{svc: s}
}

func (s *handlerBooking) Index(c *gin.Context)  {
	ctx, cancel := cntx.TimeoutLongContext(c.Request.Context())
	defer cancel()

	data, err := s.svc.QueriesBooking().GetBookings(ctx)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, errs.ContextTimeout) {
			c.JSON(http.StatusGatewayTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": data,
	})

}