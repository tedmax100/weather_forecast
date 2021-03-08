package model

import "time"

type ForecastRequest struct {
	CurrentTime time.Time
	Location    string
}

func (req ForecastRequest) GetNearestLastPeriod() (time.Time, time.Time) {
	hourNow := req.CurrentTime.Hour()
	var startTime, endTime time.Time
	quotient := hourNow / 12
	if quotient == 0 {
		startTime = time.Date(req.CurrentTime.Year(), req.CurrentTime.Month(), req.CurrentTime.Day(), 6, 0, 0, 0, req.CurrentTime.Location())
		endTime = startTime.Add(12 * time.Hour)
		return startTime, endTime
	}
	startTime = time.Date(req.CurrentTime.Year(), req.CurrentTime.Month(), req.CurrentTime.Day(), 18, 0, 0, 0, req.CurrentTime.Location())
	endTime = startTime.Add(12 * time.Hour)
	return startTime, endTime
}
