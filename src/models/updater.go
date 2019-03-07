package models

import (
	"time"
)

type (
	Updater struct {
		ID      uint `json:"id" gorm:"primary_key"`
		GDSID   uint
		StartAt time.Time `gorm:"default:NOW()"`
		StopAt  *time.Time
		Status  bool
	}
	Updaters []Updater
)
