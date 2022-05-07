package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	TeamID uint `json:"team_id" gorm:"index:,unique"`
}
