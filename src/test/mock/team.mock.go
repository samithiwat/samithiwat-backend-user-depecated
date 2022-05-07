package mock

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"google.golang.org/grpc"
	"net/http"
)

var Team1 proto.Team
var Team2 proto.Team
var Team3 proto.Team
var Team4 proto.Team
var Teams []*proto.Team

type TeamMockClient struct {
}

func (c *TeamMockClient) FindAll(ctx context.Context, in *proto.FindAllTeamRequest, opts ...grpc.CallOption) (*proto.TeamPaginationResponse, error) {
	return nil, nil
}

func (*TeamMockClient) FindOne(_ context.Context, in *proto.FindOneTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return &proto.TeamResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &Team1,
	}, nil
}

func (*TeamMockClient) FindMulti(_ context.Context, in *proto.FindMultiTeamRequest, opts ...grpc.CallOption) (*proto.TeamListResponse, error) {
	return &proto.TeamListResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       Teams,
	}, nil
}

func (*TeamMockClient) Create(_ context.Context, in *proto.CreateTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return nil, nil
}

func (*TeamMockClient) Update(_ context.Context, in *proto.UpdateTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return nil, nil
}

func (*TeamMockClient) Delete(_ context.Context, in *proto.DeleteTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return nil, nil
}

type TeamMockErrClient struct {
}

func (c *TeamMockErrClient) FindAll(ctx context.Context, in *proto.FindAllTeamRequest, opts ...grpc.CallOption) (*proto.TeamPaginationResponse, error) {
	return nil, nil
}

func (*TeamMockErrClient) FindOne(_ context.Context, in *proto.FindOneTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return &proto.TeamResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found organization"},
		Data:       nil,
	}, errors.New("Grpc error")
}

func (*TeamMockErrClient) FindMulti(_ context.Context, in *proto.FindMultiTeamRequest, opts ...grpc.CallOption) (*proto.TeamListResponse, error) {
	return nil, nil
}

func (*TeamMockErrClient) Create(_ context.Context, in *proto.CreateTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return nil, nil
}

func (*TeamMockErrClient) Update(_ context.Context, in *proto.UpdateTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return nil, nil
}

func (*TeamMockErrClient) Delete(_ context.Context, in *proto.DeleteTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return nil, nil
}

func InitializeMockTeam() {
	Team1 = proto.Team{
		Id:          1,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Team2 = proto.Team{
		Id:          2,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Team3 = proto.Team{
		Id:          3,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Team4 = proto.Team{
		Id:          4,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Teams = append(Teams, &Team1, &Team2, &Team3, &Team4)
}
