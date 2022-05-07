package mock

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"google.golang.org/grpc"
	"net/http"
)

var Location1 proto.Location
var Location2 proto.Location
var Location3 proto.Location
var Location4 proto.Location
var Locations []*proto.Location

type LocationMockClient struct {
}

func (*LocationMockClient) FindOne(_ context.Context, in *proto.FindOneLocationRequest, opts ...grpc.CallOption) (*proto.LocationResponse, error) {
	return &proto.LocationResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &Location1,
	}, nil
}

func (*LocationMockClient) FindMulti(_ context.Context, in *proto.FindMultiLocationRequest, opts ...grpc.CallOption) (*proto.LocationListResponse, error) {
	return &proto.LocationListResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       Locations,
	}, nil
}

func (*LocationMockClient) Create(_ context.Context, in *proto.CreateLocationRequest, opts ...grpc.CallOption) (*proto.LocationResponse, error) {
	return nil, nil
}

func (*LocationMockClient) Update(_ context.Context, in *proto.UpdateLocationRequest, opts ...grpc.CallOption) (*proto.LocationResponse, error) {
	return nil, nil
}

func (*LocationMockClient) Delete(_ context.Context, in *proto.DeleteLocationRequest, opts ...grpc.CallOption) (*proto.LocationResponse, error) {
	return nil, nil
}

type LocationMockErrClient struct {
}

func (*LocationMockErrClient) FindOne(_ context.Context, in *proto.FindOneLocationRequest, opts ...grpc.CallOption) (*proto.LocationResponse, error) {
	return &proto.LocationResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found contact"},
		Data:       nil,
	}, errors.New("Grpc error")
}

func (*LocationMockErrClient) FindMulti(_ context.Context, in *proto.FindMultiLocationRequest, opts ...grpc.CallOption) (*proto.LocationListResponse, error) {
	return nil, nil
}

func (*LocationMockErrClient) Create(_ context.Context, in *proto.CreateLocationRequest, opts ...grpc.CallOption) (*proto.LocationResponse, error) {
	return nil, nil
}

func (*LocationMockErrClient) Update(_ context.Context, in *proto.UpdateLocationRequest, opts ...grpc.CallOption) (*proto.LocationResponse, error) {
	return nil, nil
}

func (*LocationMockErrClient) Delete(_ context.Context, in *proto.DeleteLocationRequest, opts ...grpc.CallOption) (*proto.LocationResponse, error) {
	return nil, nil
}

func InitializeMockLocation() {
	Location1 = proto.Location{
		Id:       1,
		Address:  faker.Word(),
		District: faker.Word(),
		Province: faker.Word(),
		Country:  faker.Word(),
		Zipcode:  faker.Word(),
	}

	Location2 = proto.Location{
		Id:       2,
		Address:  faker.Word(),
		District: faker.Word(),
		Province: faker.Word(),
		Country:  faker.Word(),
		Zipcode:  faker.Word(),
	}

	Location3 = proto.Location{
		Id:       3,
		Address:  faker.Word(),
		District: faker.Word(),
		Province: faker.Word(),
		Country:  faker.Word(),
		Zipcode:  faker.Word(),
	}

	Location4 = proto.Location{
		Id:       4,
		Address:  faker.Word(),
		District: faker.Word(),
		Province: faker.Word(),
		Country:  faker.Word(),
		Zipcode:  faker.Word(),
	}

	Locations = append(Locations, &Location1, &Location2, &Location3, &Location4)
}
