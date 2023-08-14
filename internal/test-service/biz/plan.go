package biz

import (
	"context"
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
	CreateTestPlanCase(ctx context.Context, TestPlanCase *TestPlanCase) error
	UpdateTestPlanCase(ctx context.Context, id int64, TestPlanCase *TestPlanCase) error
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
