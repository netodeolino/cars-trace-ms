package models

import (
	"gorm.io/gorm"
)

type CarEntity struct {
	gorm.Model
	CarID     string
	Latitude  float64
	Longitude float64
	Operator  string
	Stopped   bool
	LineId    string
}

type Tabler interface {
	TableName() string
}

func (CarEntity) TableName() string {
	return "car"
}

func MigrateModels(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&CarEntity{})
	return db
}
