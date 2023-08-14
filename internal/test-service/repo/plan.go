package repo

import (
	"context"
	"github.com/liang21/terminator/internal/test-service/biz"
	"github.com/liang21/terminator/pkg/pagination"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type planRepo struct {
	db  *xorm.Engine
	rdb *redis.Client
}

func NewTestPlanRepo(db *xorm.Engine, rdb *redis.Client) biz.PlanRepo {
	return &planRepo{
		db:  db,
		rdb: rdb,
	}
}
func (p planRepo) ListTestPlan(ctx context.Context, meta pagination.ListMeta) (TestPlanDto *biz.TestPlanDTOList, err error) {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) GetTestPlan(ctx context.Context, id int64) (*biz.TestPlan, error) {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) CreateTestPlan(ctx context.Context, TestPlan *biz.TestPlan) error {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) UpdateTestPlan(ctx context.Context, id int64, TestPlan *biz.TestPlan) error {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) DeleteTestPlan(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) ListTestPlanCase(ctx context.Context, meta pagination.ListMeta) (TestPlanCaseDto *biz.TestPlanCaseDTOList, err error) {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) GetTestPlanCase(ctx context.Context, id int64) (*biz.TestPlanCase, error) {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) CreateTestPlanCase(ctx context.Context, TestPlanCase *biz.TestPlanCase) error {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) UpdateTestPlanCase(ctx context.Context, id int64, TestPlanCase *biz.TestPlanCase) error {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) DeleteTestPlanCase(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) ListTestPlanReport(ctx context.Context, meta pagination.ListMeta) (TestPlanReportDto *biz.TestPlanReportDTOList, err error) {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) GetTestPlanReport(ctx context.Context, id int64) (*biz.TestPlanReport, error) {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) CreateTestPlanReport(ctx context.Context, TestPlanReport *biz.TestPlanReport) error {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) UpdateTestPlanReport(ctx context.Context, id int64, TestPlanReport *biz.TestPlanReport) error {
	//TODO implement me
	panic("implement me")
}

func (p planRepo) DeleteTestPlanReport(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
