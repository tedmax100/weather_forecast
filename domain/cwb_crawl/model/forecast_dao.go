package model

import (
	"time"
)

type ForcasstDao struct {
	LocationId     uint `gorm:"primaryKey"`
	ElementId      uint `gorm:"primaryKey"`
	ParameterName  string
	ParameterUnit  *string
	ParameterValue *int
	StartTime      time.Time `gorm:"primaryKey"`
	EndTime        time.Time
}
