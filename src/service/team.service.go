package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"time"
)

type TeamService struct {
	client proto.TeamServiceClient
}

func NewTeamService(client proto.TeamServiceClient) *TeamService {
	return &TeamService{
		client: client,
	}
}

func (s *TeamService) FindOne(id uint) (res *proto.TeamResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.FindOne(ctx, &proto.FindOneTeamRequest{Id: int32(id)})
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return res, nil
	}

	return
}

func (s *TeamService) FindMulti(ids []uint32) (res *proto.TeamListResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.FindMulti(ctx, &proto.FindMultiTeamRequest{Ids: ids})
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return res, nil
	}

	return
}
