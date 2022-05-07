package service

import "github.com/samithiwat/samithiwat-backend-user/src/proto"

type OrganizationService struct {
	client proto.OrganizationServiceClient
}

func NewOrganizationService(client proto.OrganizationServiceClient) *OrganizationService {
	return &OrganizationService{
		client: client,
	}
}
