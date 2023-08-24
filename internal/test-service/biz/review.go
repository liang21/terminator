package biz

import (
	"context"
	"errors"
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
	ReviewedEndAt  time.Time `json:"reviewed_end_at"`
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
	Id            int64     `json:"id" comment:"测试报告id"`
	ReviewId      int64     `json:"review_id" comment:"review_id"`
	CaseId        int64     `json:"name" comment:"测试用例id"`
	Status        string    `json:"status" comment:"状态"`
	Deleted       int64     `json:"deleted"`
	ReviewedEndAt time.Time `json:"reviewed_end_at"`
	CreateAt      time.Time `json:"create_at"`
	UpdateAt      time.Time `json:"update_at"`
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

type ReviewUsecase struct {
	repo ReviewRepo
}

func NewReviewUsecase(repo ReviewRepo) *ReviewUsecase {
	return &ReviewUsecase{repo: repo}
}

func (uc *ReviewUsecase) ListReview(ctx context.Context, meta pagination.ListMeta) (testreview *TestCaseReviewDTOList, err error) {
	testreview, err = uc.repo.ListTestReview(ctx, meta)
	if err != nil {
		return
	}
	return testreview, nil
}

func (uc *ReviewUsecase) GetReview(ctx context.Context, id int64) (testReview *TestReview, err error) {
	if id == 0 {
		return nil, errors.New("id is empty")
	}
	testReview, err = uc.repo.GetTestReview(ctx, id)
	if err != nil {
		return
	}
	return testReview, nil
}

func (uc *ReviewUsecase) CreateReview(ctx context.Context, testReview *TestReview) error {
	if err := uc.repo.CreateTestReview(ctx, testReview); err != nil {
		return err
	}
	return nil
}

func (uc *ReviewUsecase) UpdateReview(ctx context.Context, id int64, testReview *TestReview) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err := uc.repo.UpdateTestReview(ctx, id, testReview); err != nil {
		return err
	}
	return nil
}

func (uc *ReviewUsecase) DeleteReview(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err := uc.repo.DeleteTestReview(ctx, id); err != nil {
		return err
	}
	return nil
}

func (uc *ReviewUsecase) ListReviewCase(ctx context.Context, meta pagination.ListMeta) (testReviewCase *TestReviewCaseDTOList, err error) {
	testReviewCase, err = uc.repo.ListTestReviewCase(ctx, meta)
	if err != nil {
		return
	}
	return testReviewCase, nil
}

func (uc *ReviewUsecase) GetReviewCase(ctx context.Context, id int64) (testReviewCase *TestReviewCase, err error) {
	if id == 0 {
		return nil, errors.New("id is empty")
	}
	testReviewCase, err = uc.repo.GetTestReviewCase(ctx, id)
	if err != nil {
		return
	}
	return testReviewCase, nil
}

func (uc *ReviewUsecase) CreateReviewCase(ctx context.Context, testReviewCase *TestReviewCase) error {
	if err := uc.repo.CreateTestReviewCase(ctx, testReviewCase); err != nil {
		return err
	}
	return nil
}

func (uc *ReviewUsecase) UpdateReviewCase(ctx context.Context, id int64, testReviewCase *TestReviewCase) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err := uc.repo.UpdateTestReviewCase(ctx, id, testReviewCase); err != nil {
		return err
	}
	return nil
}

func (uc *ReviewUsecase) DeleteReviewCase(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err := uc.repo.DeleteTestReviewCase(ctx, id); err != nil {
		return err
	}
	return nil
}

func (uc *ReviewUsecase) ListReviewReport(ctx context.Context, meta pagination.ListMeta) (testReviewReport *TestReviewReportDTOList, err error) {
	testReviewReport, err = uc.repo.ListTestReviewReport(ctx, meta)
	if err != nil {
		return
	}
	return testReviewReport, nil
}

func (uc *ReviewUsecase) GetReviewReport(ctx context.Context, id int64) (testReviewReport *TestReviewReport, err error) {
	if id == 0 {
		return nil, errors.New("id is empty")
	}
	testReviewReport, err = uc.repo.GetTestReviewReport(ctx, id)
	if err != nil {
		return
	}
	return testReviewReport, nil
}

func (uc *ReviewUsecase) CreateReviewReport(ctx context.Context, testReviewReport *TestReviewReport) error {
	if err := uc.repo.CreateTestReviewReport(ctx, testReviewReport); err != nil {
		return err
	}
	return nil
}

func (uc *ReviewUsecase) UpdateReviewReport(ctx context.Context, id int64, testReviewReport *TestReviewReport) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err := uc.repo.UpdateTestReviewReport(ctx, id, testReviewReport); err != nil {
		return err
	}
	return nil
}

func (uc *ReviewUsecase) DeleteReviewReport(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err := uc.repo.DeleteTestReviewReport(ctx, id); err != nil {
		return err
	}
	return nil
}
