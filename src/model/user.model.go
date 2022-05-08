package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname     string          `json:"firstname"`
	Lastname      string          `json:"lastname"`
	ImageUrl      string          `json:"image_url"`
	Organizations []*Organization `json:"organizations" gorm:"many2many:user_organization;"`
	Teams         []*Team         `json:"teams" gorm:"many2many:user_team;"`
	Roles         []*Role         `json:"roles" gorm:"many2many:user_role;"`
	Location      Location        `json:"location"`
	Contact       Contact         `json:"contact"`
}

type UserPagination struct {
	Items *[]*User
	Meta  PaginationMetadata
}
