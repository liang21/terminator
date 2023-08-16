package biz

import (
	"context"
	"errors"
	"github.com/liang21/terminator/pkg/pagination"
)

type TestCase struct {
	Id                int64  `json:"id"`
	ModuleId          int64  `json:"module_id"`
	ProjectId         int64  `json:"project_id"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Maintainer        string `json:"maintainer"`
	Priority          string `json:"priority"`
	Method            string `json:"method"`
	Prerequisite      string `json:"prerequisite"`
	Steps             string `json:"steps"`
	Remark            string `json:"remark"`
	ReviewStatus      int64  `json:"review_status"`
	Status            string `json:"status"`
	StepDescription   string `json:"step_description"`
	ExpectedResult    string `json:"expected_result"`
	CreateTestCase    string `json:"create_TestCase"`
	OriginalStatus    string `json:"original_status"`
	LastExecuteResult string `json:"last_execute_result"`
	Deleted           int64  `json:"deleted"`
	CreateAt          string `json:"create_at"`
	UpdateAt          string `json:"update_at"`
}
type TestCaseDTOList struct {
	TotalCount int64       `json:"totalCount,omitempty"` //总数
	Items      []*TestCase `json:"data"`                 //数据
}

func (u TestCase) TableName() string {
	return "TestCase"
}

type TestCaseUsecase struct {
	repo TestCaseRepo
}

func NewTestCaseUsecase(repo TestCaseRepo) *TestCaseUsecase {
	return &TestCaseUsecase{repo: repo}
}

type TestCaseRepo interface {
	// ListTestCase db
	ListTestCase(ctx context.Context, meta pagination.ListMeta) (TestCaseDto *TestCaseDTOList, err error)
	GetTestCase(ctx context.Context, id int64) (*TestCase, error)
	CreateTestCase(ctx context.Context, TestCase *TestCase) error
	UpdateTestCase(ctx context.Context, id int64, TestCase *TestCase) error
	DeleteTestCase(ctx context.Context, id int64) error
}

func (uc *TestCaseUsecase) List(ctx context.Context, meta pagination.ListMeta) (testcase *TestCaseDTOList, err error) {
	testcase, err = uc.repo.ListTestCase(ctx, meta)
	if err != nil {
		return
	}
	return testcase, nil
}

func (uc *TestCaseUsecase) Get(ctx context.Context, id int64) (testcase *TestCase, err error) {
	if id == 0 {
		return nil, errors.New("id is empty")
	}
	testcase, err = uc.repo.GetTestCase(ctx, id)
	if err != nil {
		return
	}
	return testcase, nil
}

func (uc *TestCaseUsecase) Create(ctx context.Context, testcase *TestCase) (err error) {
	if err = uc.repo.CreateTestCase(ctx, testcase); err != nil {
		return err
	}
	return
}

func (uc *TestCaseUsecase) Update(ctx context.Context, id int64, testcase *TestCase) (err error) {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err = uc.repo.UpdateTestCase(ctx, id, testcase); err != nil {
		return err
	}
	return
}

func (uc *TestCaseUsecase) Delete(ctx context.Context, id int64) (err error) {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err = uc.repo.DeleteTestCase(ctx, id); err != nil {
		return err
	}
	return
}
