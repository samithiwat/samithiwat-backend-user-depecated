package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleID uint `json:"role_id" gorm:"index:,unique"`
}
