package domain

import (
	"time"
)

type ScheduleServices struct {
	IdServices int `db:"service_id"`
	Name string `db:"name"`
	Price float64 `db:"price"`
	CreatedAt_Svc time.Time `db:"created_at_svc"`
	UpdatedAt_Svc time.Time `db:"updated_at_svc"`
	IdSchedule int `db:"schedule_id"`
	Date string `db:"date" json:"date"`
	Time string `db:"time" json:"time"`
	Duration int `db:"duration" json:"duration"`
	CreatedAt_Scdl time.Time `db:"created_at_scdl"`
	UpdatedAt_Scdl time.Time `db:"updated_at_scdl"`
}