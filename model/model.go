package model

import "gorm.io/gorm"

var DB *gorm.DB

type Farm struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null"`
	Ponds []Pond `json:"ponds"`
}

type Pond struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	FarmID uint   `json:"farm_id"`
}

type Telemetry struct {
	gorm.Model
	Ip       string
	Method   string
	Endpoint string
	Status   int
	Latency  int64
}
