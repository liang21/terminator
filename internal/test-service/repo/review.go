package repo

import (
	"context"
	"errors"
	"github.com/liang21/terminator/internal/test-service/biz"
	"github.com/liang21/terminator/pkg/pagination"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type reviewRepo struct {
	db  *xorm.Engine
	rdb *redis.Client
}

func NewReviewRepo(db *xorm.Engine, rdb *redis.Client) biz.ReviewRepo {
	return &reviewRepo{
		db:  db,
		rdb: rdb,
	}
}
func (r reviewRepo) ListTestReview(ctx context.Context, meta pagination.ListMeta) (TestCaseReviewDto *biz.TestCaseReviewDTOList, err error) {
	var testReviews []*biz.TestReview
	offset := pagination.GetPageOffset(meta.PageSize, meta.PageToken)
	count, err := r.db.Limit(int(meta.PageSize), int(offset)).Desc("name").And("deleted = ?", 0).FindAndCount(&testReviews)
	if err != nil {
		return nil, err
	}
	testReviewDtoList := &biz.TestCaseReviewDTOList{}
	testReviewDtoList.TotalCount = count
	testReviewDtoList.Items = testReviews
	return testReviewDtoList, nil
}

func (r reviewRepo) GetTestReview(ctx context.Context, id int64) (*biz.TestReview, error) {
	testReview := &biz.TestReview{Deleted: 0}
	ok, err := r.db.Where("id = ?", id).Get(testReview)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("test case review not found")
	}
	return testReview, nil
}

func (r reviewRepo) CreateTestReview(ctx context.Context, TestCaseReview *biz.TestReview) error {
	result, err := r.db.Insert(TestCaseReview)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test case review create failed")
	}
	return nil
}

func (r reviewRepo) UpdateTestReview(ctx context.Context, id int64, TestCaseReview *biz.TestReview) error {
	result, err := r.db.Where("id = ?", id).Update(TestCaseReview)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test case review update failed")
	}
	return nil
}

func (r reviewRepo) DeleteTestReview(ctx context.Context, id int64) error {
	testReview := &biz.TestReview{Deleted: 0}
	result, err := r.db.Where("id = ?", id).Update(testReview)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test case review delete failed")
	}
	return nil
}

func (r reviewRepo) ListTestReviewCase(ctx context.Context, meta pagination.ListMeta) (TestReviewCaseDto *biz.TestReviewCaseDTOList, err error) {
	var testReviewCases []*biz.TestReviewCase
	offset := pagination.GetPageOffset(meta.PageSize, meta.PageToken)
	count, err := r.db.Limit(int(meta.PageSize), int(offset)).Desc("name").And("deleted = ?", 0).FindAndCount(&testReviewCases)
	if err != nil {
		return nil, err
	}
	testReviewCaseDtoList := &biz.TestReviewCaseDTOList{}
	testReviewCaseDtoList.TotalCount = count
	testReviewCaseDtoList.Items = testReviewCases
	return testReviewCaseDtoList, nil
}

func (r reviewRepo) GetTestReviewCase(ctx context.Context, id int64) (*biz.TestReviewCase, error) {
	testReviewCase := &biz.TestReviewCase{Deleted: 0}
	ok, err := r.db.Where("id = ?", id).Get(testReviewCase)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("test case review case not found")
	}
	return testReviewCase, nil
}

func (r reviewRepo) CreateTestReviewCase(ctx context.Context, TestReviewCase *biz.TestReviewCase) error {
	result, err := r.db.Insert(TestReviewCase)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test case review case create failed")
	}
	return nil
}

func (r reviewRepo) UpdateTestReviewCase(ctx context.Context, id int64, TestReviewCase *biz.TestReviewCase) error {
	result, err := r.db.Where("id = ?", id).Update(TestReviewCase)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test case review case update failed")
	}
	return nil
}

func (r reviewRepo) DeleteTestReviewCase(ctx context.Context, id int64) error {
	testReviewCase := &biz.TestReviewCase{Deleted: 0}
	result, err := r.db.Where("id = ?", id).Update(testReviewCase)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test case review case delete failed")
	}
	return nil
}

func (r reviewRepo) ListTestReviewReport(ctx context.Context, meta pagination.ListMeta) (TestReviewReportDto *biz.TestReviewReportDTOList, err error) {
	var testReviewReports []*biz.TestReviewReport
	offset := pagination.GetPageOffset(meta.PageSize, meta.PageToken)
	count, err := r.db.Limit(int(meta.PageSize), int(offset)).Desc("name").And("deleted = ?", 0).FindAndCount(&testReviewReports)
	if err != nil {
		return nil, err
	}
	testReviewReportDtoList := &biz.TestReviewReportDTOList{}
	testReviewReportDtoList.TotalCount = count
	testReviewReportDtoList.Items = testReviewReports
	return testReviewReportDtoList, nil
}

func (r reviewRepo) GetTestReviewReport(ctx context.Context, id int64) (*biz.TestReviewReport, error) {
	testReviewReport := &biz.TestReviewReport{Deleted: 0}
	ok, err := r.db.Where("id = ?", id).Get(testReviewReport)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("test case review report not found")
	}
	return testReviewReport, nil
}

func (r reviewRepo) CreateTestReviewReport(ctx context.Context, TestReviewReport *biz.TestReviewReport) error {
	result, err := r.db.Insert(TestReviewReport)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test case review report create failed")
	}
	return nil
}

func (r reviewRepo) UpdateTestReviewReport(ctx context.Context, id int64, TestReviewReport *biz.TestReviewReport) error {
	result, err := r.db.Where("id = ?", id).Update(TestReviewReport)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test case review report update failed")
	}
	return nil
}

func (r reviewRepo) DeleteTestReviewReport(ctx context.Context, id int64) error {
	testReviewReport := &biz.TestReviewReport{Deleted: 0}
	result, err := r.db.Where("id = ?", id).Update(testReviewReport)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("test case review report delete failed")
	}
	return nil
}
