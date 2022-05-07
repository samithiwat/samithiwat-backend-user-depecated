package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"time"
)

type LocationService struct {
	client proto.LocationServiceClient
}

func NewLocationService(client proto.LocationServiceClient) *LocationService {
	return &LocationService{
		client: client,
	}
}

func (s *LocationService) FindOne(id uint) (res *proto.LocationResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.FindOne(ctx, &proto.FindOneLocationRequest{Id: int32(id)})
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return
	}

	return
}

func (s *LocationService) FindMulti(ids []uint32) (res *proto.LocationListResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.FindMulti(ctx, &proto.FindMultiLocationRequest{Ids: ids})
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return
	}

	return
}
