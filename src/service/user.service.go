package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-user/src/model"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"net/http"
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

func (s *UserService) FindAll(_ context.Context, req *proto.FindAllUserRequest) (res *proto.UserPaginationResponse, err error) {

	meta := proto.PaginationMetadata{
		ItemsPerPage: req.Limit,
		CurrentPage:  req.Page,
	}

	var orgs []*model.User
	var errors []string

	res = &proto.UserPaginationResponse{
		Data: &proto.UserPagination{
			Items: nil,
			Meta:  &meta,
		},
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindAll(&meta, &orgs)
	if err != nil {
		errors = append(errors, err.Error())
		res.StatusCode = http.StatusBadRequest
		return
	}

	var result []*proto.User

	for _, org := range orgs {
		result = append(result, RawToDtoUser(org))
	}

	res.Data.Items = result

	return
}

func (s *UserService) FindOne(_ context.Context, req *proto.FindOneUserRequest) (res *proto.UserResponse, err error) {
	org := model.User{}
	var errors []string

	res = &proto.UserResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindOne(uint(req.Id), &org)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoUser(&org)
	res.Data = result
	return
}

func (s *UserService) FindMulti(_ context.Context, req *proto.FindMultiUserRequest) (res *proto.UserListResponse, err error) {
	var orgs []*model.User
	var errors []string

	res = &proto.UserListResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindMulti(req.Ids, &orgs)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	var result []*proto.User
	for _, org := range orgs {
		result = append(result, RawToDtoUser(org))
	}

	res.Data = result

	return
}

func (s *UserService) Create(_ context.Context, req *proto.CreateUserRequest) (res *proto.UserResponse, err error) {
	org := DtoToRawUser(req.User)
	var errors []string

	res = &proto.UserResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	err = s.repository.Create(org)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusUnprocessableEntity
		return
	}

	result := RawToDtoUser(org)
	res.Data = result

	return
}

func (s *UserService) Update(_ context.Context, req *proto.UpdateUserRequest) (res *proto.UserResponse, err error) {
	org := DtoToRawUser(req.User)
	var errors []string

	res = &proto.UserResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Update(uint(org.ID), org)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoUser(org)
	res.Data = result

	return
}

func (s *UserService) Delete(_ context.Context, req *proto.DeleteUserRequest) (res *proto.UserResponse, err error) {
	org := model.User{}
	var errors []string

	res = &proto.UserResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Delete(uint(req.Id), &org)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoUser(&org)
	res.Data = result

	return
}
