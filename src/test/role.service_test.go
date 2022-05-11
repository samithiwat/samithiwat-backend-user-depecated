package test

import (
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"github.com/samithiwat/samithiwat-backend-user/src/service"
	"github.com/samithiwat/samithiwat-backend-user/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindOneRole(t *testing.T) {
	mock.InitializeMockRole()

	var errors []string

	assert := assert.New(t)
	want := &proto.RoleResponse{
		Data:       &mock.Role1,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	locService := service.NewRoleService(&mock.RoleMockClient{})
	locRes, err := locService.FindOne(1)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}

func TestFindOneErrGrpcRole(t *testing.T) {
	mock.InitializeMockRole()

	errors := []string{"Not found role", "Grpc error"}

	assert := assert.New(t)
	want := &proto.RoleResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	locService := service.NewRoleService(&mock.RoleMockErrClient{})
	locRes, _ := locService.FindOne(1)

	assert.Equal(want, locRes)
}

func TestFindMultiRole(t *testing.T) {
	mock.InitializeMockRole()

	var errors []string

	assert := assert.New(t)
	want := &proto.RoleListResponse{
		Data:       mock.Roles,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	locService := service.NewRoleService(&mock.RoleMockClient{})
	locRes, err := locService.FindMulti([]uint32{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}
