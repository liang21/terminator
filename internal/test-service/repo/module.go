package repo

import (
	"context"
	"github.com/liang21/terminator/internal/test-service/biz"
	"github.com/liang21/terminator/pkg/pagination"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type testModuletRepo struct {
	db  *xorm.Engine
	rdb *redis.Client
}

func NewTestModuleRepo(db *xorm.Engine, rdb *redis.Client) biz.TestModuleRepo {
	return &testModuletRepo{
		db:  db,
		rdb: rdb,
	}
}

func (t testModuletRepo) ListTestModule(ctx context.Context, meta pagination.ListMeta) (TestModuleDto *biz.TestModuleDTOList, err error) {
	//TODO implement me
	panic("implement me")
}

func (t testModuletRepo) GetTestModule(ctx context.Context, id int64) (*biz.TestModule, error) {
	//TODO implement me
	panic("implement me")
}

func (t testModuletRepo) CreateTestModule(ctx context.Context, TestModule *biz.TestModule) error {
	//TODO implement me
	panic("implement me")
}

func (t testModuletRepo) UpdateTestModule(ctx context.Context, id int64, TestModule *biz.TestModule) error {
	//TODO implement me
	panic("implement me")
}

func (t testModuletRepo) DeleteTestModule(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
