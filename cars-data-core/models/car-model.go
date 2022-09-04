package models

import (
	"gorm.io/gorm"
)

type CarModel struct {
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

func (CarModel) TableName() string {
	return "car"
}

func MigrateModels(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&CarModel{})
	return db
}
