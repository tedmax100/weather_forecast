package model

import "time"

type ForecastQueryDao struct {
	StartTime time.Time
	EndTime   time.Time
	Location  string
}
