package service

import (
	"context"
	v1 "github.com/liang21/terminator/api/system/v1"
	"github.com/liang21/terminator/internal/system/biz"
	"github.com/liang21/terminator/pkg/pagination"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
	user *biz.UserUsecase
}

func NewUserService(user *biz.UserUsecase) *UserService {
	return &UserService{user: user}
}

func (u *UserService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	meta := pagination.ListMeta{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
	}
	users, err := u.user.List(ctx, meta)
	user := make([]*v1.User, 0, len(users.Items))
	for _, item := range users.Items {
		user = append(user, &v1.User{Id: item.Id, Name: item.Name, Email: item.Email, Phone: item.Phone, RoleId: item.RoleId, Password: item.Password})
	}
	if err != nil {
		return nil, err
	}
	return &v1.ListUserReply{Total: users.TotalCount, Results: user}, nil
}

func (u *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user, err := u.user.Get(ctx, req.GetId())
	return &v1.GetUserReply{User: &v1.User{Id: user.Id, Name: user.Name, Email: user.Email, Phone: user.Phone, RoleId: user.RoleId, Password: user.Password}}, err
}

func (u *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	err := u.user.Create(ctx, &biz.User{Name: req.GetName(), Email: req.GetEmail(), Phone: req.GetPhone(), RoleId: req.GetRoleId(), Password: req.GetPassword()})
	return &v1.CreateUserReply{User: &v1.User{Name: req.Name, Email: req.Email, Phone: req.Phone, RoleId: req.RoleId, Password: req.Password}}, err
}

func (u *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	err := u.user.Update(ctx, req.GetId(), &biz.User{Id: req.GetId(), Name: req.GetName(), Email: req.GetEmail(), Phone: req.GetPhone(), RoleId: req.GetRoleId(), Password: req.GetPassword()})
	return &v1.UpdateUserReply{}, err
}

func (u *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	err := u.user.Delete(ctx, req.GetId())
	return &v1.DeleteUserReply{}, err
}
