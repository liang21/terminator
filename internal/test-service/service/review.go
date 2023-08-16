package service

import (
	"context"
	v1 "github.com/liang21/terminator/api/test/v1"
	"github.com/liang21/terminator/internal/test-service/biz"
)

type ReviewService struct {
	v1.UnimplementedTestReviewServiceServer
	v1.UnimplementedTestReviewCaseServiceServer
	v1.UnimplementedTestReviewReportServiceServer
	review *biz.ReviewUsecase
}

func NewReviewService(review *biz.ReviewUsecase) *ReviewService {
	return &ReviewService{review: review}
}

func (s *ReviewService) CreateReview(ctx context.Context, req *v1.CreateTestReviewRequest) (*v1.CreateTestReviewReply, error) {
	return &v1.CreateTestReviewReply{}, nil
}

func (s *ReviewService) UpdateReview(ctx context.Context, req *v1.UpdateTestReviewRequest) (*v1.UpdateTestReviewReply, error) {
	return &v1.UpdateTestReviewReply{}, nil
}

func (s *ReviewService) DeleteReview(ctx context.Context, req *v1.DeleteTestReviewRequest) (*v1.DeleteTestReviewReply, error) {
	return &v1.DeleteTestReviewReply{}, nil
}

func (s *ReviewService) GetReview(ctx context.Context, req *v1.GetTestReviewRequest) (*v1.GetTestReviewReply, error) {
	return &v1.GetTestReviewReply{}, nil
}

func (s *ReviewService) ListReview(ctx context.Context, req *v1.ListTestReviewRequest) (*v1.ListTestReviewReply, error) {
	return &v1.ListTestReviewReply{}, nil
}

func (s *ReviewService) CreateReviewCase(ctx context.Context, req *v1.CreateTestReviewCaseRequest) (*v1.CreateTestReviewCaseReply, error) {
	return &v1.CreateTestReviewCaseReply{}, nil
}

func (s *ReviewService) UpdateReviewCase(ctx context.Context, req *v1.UpdateTestReviewCaseRequest) (*v1.UpdateTestReviewCaseReply, error) {
	return &v1.UpdateTestReviewCaseReply{}, nil
}

func (s *ReviewService) DeleteReviewCase(ctx context.Context, req *v1.DeleteTestReviewCaseRequest) (*v1.DeleteTestReviewCaseReply, error) {
	return &v1.DeleteTestReviewCaseReply{}, nil
}

func (s *ReviewService) GetReviewCase(ctx context.Context, req *v1.GetTestReviewCaseRequest) (*v1.GetTestReviewCaseReply, error) {
	return &v1.GetTestReviewCaseReply{}, nil
}

func (s *ReviewService) ListReviewCase(ctx context.Context, req *v1.ListTestReviewCaseRequest) (*v1.ListTestReviewCaseReply, error) {
	return &v1.ListTestReviewCaseReply{}, nil
}

func (s *ReviewService) GetReviewReport(ctx context.Context, req *v1.GetTestReviewReportRequest) (*v1.GetTestReviewReportReply, error) {
	return &v1.GetTestReviewReportReply{}, nil
}
