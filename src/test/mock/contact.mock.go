package mock

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"google.golang.org/grpc"
	"net/http"
)

var Contact1 proto.Contact
var Contact2 proto.Contact
var Contact3 proto.Contact
var Contact4 proto.Contact
var Contacts []*proto.Contact

type ContactMockClient struct {
}

func (*ContactMockClient) FindOne(_ context.Context, in *proto.FindOneContactRequest, opts ...grpc.CallOption) (*proto.ContactResponse, error) {
	return &proto.ContactResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &Contact1,
	}, nil
}

func (*ContactMockClient) FindMulti(_ context.Context, in *proto.FindMultiContactRequest, opts ...grpc.CallOption) (*proto.ContactListResponse, error) {
	return &proto.ContactListResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       Contacts,
	}, nil
}

func (*ContactMockClient) Create(_ context.Context, in *proto.CreateContactRequest, opts ...grpc.CallOption) (*proto.ContactResponse, error) {
	return nil, nil
}

func (*ContactMockClient) Update(_ context.Context, in *proto.UpdateContactRequest, opts ...grpc.CallOption) (*proto.ContactResponse, error) {
	return nil, nil
}

func (*ContactMockClient) Delete(_ context.Context, in *proto.DeleteContactRequest, opts ...grpc.CallOption) (*proto.ContactResponse, error) {
	return nil, nil
}

type ContactMockErrClient struct {
}

func (*ContactMockErrClient) FindOne(_ context.Context, in *proto.FindOneContactRequest, opts ...grpc.CallOption) (*proto.ContactResponse, error) {
	return &proto.ContactResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found contact"},
		Data:       nil,
	}, errors.New("Grpc error")
}

func (*ContactMockErrClient) FindMulti(_ context.Context, in *proto.FindMultiContactRequest, opts ...grpc.CallOption) (*proto.ContactListResponse, error) {
	return nil, nil
}

func (*ContactMockErrClient) Create(_ context.Context, in *proto.CreateContactRequest, opts ...grpc.CallOption) (*proto.ContactResponse, error) {
	return nil, nil
}

func (*ContactMockErrClient) Update(_ context.Context, in *proto.UpdateContactRequest, opts ...grpc.CallOption) (*proto.ContactResponse, error) {
	return nil, nil
}

func (*ContactMockErrClient) Delete(_ context.Context, in *proto.DeleteContactRequest, opts ...grpc.CallOption) (*proto.ContactResponse, error) {
	return nil, nil
}

func InitializeMockContact() {
	Contact1 = proto.Contact{
		Id:        1,
		Facebook:  faker.URL(),
		Instagram: faker.URL(),
		Linkedin:  faker.URL(),
		Twitter:   faker.URL(),
	}

	Contact2 = proto.Contact{
		Id:        2,
		Facebook:  faker.URL(),
		Instagram: faker.URL(),
		Linkedin:  faker.URL(),
		Twitter:   faker.URL(),
	}

	Contact3 = proto.Contact{
		Id:        3,
		Facebook:  faker.URL(),
		Instagram: faker.URL(),
		Linkedin:  faker.URL(),
		Twitter:   faker.URL(),
	}

	Contact4 = proto.Contact{
		Id:        4,
		Facebook:  faker.URL(),
		Instagram: faker.URL(),
		Linkedin:  faker.URL(),
		Twitter:   faker.URL(),
	}

	Contacts = append(Contacts, &Contact1, &Contact2, &Contact3, &Contact4)
}
