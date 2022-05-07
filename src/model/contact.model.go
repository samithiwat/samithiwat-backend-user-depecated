package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	ContactID uint `json:"contact_id" gorm:"index:,unique"`
}
