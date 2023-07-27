package repo

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/liang21/terminator/internal/system/biz"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

func NewUserStore(db *xorm.Engine, rdb *redis.Client) biz.UserRepo {
	return &userRepo{
		db:  db,
		rdb: rdb,
	}
}

type userRepo struct {
	db  *xorm.Engine
	rdb *redis.Client
}

func (u userRepo) ListUser(ctx context.Context) ([]*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) UpdateUser(ctx context.Context, id int64, user *biz.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) DeleteUser(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) GetUserLike(ctx context.Context, id int64) (rv int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) IncUserLike(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
