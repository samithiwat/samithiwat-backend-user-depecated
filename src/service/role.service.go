package service

import "github.com/samithiwat/samithiwat-backend-user/src/proto"

type RoleService struct {
	client proto.RoleServiceClient
}

func NewRoleService(client proto.RoleServiceClient) *RoleService {
	return &RoleService{
		client: client,
	}
}
