package biz

import (
	"context"
	"errors"
	"github.com/liang21/terminator/pkg/pagination"
	"time"
)

type User struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	RoleId   int64     `json:"role_id"`
	Password string    `json:"password"`
	Deleted  int64     `json:"deleted"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
type UserDTOList struct {
	TotalCount int64   `json:"totalCount,omitempty"` //总数
	Items      []*User `json:"data"`                 //数据
}

func (u User) TableName() string {
	return "user"
}

type UserRepo interface {
	// ListUser db
	ListUser(ctx context.Context, meta pagination.ListMeta) (userDto *UserDTOList, err error)
	GetUser(ctx context.Context, id int64) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id int64, user *User) error
	DeleteUser(ctx context.Context, id int64) error
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) List(ctx context.Context, meta pagination.ListMeta) (userDtoList *UserDTOList, err error) {
	userDtoList, err = uc.repo.ListUser(ctx, meta)
	if err != nil {
		return
	}
	return userDtoList, nil
}

func (uc *UserUsecase) Get(ctx context.Context, id int64) (user *User, err error) {
	if id == 0 {
		return nil, errors.New("id is empty")
	}
	user, err = uc.repo.GetUser(ctx, id)
	if err != nil {
		return
	}
	return user, nil
}

func (uc *UserUsecase) Create(ctx context.Context, user *User) error {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUsecase) Update(ctx context.Context, id int64, user *User) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	return uc.repo.UpdateUser(ctx, id, user)
}

func (uc *UserUsecase) Delete(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	return uc.repo.DeleteUser(ctx, id)
}
