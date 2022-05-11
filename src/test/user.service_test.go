package test

import (
	"fmt"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"github.com/samithiwat/samithiwat-backend-user/src/service"
	"github.com/samithiwat/samithiwat-backend-user/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindAllUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)

	var result []*proto.User
	for _, usr := range mock.Users {
		result = append(result, RawToDtoUser(usr))
	}

	var errors []string

	want := &proto.UserPaginationResponse{
		Data: &proto.UserPagination{
			Items: result,
			Meta: &proto.PaginationMetadata{
				ItemsPerPage: 10,
				ItemCount:    4,
				TotalItem:    4,
				CurrentPage:  1,
				TotalPage:    1,
			},
		},
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	usrService := service.NewUserService(&mock.UserMockRepo{})
	usrRes, err := usrService.FindAll(mock.Context{}, &proto.FindAllUserRequest{Limit: 10, Page: 1})

	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, usrRes, fmt.Sprintf("Want %v but got %v", want, usrRes))
}

func TestFindOneUser(t *testing.T) {
	mock.InitializeMockUser()

	var errors []string

	assert := assert.New(t)
	want := &proto.UserResponse{
		Data:       RawToDtoUser(&mock.User1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	usrService := service.NewUserService(&mock.UserMockRepo{})
	usrRes, err := usrService.FindOne(mock.Context{}, &proto.FindOneUserRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, usrRes)
}

func TestFindOneErrNotFoundUser(t *testing.T) {
	mock.InitializeMockUser()

	errors := []string{"Not found user"}

	assert := assert.New(t)
	want := &proto.UserResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	usrService := service.NewUserService(&mock.UserMockErrRepo{})
	usrRes, _ := usrService.FindOne(mock.Context{}, &proto.FindOneUserRequest{Id: 1})

	assert.Equal(want, usrRes)
}

func TestCreateUser(t *testing.T) {
	mock.InitializeMockUser()

	var errors []string

	assert := assert.New(t)
	want := &proto.UserResponse{
		Data:       RawToDtoUser(&mock.User1),
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	usrService := service.NewUserService(&mock.UserMockRepo{})
	usrRes, err := usrService.Create(mock.Context{}, &mock.CreateUserReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, usrRes)
}

func TestUpdateUser(t *testing.T) {
	mock.InitializeMockUser()

	var errors []string

	assert := assert.New(t)
	want := &proto.UserResponse{
		Data:       RawToDtoUser(&mock.User1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	usrService := service.NewUserService(&mock.UserMockRepo{})
	usrRes, err := usrService.Update(mock.Context{}, &mock.UpdateUserReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, usrRes)
}

func TestUpdateErrNotFoundUser(t *testing.T) {
	mock.InitializeMockUser()

	errors := []string{"Not found user"}

	assert := assert.New(t)

	want := &proto.UserResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	usrService := service.NewUserService(&mock.UserMockErrRepo{})
	usrRes, _ := usrService.Update(mock.Context{}, &mock.UpdateUserReqMock)

	assert.Equal(want, usrRes)
}

func TestDeleteUser(t *testing.T) {
	mock.InitializeMockUser()

	var errors []string

	assert := assert.New(t)
	want := &proto.UserResponse{
		Data:       RawToDtoUser(&mock.User1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	usrService := service.NewUserService(&mock.UserMockRepo{})
	usrRes, err := usrService.Delete(mock.Context{}, &proto.DeleteUserRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, usrRes)
}

func TestDeleteErrNotFoundUser(t *testing.T) {
	mock.InitializeMockUser()

	errors := []string{"Not found user"}

	assert := assert.New(t)

	want := &proto.UserResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	usrService := service.NewUserService(&mock.UserMockErrRepo{})
	usrRes, _ := usrService.Delete(mock.Context{}, &proto.DeleteUserRequest{Id: 1})

	assert.Equal(want, usrRes)
}
