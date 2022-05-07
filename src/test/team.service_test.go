package test

import (
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"github.com/samithiwat/samithiwat-backend-user/src/service"
	"github.com/samithiwat/samithiwat-backend-user/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindOneTeam(t *testing.T) {
	mock.InitializeMockTeam()

	var errors []string

	assert := assert.New(t)
	want := &proto.TeamResponse{
		Data:       &mock.Team1,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	locService := service.NewTeamService(&mock.TeamMockClient{})
	locRes, err := locService.FindOne(1)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}

func TestFindOneErrGrpcTeam(t *testing.T) {
	mock.InitializeMockTeam()

	errors := []string{"Not found location", "Grpc error"}

	assert := assert.New(t)
	want := &proto.TeamResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	locService := service.NewTeamService(&mock.TeamMockErrClient{})
	locRes, err := locService.FindOne(1)

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, locRes)
}

func TestFindMultiTeam(t *testing.T) {
	mock.InitializeMockTeam()

	var errors []string

	assert := assert.New(t)
	want := &proto.TeamListResponse{
		Data:       mock.Teams,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	locService := service.NewTeamService(&mock.TeamMockClient{})
	locRes, err := locService.FindMulti([]uint32{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}
