package biz

import (
	"context"
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
