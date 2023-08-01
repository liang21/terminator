package data

import (
	"context"
	v1 "github.com/liang21/terminator/api/system/v1"
	"github.com/liang21/terminator/internal/system/biz"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
	user *biz.UserUsecase
}

func NewUserService(user *biz.UserUsecase) *UserService {
	return &UserService{user: user}
}

func (u *UserService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	//TODO implement me
	return nil, nil
}

func (u *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user, err := u.user.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetUserReply{User: &v1.User{Id: user.Id, Name: user.Name, Email: user.Email, Phone: user.Phone, RoleId: user.RoleId, Password: user.Password}}, nil
}

func (u *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	//TODO implement me
	return nil, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	err := u.user.Update(ctx, req.GetId(), &biz.User{Id: req.GetId(), Name: req.GetName(), Email: req.GetEmail(), Phone: req.GetPhone(), RoleId: req.GetRoleId(), Password: req.GetPassword()})
	return &v1.UpdateUserReply{}, err
}

func (u *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	//TODO implement me
	return nil, nil
}
