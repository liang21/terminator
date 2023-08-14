package repo

import (
	"context"
	"github.com/liang21/terminator/internal/test-service/biz"
	"github.com/liang21/terminator/pkg/pagination"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type testCasetRepo struct {
	db  *xorm.Engine
	rdb *redis.Client
}

func NewTestCaseRepo(db *xorm.Engine, rdb *redis.Client) biz.TestCaseRepo {
	return &testCasetRepo{
		db:  db,
		rdb: rdb,
	}
}
func (t testCasetRepo) ListTestCase(ctx context.Context, meta pagination.ListMeta) (TestCaseDto *biz.TestCaseDTOList, err error) {
	//TODO implement me
	panic("implement me")
}

func (t testCasetRepo) GetTestCase(ctx context.Context, id int64) (*biz.TestCase, error) {
	//TODO implement me
	panic("implement me")
}

func (t testCasetRepo) CreateTestCase(ctx context.Context, TestCase *biz.TestCase) error {
	//TODO implement me
	panic("implement me")
}

func (t testCasetRepo) UpdateTestCase(ctx context.Context, id int64, TestCase *biz.TestCase) error {
	//TODO implement me
	panic("implement me")
}

func (t testCasetRepo) DeleteTestCase(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
