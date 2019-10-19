package users

import (
	"context"

	"github.com/google/uuid"
	"github.com/n7down/microservices/internal/client/users/pb"
)

type UsersServer struct {
	db *UsersDB
}

func NewUsersServer() (*UsersServer, error) {
	usersServer := &UsersServer{}
	db, err := NewUsersDB()
	if err != nil {
		return usersServer, err
	}
	return &UsersServer{db: db}, nil
}

func (u *UsersServer) UsersGetPassword(ctx context.Context, req *users_pb.UsersGetPasswordRequest) (*users_pb.UsersGetPasswordResponse, error) {
	res := &users_pb.UsersGetPasswordResponse{}
	id, password, err := u.db.GetPassword(req.Username)
	if err != nil {
		return res, err
	}
	res.ID = id
	res.Password = password
	return res, nil
}

func (u *UsersServer) UsersCreate(ctx context.Context, req *users_pb.UsersCreateRequest) (*users_pb.UsersCreateResponse, error) {
	res := &users_pb.UsersCreateResponse{}
	uuid := uuid.New()
	err := u.db.Create(uuid.String(), req.Username, req.Password, req.Firstname, req.Lastname)
	if err != nil {
		return res, err
	}
	res.ID = uuid.String()
	return res, nil
}

// FIXME: implement
func (u *UsersServer) UsersUpdate(ctx context.Context, req *users_pb.UsersUpdateRequest) (*users_pb.UsersUpdateResponse, error) {
	return &users_pb.UsersUpdateResponse{}, nil
}

// FIXME: implement
func (u *UsersServer) UsersList(ctx context.Context, req *users_pb.UsersListRequest) (*users_pb.UsersListResponse, error) {
	return &users_pb.UsersListResponse{}, nil
}

// FIXME: implement
func (u *UsersServer) UsersByID(ctx context.Context, req *users_pb.UsersByIDRequest) (*users_pb.UsersByIDResponse, error) {
	return &users_pb.UsersByIDResponse{}, nil
}

func (u *UsersServer) UsersDelete(ctx context.Context, req *users_pb.UsersDeleteRequest) (*users_pb.UsersDeleteResponse, error) {
	res := &users_pb.UsersDeleteResponse{}
	err := u.db.Delete(req.ID)
	if err != nil {
		return res, err
	}
	res.ID = req.ID
	return res, nil
}