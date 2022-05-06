package test

import (
	"github.com/samithiwat/samithiwat-backend-user/src/model"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
)

func RawToDtoUser(user *model.User) *proto.User {
	return &proto.User{
		Id:        uint32(user.ID),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		ImageUrl:  user.ImageUrl,
	}
}
