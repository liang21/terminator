package biz

import (
	"context"
	"errors"
	"time"
)

type User struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	RoleId   int64     `json:"role_id"`
	Password string    `json:"password"`
	Delete   int64     `json:"delete"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (u User) TableName() string {
	return "user"
}

type UserRepo interface {
	// db
	ListUser(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id int64, user *User) error
	DeleteUser(ctx context.Context, id int64) error
	// redis
	GetUserLike(ctx context.Context, id int64) (rv int64, err error)
	IncUserLike(ctx context.Context, id int64) error
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) List(ctx context.Context) (ps []*User, err error) {
	ps, err = uc.repo.ListUser(ctx)
	if err != nil {
		return
	}
	return
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
	if user.Name == "" {
		return errors.New("name is empty")
	}
	if user.Phone == "" {
		return errors.New("phone is empty")
	}
	if user.Email == "" {
		return errors.New("email is empty")
	}
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUsecase) Update(ctx context.Context, id int64, user *User) error {
	return uc.repo.UpdateUser(ctx, id, user)
}

func (uc *UserUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteUser(ctx, id)
}
