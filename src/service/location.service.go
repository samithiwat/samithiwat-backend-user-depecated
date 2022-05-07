package service

import "github.com/samithiwat/samithiwat-backend-user/src/proto"

type LocationService struct {
	client proto.LocationServiceClient
}

func NewLocationService(client proto.LocationServiceClient) *LocationService {
	return &LocationService{
		client: client,
	}
}
