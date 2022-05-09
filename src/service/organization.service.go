package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"time"
)

type OrganizationService struct {
	client proto.OrganizationServiceClient
}

func NewOrganizationService(client proto.OrganizationServiceClient) *OrganizationService {
	return &OrganizationService{
		client: client,
	}
}

func (s *OrganizationService) FindOne(id uint) (res *proto.OrganizationResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.FindOne(ctx, &proto.FindOneOrganizationRequest{Id: int32(id)})
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return res, nil
	}

	return
}

func (s *OrganizationService) FindMulti(ids []uint32) (res *proto.OrganizationListResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.FindMulti(ctx, &proto.FindMultiOrganizationRequest{Ids: ids})
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return res, nil
	}

	return
}
