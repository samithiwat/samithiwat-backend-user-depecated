package test

import (
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"github.com/samithiwat/samithiwat-backend-user/src/service"
	"github.com/samithiwat/samithiwat-backend-user/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindOneOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	var errors []string

	assert := assert.New(t)
	want := &proto.OrganizationResponse{
		Data:       &mock.Organization1,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	locService := service.NewOrganizationService(&mock.OrganizationMockClient{})
	locRes, err := locService.FindOne(1)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}

func TestFindOneErrGrpcOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	errors := []string{"Not found organization", "Grpc error"}

	assert := assert.New(t)
	want := &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	locService := service.NewOrganizationService(&mock.OrganizationMockErrClient{})
	locRes, err := locService.FindOne(1)

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, locRes)
}

func TestFindMultiOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	var errors []string

	assert := assert.New(t)
	want := &proto.OrganizationListResponse{
		Data:       mock.Organizations,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	locService := service.NewOrganizationService(&mock.OrganizationMockClient{})
	locRes, err := locService.FindMulti([]uint32{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}
