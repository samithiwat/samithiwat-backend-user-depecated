package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	OrganizationID uint `json:"organization_id" gorm:"index:,unique"`
}
