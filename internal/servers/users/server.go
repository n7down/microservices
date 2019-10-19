package users

import (
	"context"

	"github.com/n7down/microservices/internal/users/pb"
)

type UsersServer struct{}

func (u *UsersServer) UsersGetPassword(ctx context.Context, req *users_pb.UsersGetPasswordRequest) (*users_pb.UsersGetPasswordResponse, error) {
	return &users_pb.UsersGetPasswordResponse{}, nil
}

func (u *UsersServer) UsersCreate(ctx context.Context, req *users_pb.UsersCreateRequest) (*users_pb.UsersCreateResponse, error) {
	return &users_pb.UsersCreateResponse{}, nil
}

func (u *UsersServer) UsersUpdate(ctx context.Context, req *users_pb.UsersUpdateRequest) (*users_pb.UsersUpdateResponse, error) {
	return &users_pb.UsersUpdateResponse{}, nil
}

func (u *UsersServer) UsersList(ctx context.Context, req *users_pb.UsersListRequest) (*users_pb.UsersListResponse, error) {
	return &users_pb.UsersListResponse{}, nil
}

func (u *UsersServer) UsersByID(ctx context.Context, req *users_pb.UsersByIDRequest) (*users_pb.UsersByIDResponse, error) {
	return &users_pb.UsersByIDResponse{}, nil
}

// FIXME: just sets the is_active = false in the database
func (u *UsersServer) UsersDelete(ctx context.Context, req *users_pb.UsersDeleteRequest) (*users_pb.UsersDeleteResponse, error) {
	return &users_pb.UsersDeleteResponse{}, nil
}
