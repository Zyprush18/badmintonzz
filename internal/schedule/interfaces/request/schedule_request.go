package request


type ScheduleRequest struct {
	Date string `json:"date" validate:"required"`
	Time string `json:"time" validate:"required"`
	Duration int `json:"duration" validate:"required,number"`
	Service_id int `json:"service_id" validate:"required,number"`
}