package domain

import (
	"database/sql"
	"time"
)


type Schedules struct {
	ID int
	Date string
	Time time.Time
	Duration int
	Service_id int
	Created_at sql.NullTime
	Updated_at sql.NullTime
}