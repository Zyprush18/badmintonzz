package interfaces

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Zyprush18/badmintonzz/internal/booking/application"
	"github.com/Zyprush18/badmintonzz/internal/booking/interfaces/request"
	"github.com/Zyprush18/badmintonzz/internal/shared/cntx"
	"github.com/Zyprush18/badmintonzz/internal/shared/errs"
	"github.com/Zyprush18/badmintonzz/internal/shared/validation"
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

	role := c.GetString("role")
	user_id := c.GetInt("user_id")

	data, err := s.svc.QueriesBooking().GetBookings(ctx, user_id, role)
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




func (s *handlerBooking) Show(c *gin.Context) {
	ctx, cancel := cntx.TimeoutShortContext(c.Request.Context())
	defer cancel()

	id := c.Param("id")
	bookingID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid booking ID"})
		return
	}

	role := c.GetString("role")
	user_id := c.GetInt("user_id")

	data, err := s.svc.QueriesBooking().GetBooking(ctx, bookingID, user_id, role)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, errs.NotFoundRow) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Booking not found"})
			return
		}

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

func (s *handlerBooking) Create(c *gin.Context) {
	ctx, cancel := cntx.TimeoutShortContext(c.Request.Context())
	defer cancel()

	booking := new(request.BookingRequest)
	if err := c.ShouldBindJSON(booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid booking data"})
		return
	}

	if err:= validation.ValidateCheckFields(ctx, booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errs.InvalidValidation,
			"error": err.Error(),
		})
		return
	}


	// role := c.GetString("role") -> aktifkan jika sudah terapkan login
	role := "admin"
	user_id := c.GetInt("user_id")
	

	token, redirectURL, err := s.svc.CommandsBooking().CreateBooking(ctx, booking, user_id, role)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, errs.ContextTimeout) {
			c.JSON(http.StatusGatewayTimeout, gin.H{"message": errs.RequestTimeout})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": errs.ServerError})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Booking created successfully",
		"token": token,
		"redirect_url": redirectURL,
	})
}