package biz

import (
	"context"
	"github.com/liang21/terminator/pkg/pagination"
	"time"
)

type TestReview struct {
	Id             int64     `json:"id" comment:"测试用例评审id"`
	ProjectId      int64     `json:"project_id" comment:"项目id"`
	Name           string    `json:"name" comment:"测试用例评审名称"`
	Status         string    `json:"status" comment:"状态"`
	Description    string    `json:"description" comment:"描述"`
	ReviewTotal    int64     `json:"review_total" comment:"用例数"`
	ReviewPassRule int64     `json:"review_pass_rule" comment:"评审通过率"`
	Deleted        int64     `json:"deleted"`
	CreateAt       time.Time `json:"create_at"`
	UpdateAt       time.Time `json:"update_at"`
}
type TestCaseReviewDTOList struct {
	TotalCount int64         `json:"totalCount,omitempty"` //总数
	Items      []*TestReview `json:"data"`                 //数据
}

func (m TestReview) TableName() string {
	return "test_case_review"
}

type TestReviewRepo interface {
	// ListTestReview db
	ListTestReview(ctx context.Context, meta pagination.ListMeta) (TestCaseReviewDto *TestCaseReviewDTOList, err error)
	GetTestReview(ctx context.Context, id int64) (*TestReview, error)
	CreateTestReview(ctx context.Context, TestCaseReview *TestReview) error
	UpdateTestReview(ctx context.Context, id int64, TestCaseReview *TestReview) error
	DeleteTestReview(ctx context.Context, id int64) error
}
type TestReviewCase struct {
	Id       int64     `json:"id" comment:"测试报告id"`
	ReviewId int64     `json:"review_id" comment:"review_id"`
	CaseId   int64     `json:"name" comment:"测试用例id"`
	Status   string    `json:"status" comment:"状态"`
	Deleted  int64     `json:"deleted"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
type TestReviewCaseDTOList struct {
	TotalCount int64             `json:"totalCount,omitempty"` //总数
	Items      []*TestReviewCase `json:"data"`                 //数据
}

func (c TestReviewCase) TableName() string {
	return "test_review_case"
}

type TestReviewCaseRepo interface {
	// ListTestReviewCase db
	ListTestReviewCase(ctx context.Context, meta pagination.ListMeta) (TestReviewCaseDto *TestReviewCaseDTOList, err error)
	GetTestReviewCase(ctx context.Context, id int64) (*TestReviewCase, error)
	CreateTestReviewCase(ctx context.Context, TestReviewCase *TestReviewCase) error
	UpdateTestReviewCase(ctx context.Context, id int64, TestReviewCase *TestReviewCase) error
	DeleteTestReviewCase(ctx context.Context, id int64) error
}

type TestReviewReport struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	ProjectId int64     `json:"project_id"`
	Deleted   int64     `json:"deleted"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type TestReviewReportDTOList struct {
	TotalCount int64               `json:"totalCount,omitempty"` //总数
	Items      []*TestReviewReport `json:"data"`                 //数据
}

func (m TestReviewReport) TableName() string {
	return "test_review_report"
}

type TestReviewReportRepo interface {
	// ListTestReviewReport db
	ListTestReviewReport(ctx context.Context, meta pagination.ListMeta) (TestReviewReportDto *TestReviewReportDTOList, err error)
	GetTestReviewReport(ctx context.Context, id int64) (*TestReviewReport, error)
	CreateTestReviewReport(ctx context.Context, TestReviewReport *TestReviewReport) error
	UpdateTestReviewReport(ctx context.Context, id int64, TestReviewReport *TestReviewReport) error
	DeleteTestReviewReport(ctx context.Context, id int64) error
}
type ReviewRepo interface {
	TestReviewRepo
	TestReviewCaseRepo
	TestReviewReportRepo
}
