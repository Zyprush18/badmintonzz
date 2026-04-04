package response

import (
	"time"

	"github.com/Zyprush18/badmintonzz/internal/services/domain"
)


type Schedules struct {
	ID int `json:"id"`
	Date string `json:"date"`
	Time string `json:"time"`
	Duration int `json:"duration"`
	Service_id int `json:"service_id"`
	Services domain.Services `json:"services"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}