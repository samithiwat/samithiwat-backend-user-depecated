package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"time"
)

type ContactService struct {
	client proto.ContactServiceClient
}

func NewContactService(client proto.ContactServiceClient) *ContactService {
	return &ContactService{
		client: client,
	}
}

func (s *ContactService) FindOne(id uint) (res *proto.ContactResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.FindOne(ctx, &proto.FindOneContactRequest{Id: int32(id)})
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return res, nil
	}

	return
}

func (s *ContactService) FindMulti(ids []uint32) (res *proto.ContactListResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.FindMulti(ctx, &proto.FindMultiContactRequest{Ids: ids})
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return res, nil
	}

	return
}
