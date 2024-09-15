package models

import (
	"time"
)

type Cab struct {
	ID            uint   `gorm:"primaryKey"`
	LicensePlate  string `gorm:"unique"`
	CurrentCityID uint
	State         string `gorm:"default:IDLE"`
	LastIdleTime  time.Time
}

type City struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

type CabHistory struct {
	ID        uint `gorm:"primaryKey"`
	CabID     uint
	CityID    uint
	State     string
	Timestamp time.Time
}
