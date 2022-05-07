package service

import "github.com/samithiwat/samithiwat-backend-user/src/proto"

type TeamService struct {
	client proto.TeamServiceClient
}

func NewTeamService(client proto.TeamServiceClient) *TeamService {
	return &TeamService{
		client: client,
	}
}
