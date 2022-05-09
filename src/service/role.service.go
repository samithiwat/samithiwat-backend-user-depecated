package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"time"
)

type RoleService struct {
	client proto.RoleServiceClient
}

func NewRoleService(client proto.RoleServiceClient) *RoleService {
	return &RoleService{
		client: client,
	}
}

func (s *RoleService) FindOne(id uint) (res *proto.RoleResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.FindOne(ctx, &proto.FindOneRoleRequest{Id: int32(id)})
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return res, nil
	}

	return
}

func (s *RoleService) FindMulti(ids []uint32) (res *proto.RoleListResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.FindMulti(ctx, &proto.FindMultiRoleRequest{Ids: ids})
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return res, nil
	}

	return
}
