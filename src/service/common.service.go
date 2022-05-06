package service

import (
	"github.com/samithiwat/samithiwat-backend-user/src/model"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"gorm.io/gorm"
)

func RawToDtoUser(user *model.User) *proto.User {
	return &proto.User{
		Id:        uint32(user.ID),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		ImageUrl:  user.ImageUrl,
	}
}

func DtoToRawUser(user *proto.User) *model.User {
	return &model.User{
		Model:     gorm.Model{ID: uint(user.Id)},
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		ImageUrl:  user.ImageUrl,
	}
}
