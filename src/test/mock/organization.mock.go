package mock

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"google.golang.org/grpc"
	"net/http"
)

var Organization1 proto.Organization
var Organization2 proto.Organization
var Organization3 proto.Organization
var Organization4 proto.Organization
var Organizations []*proto.Organization

type OrganizationMockClient struct {
}

func (c *OrganizationMockClient) FindAll(ctx context.Context, in *proto.FindAllOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationPaginationResponse, error) {
	return nil, nil
}

func (*OrganizationMockClient) FindOne(_ context.Context, in *proto.FindOneOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return &proto.OrganizationResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &Organization1,
	}, nil
}

func (*OrganizationMockClient) FindMulti(_ context.Context, in *proto.FindMultiOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationListResponse, error) {
	return &proto.OrganizationListResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       Organizations,
	}, nil
}

func (*OrganizationMockClient) Create(_ context.Context, in *proto.CreateOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return nil, nil
}

func (*OrganizationMockClient) Update(_ context.Context, in *proto.UpdateOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return nil, nil
}

func (*OrganizationMockClient) Delete(_ context.Context, in *proto.DeleteOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return nil, nil
}

type OrganizationMockErrClient struct {
}

func (c *OrganizationMockErrClient) FindAll(ctx context.Context, in *proto.FindAllOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationPaginationResponse, error) {
	return nil, nil
}

func (*OrganizationMockErrClient) FindOne(_ context.Context, in *proto.FindOneOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return &proto.OrganizationResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found organization"},
		Data:       nil,
	}, errors.New("Grpc error")
}

func (*OrganizationMockErrClient) FindMulti(_ context.Context, in *proto.FindMultiOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationListResponse, error) {
	return nil, nil
}

func (*OrganizationMockErrClient) Create(_ context.Context, in *proto.CreateOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return nil, nil
}

func (*OrganizationMockErrClient) Update(_ context.Context, in *proto.UpdateOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return nil, nil
}

func (*OrganizationMockErrClient) Delete(_ context.Context, in *proto.DeleteOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return nil, nil
}

func InitializeMockOrganization() {
	Organization1 = proto.Organization{
		Id:          1,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organization2 = proto.Organization{
		Id:          2,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organization3 = proto.Organization{
		Id:          3,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organization4 = proto.Organization{
		Id:          4,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organizations = append(Organizations, &Organization1, &Organization2, &Organization3, &Organization4)
}
