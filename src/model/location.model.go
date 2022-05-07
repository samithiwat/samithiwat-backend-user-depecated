package model

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	LocationID uint `json:"location_id" gorm:"index:,unique"`
}
