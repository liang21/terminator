package repo

import (
	"context"
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
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) GetTestReview(ctx context.Context, id int64) (*biz.TestReview, error) {
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) CreateTestReview(ctx context.Context, TestCaseReview *biz.TestReview) error {
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) UpdateTestReview(ctx context.Context, id int64, TestCaseReview *biz.TestReview) error {
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) DeleteTestReview(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) ListTestReviewCase(ctx context.Context, meta pagination.ListMeta) (TestReviewCaseDto *biz.TestReviewCaseDTOList, err error) {
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) GetTestReviewCase(ctx context.Context, id int64) (*biz.TestReviewCase, error) {
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) CreateTestReviewCase(ctx context.Context, TestReviewCase *biz.TestReviewCase) error {
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) UpdateTestReviewCase(ctx context.Context, id int64, TestReviewCase *biz.TestReviewCase) error {
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) DeleteTestReviewCase(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) ListTestReviewReport(ctx context.Context, meta pagination.ListMeta) (TestReviewReportDto *biz.TestReviewReportDTOList, err error) {
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) GetTestReviewReport(ctx context.Context, id int64) (*biz.TestReviewReport, error) {
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) CreateTestReviewReport(ctx context.Context, TestReviewReport *biz.TestReviewReport) error {
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) UpdateTestReviewReport(ctx context.Context, id int64, TestReviewReport *biz.TestReviewReport) error {
	//TODO implement me
	panic("implement me")
}

func (r reviewRepo) DeleteTestReviewReport(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
