package biz

import (
	"context"
	"errors"
	"github.com/liang21/terminator/pkg/pagination"
	"time"
)

type TestPlan struct {
	Id             int64     `json:"id"`
	ProjectId      int64     `json:"project_id"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	Stage          string    `json:"stage"`
	Description    string    `json:"description"`
	ReviewTotal    int64     `json:"review_total"`
	ReviewPassRule int64     `json:"review_pass_rule"`
	PlanedStartAt  time.Time `json:"planed_start_at"`
	PlanedEndAt    time.Time `json:"planed_end_at"`
	ActualStartAt  time.Time `json:"actual_start_at"`
	ActualEndAt    time.Time `json:"actual_end_at"`
	Deleted        int64     `json:"deleted"`
	CreateAt       time.Time `json:"create_at"`
	UpdateAt       time.Time `json:"update_at"`
}

func (p TestPlan) TableName() string {
	return "test_plan"
}

type TestPlanDTOList struct {
	TotalCount int64       `json:"totalCount,omitempty"` //总数
	Items      []*TestPlan `json:"data"`                 //数据
}
type TestPlanRepo interface {
	// ListTestPlan db
	ListTestPlan(ctx context.Context, meta pagination.ListMeta) (TestPlanDto *TestPlanDTOList, err error)
	GetTestPlan(ctx context.Context, id int64) (*TestPlan, error)
	CreateTestPlan(ctx context.Context, TestPlan *TestPlan) error
	UpdateTestPlan(ctx context.Context, id int64, TestPlan *TestPlan) error
	DeleteTestPlan(ctx context.Context, id int64) error
}

type TestPlanCase struct {
	Id       int64     `json:"id" comment:"测试报告id"`
	PlanId   int64     `json:"test_plan_id" comment:"测试计划id"`
	CaseId   int64     `json:"name" comment:"测试用例id"`
	Status   string    `json:"status" comment:"状态"`
	Result   string    `json:"result" comment:"执行结果"`
	Deleted  int64     `json:"deleted"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
type TestPlanCaseDTOList struct {
	TotalCount int64           `json:"totalCount,omitempty"` //总数
	Items      []*TestPlanCase `json:"data"`                 //数据
}

func (c TestPlanCase) TableName() string {
	return "test_plan_case"
}

type TestPlanCaseRepo interface {
	// ListTestPlanCase db
	ListTestPlanCase(ctx context.Context, meta pagination.ListMeta) (TestPlanCaseDto *TestPlanCaseDTOList, err error)
	GetTestPlanCase(ctx context.Context, id int64) (*TestPlanCase, error)
	CreateTestPlanCase(ctx context.Context, planCase *TestPlanCase) error
	UpdateTestPlanCase(ctx context.Context, id int64, planCase *TestPlanCase) error
	DeleteTestPlanCase(ctx context.Context, id int64) error
}

type TestPlanReport struct {
	Id         int64     `json:"id" comment:"测试报告id"`
	TestPlanId int64     `json:"test_plan_id" comment:"测试计划id"`
	Name       string    `json:"name" comment:"测试报告名称"`
	Status     string    `json:"status" comment:"状态"`
	Success    int32     `json:"success" comment:"成功数"`
	Deleted    int64     `json:"deleted"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
}
type TestPlanReportDTOList struct {
	TotalCount int64             `json:"totalCount,omitempty"` //总数
	Items      []*TestPlanReport `json:"data"`                 //数据
}

func (r TestPlanReport) TableName() string {
	return "test_plan_report"
}

type TestPlanReportRepo interface {
	// ListTestPlanReport db
	ListTestPlanReport(ctx context.Context, meta pagination.ListMeta) (TestPlanReportDto *TestPlanReportDTOList, err error)
	GetTestPlanReport(ctx context.Context, id int64) (*TestPlanReport, error)
	CreateTestPlanReport(ctx context.Context, TestPlanReport *TestPlanReport) error
	UpdateTestPlanReport(ctx context.Context, id int64, TestPlanReport *TestPlanReport) error
	DeleteTestPlanReport(ctx context.Context, id int64) error
}

type PlanRepo interface {
	TestPlanRepo
	TestPlanCaseRepo
	TestPlanReportRepo
}

type PlanUsecase struct {
	repo PlanRepo
}

func NewPlanUsecase(repo PlanRepo) *PlanUsecase {
	return &PlanUsecase{repo: repo}
}

func (uc *PlanUsecase) ListPlan(ctx context.Context, meta pagination.ListMeta) (plan *TestPlanDTOList, err error) {
	plan, err = uc.repo.ListTestPlan(ctx, meta)
	if err != nil {
		return
	}
	return plan, nil
}

func (uc *PlanUsecase) GetPlanById(ctx context.Context, id int64) (plan *TestPlan, err error) {
	if id == 0 {
		return nil, errors.New("id is empty")
	}
	plan, err = uc.repo.GetTestPlan(ctx, id)
	if err != nil {
		return
	}
	return plan, nil
}

func (uc *PlanUsecase) CreatePlan(ctx context.Context, plan *TestPlan) error {
	if err := uc.repo.CreateTestPlan(ctx, plan); err != nil {
		return err
	}
	return nil
}

func (uc *PlanUsecase) UpdatePlan(ctx context.Context, id int64, plan *TestPlan) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err := uc.repo.UpdateTestPlan(ctx, id, plan); err != nil {
		return err
	}
	return nil
}

func (uc *PlanUsecase) DeletePlan(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err := uc.repo.DeleteTestPlan(ctx, id); err != nil {
		return err
	}
	return nil
}

func (uc *PlanUsecase) ListPlanCase(ctx context.Context, meta pagination.ListMeta) (planCase *TestPlanCaseDTOList, err error) {
	planCase, err = uc.repo.ListTestPlanCase(ctx, meta)
	if err != nil {
		return
	}
	return planCase, nil
}

func (uc *PlanUsecase) GetPlanCaseById(ctx context.Context, id int64) (planCase *TestPlanCase, err error) {
	if id == 0 {
		return nil, errors.New("id is empty")
	}
	planCase, err = uc.repo.GetTestPlanCase(ctx, id)
	if err != nil {
		return
	}
	return planCase, nil
}

func (uc *PlanUsecase) CreatePlanCase(ctx context.Context, planCase *TestPlanCase) error {
	if err := uc.repo.CreateTestPlanCase(ctx, planCase); err != nil {
		return err
	}
	return nil
}

func (uc *PlanUsecase) UpdatePlanCase(ctx context.Context, id int64, planCase *TestPlanCase) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err := uc.repo.UpdateTestPlanCase(ctx, id, planCase); err != nil {
		return err
	}
	return nil
}

func (uc *PlanUsecase) DeletePlanCase(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err := uc.repo.DeleteTestPlanCase(ctx, id); err != nil {
		return err
	}
	return nil
}

func (uc *PlanUsecase) ListPlanReport(ctx context.Context, meta pagination.ListMeta) (planReport *TestPlanReportDTOList, err error) {
	planReport, err = uc.repo.ListTestPlanReport(ctx, meta)
	if err != nil {
		return
	}
	return planReport, nil
}

func (uc *PlanUsecase) GetPlanReportById(ctx context.Context, id int64) (planReport *TestPlanReport, err error) {
	if id == 0 {
		return nil, errors.New("id is empty")
	}
	planReport, err = uc.repo.GetTestPlanReport(ctx, id)
	if err != nil {
		return
	}
	return planReport, nil
}

func (uc *PlanUsecase) CreatePlanReport(ctx context.Context, planReport *TestPlanReport) error {
	if err := uc.repo.CreateTestPlanReport(ctx, planReport); err != nil {
		return err
	}
	return nil
}

func (uc *PlanUsecase) UpdatePlanReport(ctx context.Context, id int64, planReport *TestPlanReport) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err := uc.repo.UpdateTestPlanReport(ctx, id, planReport); err != nil {
		return err
	}
	return nil
}

func (uc *PlanUsecase) DeletePlanReport(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err := uc.repo.DeleteTestPlanReport(ctx, id); err != nil {
		return err
	}
	return nil
}
