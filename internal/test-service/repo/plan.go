package repo

import (
	"context"
	"errors"
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
	var testPlans []*biz.TestPlan
	offset := pagination.GetPageOffset(meta.PageSize, meta.PageToken)
	count, err := p.db.Limit(int(meta.PageSize), int(offset)).Desc("name").And("deleted = ?", 0).FindAndCount(&testPlans)
	if err != nil {
		return nil, err
	}
	testPlanDtoList := &biz.TestPlanDTOList{}
	testPlanDtoList.TotalCount = count
	testPlanDtoList.Items = testPlans
	return testPlanDtoList, nil
}

func (p planRepo) GetTestPlan(ctx context.Context, id int64) (*biz.TestPlan, error) {
	testPlan := &biz.TestPlan{Deleted: 0}
	ok, err := p.db.Where("id = ?", id).Get(testPlan)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("test plan not found")
	}
	return testPlan, nil
}

func (p planRepo) CreateTestPlan(ctx context.Context, testPlan *biz.TestPlan) error {
	result, err := p.db.Insert(testPlan)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test plan create failed")
	}
	return nil
}

func (p planRepo) UpdateTestPlan(ctx context.Context, id int64, testPlan *biz.TestPlan) error {
	result, err := p.db.Where("id = ?", id).Update(testPlan)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test plan update failed")
	}
	return nil
}

func (p planRepo) DeleteTestPlan(ctx context.Context, id int64) error {
	testPlan := &biz.TestPlan{Deleted: 1}
	result, err := p.db.Where("id = ?", id).Update(testPlan)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test plan delete failed")
	}
	return nil
}

func (p planRepo) ListTestPlanCase(ctx context.Context, meta pagination.ListMeta) (TestPlanCaseDto *biz.TestPlanCaseDTOList, err error) {
	var testPlanCases []*biz.TestPlanCase
	offset := pagination.GetPageOffset(meta.PageSize, meta.PageToken)
	count, err := p.db.Limit(int(meta.PageSize), int(offset)).Desc("name").And("deleted = ?", 0).FindAndCount(&testPlanCases)
	if err != nil {
		return nil, err
	}
	testPlanCaseDtoList := &biz.TestPlanCaseDTOList{}
	testPlanCaseDtoList.TotalCount = count
	testPlanCaseDtoList.Items = testPlanCases
	return testPlanCaseDtoList, nil
}

func (p planRepo) GetTestPlanCase(ctx context.Context, id int64) (*biz.TestPlanCase, error) {
	planCase := &biz.TestPlanCase{}
	ok, err := p.db.Where("id = ?", id).Get(planCase)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("test plan case not found")
	}
	return planCase, nil
}

func (p planRepo) CreateTestPlanCase(ctx context.Context, planCase *biz.TestPlanCase) error {
	result, err := p.db.Insert(planCase)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test plan case create failed")
	}
	return nil
}

func (p planRepo) UpdateTestPlanCase(ctx context.Context, id int64, planCase *biz.TestPlanCase) error {
	result, err := p.db.Where("id = ?", id).Update(planCase)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test plan case update failed")
	}
	return nil
}

func (p planRepo) DeleteTestPlanCase(ctx context.Context, id int64) error {
	planCase := &biz.TestPlanCase{Deleted: 1}
	result, err := p.db.Where("id = ?", id).Update(planCase)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test plan case delete failed")
	}
	return nil
}

func (p planRepo) ListTestPlanReport(ctx context.Context, meta pagination.ListMeta) (TestPlanReportDto *biz.TestPlanReportDTOList, err error) {
	var testPlanReports []*biz.TestPlanReport
	offset := pagination.GetPageOffset(meta.PageSize, meta.PageToken)
	count, err := p.db.Limit(int(meta.PageSize), int(offset)).Desc("name").And("deleted = ?", 0).FindAndCount(&testPlanReports)
	if err != nil {
		return nil, err
	}
	testPlanReportDtoList := &biz.TestPlanReportDTOList{}
	testPlanReportDtoList.TotalCount = count
	testPlanReportDtoList.Items = testPlanReports
	return testPlanReportDtoList, nil
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
