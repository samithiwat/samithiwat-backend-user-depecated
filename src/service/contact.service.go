package service

import "github.com/samithiwat/samithiwat-backend-user/src/proto"

type ContactService struct {
	client proto.ContactServiceClient
}

func NewContactService(client proto.ContactServiceClient) *ContactService {
	return &ContactService{
		client: client,
	}
}
