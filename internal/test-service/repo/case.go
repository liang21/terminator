package repo

import (
	"context"
	"errors"
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
func (t *testCasetRepo) ListTestCase(ctx context.Context, meta pagination.ListMeta) (TestCaseDto *biz.TestCaseDTOList, err error) {
	var testcases []*biz.TestCase
	offset := pagination.GetPageOffset(meta.PageSize, meta.PageToken)
	count, err := t.db.Limit(int(meta.PageSize), int(offset)).Desc("name").And("deleted = ?", 0).FindAndCount(&testcases)
	if err != nil {
		return nil, err
	}
	testcaseDtoList := &biz.TestCaseDTOList{}
	testcaseDtoList.TotalCount = count
	testcaseDtoList.Items = testcases
	return testcaseDtoList, nil
}

func (t *testCasetRepo) GetTestCase(ctx context.Context, id int64) (*biz.TestCase, error) {
	testcase := &biz.TestCase{Deleted: 0}
	ok, err := t.db.Where("id = ?", id).Get(testcase)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("test case not found")
	}
	return testcase, nil
}

func (t *testCasetRepo) CreateTestCase(ctx context.Context, testcase *biz.TestCase) error {
	result, err := t.db.Insert(testcase)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test case create failed")
	}
	return nil
}

func (t *testCasetRepo) UpdateTestCase(ctx context.Context, id int64, testcase *biz.TestCase) error {
	resulr, err := t.db.Where("id = ?", id).Update(testcase)
	if err != nil {
		return err
	}
	if resulr == 0 {
		return errors.New("test case update failed")
	}
	return nil
}

func (t *testCasetRepo) DeleteTestCase(ctx context.Context, id int64) error {
	testcase := &biz.TestCase{Deleted: 1}
	result, err := t.db.Where("id = ?", id).Update(testcase)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test case delete failed")
	}
	return nil
}
