package repo

import (
	"context"
	"errors"
	"github.com/liang21/terminator/internal/system-service/biz"
	"github.com/liang21/terminator/pkg/pagination"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

func NewUserRepo(db *xorm.Engine, rdb *redis.Client) biz.UserRepo {
	return &userRepo{
		db:  db,
		rdb: rdb,
	}
}

type userRepo struct {
	db  *xorm.Engine
	rdb *redis.Client
}

func (u *userRepo) ListUser(ctx context.Context, meta pagination.ListMeta) (*biz.UserDTOList, error) {
	var users []*biz.User
	offset := pagination.GetPageOffset(meta.PageSize, meta.PageToken)
	count, err := u.db.Limit(int(meta.PageSize), int(offset)).Desc("name").And("deleted = ?", 0).FindAndCount(&users)
	if err != nil {
		return nil, err
	}
	userDtoList := &biz.UserDTOList{}
	userDtoList.TotalCount = count
	userDtoList.Items = users
	return userDtoList, nil
}

func (u *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	user := &biz.User{Deleted: 0}
	ok, err := u.db.Where("id = ?", id).Get(user)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	result, err := u.db.Insert(user)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("user create failed")
	}
	return nil
}

func (u *userRepo) UpdateUser(ctx context.Context, id int64, user *biz.User) error {
	result, err := u.db.Update(user, &biz.User{Id: id})
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (u *userRepo) DeleteUser(ctx context.Context, id int64) error {
	user := &biz.User{Deleted: 1}
	result, err := u.db.Update(user, &biz.User{Id: id})
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("user not found")
	}
	return nil
}
