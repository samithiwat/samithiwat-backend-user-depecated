package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	ImageUrl  string `json:"image_url"`
}

type UserPagination struct {
	Items *[]*User
	Meta  PaginationMetadata
}
