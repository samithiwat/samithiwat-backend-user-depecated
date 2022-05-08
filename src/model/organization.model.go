package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	OrganizationID uint    `json:"organization_id" gorm:"index:,unique"`
	Members        []*User `json:"users" gorm:"many2many:user_organization;"`
}
