package service

import (
	"github.com/samithiwat/samithiwat-backend-user/src/model"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
)

type UserService struct {
	repository UserRepository
}

type UserRepository interface {
	FindAll(*proto.PaginationMetadata, *[]*model.User) error
	FindOne(uint, *model.User) error
	FindMulti([]uint32, *[]*model.User) error
	Create(*model.User) error
	Update(uint, *model.User) error
	Delete(uint, *model.User) error
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}
