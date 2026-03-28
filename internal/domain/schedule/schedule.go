package schedule

import "time"


type Schedules struct {
	ID int
	Date string
	Time time.Time
	Duration int
	Service_id int
}