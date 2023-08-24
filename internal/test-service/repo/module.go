package repo

import (
	"context"
	"errors"
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

func (m testModuletRepo) ListTestModule(ctx context.Context, meta pagination.ListMeta) (module *biz.TestModuleDTOList, err error) {
	var modules []*biz.TestModule
	offset := pagination.GetPageOffset(meta.PageSize, meta.PageToken)
	count, err := m.db.Limit(int(meta.PageSize), int(offset)).Desc("name").And("deleted = ?", 0).FindAndCount(&modules)
	if err != nil {
		return nil, err
	}
	moduleDtoList := &biz.TestModuleDTOList{}
	moduleDtoList.TotalCount = count
	moduleDtoList.Items = modules
	return moduleDtoList, nil
}

func (m testModuletRepo) GetTestModule(ctx context.Context, id int64) (*biz.TestModule, error) {
	module := &biz.TestModule{Deleted: 0}
	ok, err := m.db.Where("id = ?", id).Get(module)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("test module not found")
	}
	return module, nil
}

func (m testModuletRepo) CreateTestModule(ctx context.Context, module *biz.TestModule) error {
	result, err := m.db.Insert(module)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test module create failed")
	}
	return nil
}

func (m testModuletRepo) UpdateTestModule(ctx context.Context, id int64, module *biz.TestModule) error {
	result, err := m.db.Where("id = ?", id).Update(module)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test module update failed")
	}
	return nil
}

func (m testModuletRepo) DeleteTestModule(ctx context.Context, id int64) error {
	module := &biz.TestModule{Deleted: 1}
	result, err := m.db.Where("id = ?", id).Update(module)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test module delete failed")
	}
	return nil
}
