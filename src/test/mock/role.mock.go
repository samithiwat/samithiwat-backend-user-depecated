package mock

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"google.golang.org/grpc"
	"net/http"
)

var Role1 proto.Role
var Role2 proto.Role
var Role3 proto.Role
var Role4 proto.Role
var Roles []*proto.Role

type RoleMockClient struct {
}

func (c *RoleMockClient) FindAll(ctx context.Context, in *proto.FindAllRoleRequest, opts ...grpc.CallOption) (*proto.RolePaginationResponse, error) {
	return nil, nil
}

func (*RoleMockClient) FindOne(_ context.Context, in *proto.FindOneRoleRequest, opts ...grpc.CallOption) (*proto.RoleResponse, error) {
	return &proto.RoleResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &Role1,
	}, nil
}

func (*RoleMockClient) FindMulti(_ context.Context, in *proto.FindMultiRoleRequest, opts ...grpc.CallOption) (*proto.RoleListResponse, error) {
	return &proto.RoleListResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       Roles,
	}, nil
}

func (*RoleMockClient) Create(_ context.Context, in *proto.CreateRoleRequest, opts ...grpc.CallOption) (*proto.RoleResponse, error) {
	return nil, nil
}

func (*RoleMockClient) Update(_ context.Context, in *proto.UpdateRoleRequest, opts ...grpc.CallOption) (*proto.RoleResponse, error) {
	return nil, nil
}

func (*RoleMockClient) Delete(_ context.Context, in *proto.DeleteRoleRequest, opts ...grpc.CallOption) (*proto.RoleResponse, error) {
	return nil, nil
}

type RoleMockErrClient struct {
}

func (c *RoleMockErrClient) FindAll(ctx context.Context, in *proto.FindAllRoleRequest, opts ...grpc.CallOption) (*proto.RolePaginationResponse, error) {
	return nil, nil
}

func (*RoleMockErrClient) FindOne(_ context.Context, in *proto.FindOneRoleRequest, opts ...grpc.CallOption) (*proto.RoleResponse, error) {
	return &proto.RoleResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found role"},
		Data:       nil,
	}, errors.New("Grpc error")
}

func (*RoleMockErrClient) FindMulti(_ context.Context, in *proto.FindMultiRoleRequest, opts ...grpc.CallOption) (*proto.RoleListResponse, error) {
	return nil, nil
}

func (*RoleMockErrClient) Create(_ context.Context, in *proto.CreateRoleRequest, opts ...grpc.CallOption) (*proto.RoleResponse, error) {
	return nil, nil
}

func (*RoleMockErrClient) Update(_ context.Context, in *proto.UpdateRoleRequest, opts ...grpc.CallOption) (*proto.RoleResponse, error) {
	return nil, nil
}

func (*RoleMockErrClient) Delete(_ context.Context, in *proto.DeleteRoleRequest, opts ...grpc.CallOption) (*proto.RoleResponse, error) {
	return nil, nil
}

func InitializeMockRole() {
	Role1 = proto.Role{
		Id:          1,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Role2 = proto.Role{
		Id:          2,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Role3 = proto.Role{
		Id:          3,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Role4 = proto.Role{
		Id:          4,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Roles = append(Roles, &Role1, &Role2, &Role3, &Role4)
}
