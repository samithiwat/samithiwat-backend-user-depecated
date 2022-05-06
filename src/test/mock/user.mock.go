package mock

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-user/src/model"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"gorm.io/gorm"
)

var User1 model.User
var User2 model.User
var User3 model.User
var User4 model.User
var Users []*model.User
var CreateUserReqMock proto.CreateUserRequest
var UpdateUserReqMock proto.UpdateUserRequest

type UserMockRepo struct {
}

func (*UserMockRepo) FindAll(meta *proto.PaginationMetadata, users *[]*model.User) error {
	meta.CurrentPage = 1
	meta.TotalPage = 1
	meta.ItemCount = 4
	meta.TotalItem = 4
	meta.ItemsPerPage = 10

	*users = Users
	return nil
}

func (*UserMockRepo) FindOne(_ uint, user *model.User) error {
	*user = User1
	return nil
}

func (*UserMockRepo) FindMulti(_ []uint32, users *[]*model.User) error {
	*users = Users
	return nil
}

func (*UserMockRepo) Create(user *model.User) error {
	*user = User1
	return nil
}

func (*UserMockRepo) Update(_ uint, user *model.User) error {
	*user = User1
	return nil
}

func (*UserMockRepo) Delete(_ uint, user *model.User) error {
	*user = User1
	return nil
}

type UserMockErrRepo struct {
}

func (*UserMockErrRepo) FindAll(meta *proto.PaginationMetadata, users *[]*model.User) error {
	return nil
}

func (*UserMockErrRepo) FindOne(_ uint, user *model.User) error {
	return errors.New("Not found user")
}

func (*UserMockErrRepo) FindMulti(_ []uint32, users *[]*model.User) error {
	return nil
}

func (*UserMockErrRepo) Create(user *model.User) error {
	return nil
}

func (*UserMockErrRepo) Update(_ uint, user *model.User) error {
	return errors.New("Not found user")
}

func (*UserMockErrRepo) Delete(_ uint, user *model.User) error {
	return errors.New("Not found user")
}

func InitializeMockUser() (err error) {
	User1 = model.User{
		Model:     gorm.Model{ID: 1},
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	User2 = model.User{
		Model:     gorm.Model{ID: 2},
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	User3 = model.User{
		Model:     gorm.Model{ID: 3},
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	User4 = model.User{
		Model:     gorm.Model{ID: 4},
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	CreateUserReqMock = proto.CreateUserRequest{
		User: &proto.User{
			Firstname: faker.FirstName(),
			Lastname:  faker.LastName(),
			ImageUrl:  faker.URL(),
		},
	}
	if err != nil {
		panic("Error occur while mocking data")
	}

	UpdateUserReqMock = proto.UpdateUserRequest{
		User: &proto.User{
			Id:        uint32(User1.ID),
			Firstname: faker.FirstName(),
			Lastname:  faker.LastName(),
			ImageUrl:  faker.URL(),
		},
	}
	if err != nil {
		panic("Error occur while mocking data")
	}

	Users = append(Users, &User1, &User2, &User3, &User4)

	return nil
}
