package test

import (
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"github.com/samithiwat/samithiwat-backend-user/src/service"
	"github.com/samithiwat/samithiwat-backend-user/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindOneLocation(t *testing.T) {
	mock.InitializeMockLocation()

	var errors []string

	assert := assert.New(t)
	want := &proto.LocationResponse{
		Data:       &mock.Location1,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	locService := service.NewLocationService(&mock.LocationMockClient{})
	locRes, err := locService.FindOne(1)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}

func TestFindOneErrGrpcLocation(t *testing.T) {
	mock.InitializeMockLocation()

	errors := []string{"Not found location", "Grpc error"}

	assert := assert.New(t)
	want := &proto.LocationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	locService := service.NewLocationService(&mock.LocationMockErrClient{})
	locRes, _ := locService.FindOne(1)

	assert.Equal(want, locRes)
}

func TestFindMultiLocation(t *testing.T) {
	mock.InitializeMockLocation()

	var errors []string

	assert := assert.New(t)
	want := &proto.LocationListResponse{
		Data:       mock.Locations,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	locService := service.NewLocationService(&mock.LocationMockClient{})
	locRes, err := locService.FindMulti([]uint32{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}
