package service

import (
	"context"
	project_v1 "github.com/liang21/terminator/api/system/v1"
	plan_v1 "github.com/liang21/terminator/api/test/v1"
	v1 "github.com/liang21/terminator/api/test/v1"
	"github.com/liang21/terminator/internal/test-service/biz"
	"github.com/liang21/terminator/pkg/pagination"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

type ReviewService struct {
	v1.UnimplementedTestReviewServiceServer
	v1.UnimplementedTestReviewCaseServiceServer
	v1.UnimplementedTestReviewReportServiceServer
	review         *biz.ReviewUsecase
	projectClient  project_v1.ProjectServiceClient
	planClient     plan_v1.PlanServiceClient
	reviewClient   v1.TestReviewServiceClient
	testcaseClient v1.TestCaseServiceClient
}

func NewReviewService(review *biz.ReviewUsecase) *ReviewService {
	return &ReviewService{review: review}
}

func (s *ReviewService) CreateReview(ctx context.Context, req *v1.CreateTestReviewRequest) (*v1.CreateTestReviewReply, error) {
	if _, err := s.projectClient.GetProject(ctx, &project_v1.GetProjectRequest{Id: req.GetProjectId()}); err != nil {
		return nil, err
	}
	if err := s.review.CreateReview(ctx, &biz.TestReview{Name: req.GetName(), ProjectId: req.GetProjectId(), Status: req.GetStatus(), Description: req.GetDescription(), ReviewedEndAt: req.GetReviewedEndAt().AsTime()}); err != nil {
		return nil, err
	}
	return &v1.CreateTestReviewReply{}, nil
}

func (s *ReviewService) UpdateReview(ctx context.Context, req *v1.UpdateTestReviewRequest) (*v1.UpdateTestReviewReply, error) {
	if _, err := s.projectClient.GetProject(ctx, &project_v1.GetProjectRequest{Id: req.GetProjectId()}); err != nil {
		return nil, err
	}
	if err := s.review.UpdateReview(ctx, req.GetId(), &biz.TestReview{Name: req.GetName(), ProjectId: req.GetProjectId(), Status: req.GetStatus(), Description: req.GetDescription(), ReviewedEndAt: req.GetReviewedEndAt().AsTime()}); err != nil {
		return nil, err
	}
	return &v1.UpdateTestReviewReply{}, nil
}

func (s *ReviewService) DeleteReview(ctx context.Context, req *v1.DeleteTestReviewRequest) (*v1.DeleteTestReviewReply, error) {
	if err := s.review.DeleteReview(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &v1.DeleteTestReviewReply{}, nil
}

func (s *ReviewService) GetReview(ctx context.Context, req *v1.GetTestReviewRequest) (*v1.GetTestReviewReply, error) {
	testReview, err := s.review.GetReviewById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetTestReviewReply{TestReview: &v1.TestReview{Id: testReview.Id, ProjectId: testReview.ProjectId, Name: testReview.Name, Status: testReview.Status, Description: testReview.Description, ReviewTotal: testReview.ReviewTotal, ReviewPassRule: testReview.ReviewPassRule, ReviewedEndAt: timestamppb.New(testReview.ReviewedEndAt.Local())}}, nil
}

func (s *ReviewService) ListReview(ctx context.Context, req *v1.ListTestReviewRequest) (*v1.ListTestReviewReply, error) {
	parseIndex, err := strconv.ParseInt(req.GetPageToken(), 10, 64)
	if err != nil {
		return nil, err
	}
	testReviews, err := s.review.ListReview(ctx, pagination.ListMeta{PageSize: int64(req.GetPageSize()), PageToken: parseIndex})
	if err != nil {
		return nil, err
	}
	testReview := make([]*v1.TestReview, 0, len(testReviews.Items))
	for _, item := range testReviews.Items {
		testReview = append(testReview, &v1.TestReview{Id: item.Id, ProjectId: item.ProjectId, Name: item.Name, Status: item.Status, Description: item.Description, ReviewTotal: item.ReviewTotal, ReviewPassRule: item.ReviewPassRule, ReviewedEndAt: timestamppb.New(item.ReviewedEndAt.Local())})
	}
	return &v1.ListTestReviewReply{Total: testReviews.TotalCount, Results: testReview}, nil
}

func (s *ReviewService) CreateReviewCase(ctx context.Context, req *v1.CreateTestReviewCaseRequest) (*v1.CreateTestReviewCaseReply, error) {
	if _, err := s.reviewClient.GetTestReview(ctx, &v1.GetTestReviewRequest{Id: req.GetReviewId()}); err != nil {
		return nil, err
	}
	if _, err := s.testcaseClient.GetTestCase(ctx, &v1.GetTestCaseRequest{Id: req.GetCaseId()}); err != nil {
		return nil, err
	}
	if err := s.review.CreateReviewCase(ctx, &biz.TestReviewCase{ReviewId: req.GetReviewId(), CaseId: req.GetCaseId(), Status: req.GetStatus()}); err != nil {
		return nil, err
	}
	return &v1.CreateTestReviewCaseReply{}, nil
}

func (s *ReviewService) UpdateReviewCase(ctx context.Context, req *v1.UpdateTestReviewCaseRequest) (*v1.UpdateTestReviewCaseReply, error) {
	if _, err := s.reviewClient.GetTestReview(ctx, &v1.GetTestReviewRequest{Id: req.GetReviewId()}); err != nil {
		return nil, err
	}
	if _, err := s.testcaseClient.GetTestCase(ctx, &v1.GetTestCaseRequest{Id: req.GetCaseId()}); err != nil {
		return nil, err
	}
	if err := s.review.UpdateReviewCase(ctx, req.GetId(), &biz.TestReviewCase{ReviewId: req.GetReviewId(), CaseId: req.GetCaseId(), Status: req.GetStatus()}); err != nil {
		return nil, err
	}
	return &v1.UpdateTestReviewCaseReply{}, nil
}

func (s *ReviewService) DeleteReviewCase(ctx context.Context, req *v1.DeleteTestReviewCaseRequest) (*v1.DeleteTestReviewCaseReply, error) {
	if err := s.review.DeleteReviewCase(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &v1.DeleteTestReviewCaseReply{}, nil
}

func (s *ReviewService) GetReviewCase(ctx context.Context, req *v1.GetTestReviewCaseRequest) (*v1.GetTestReviewCaseReply, error) {
	testReviewCase, err := s.review.GetReviewCaseById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetTestReviewCaseReply{TestReviewCase: &v1.TestReviewCase{Id: testReviewCase.Id, ReviewId: testReviewCase.ReviewId, CaseId: testReviewCase.CaseId, Status: testReviewCase.Status}}, nil
}

func (s *ReviewService) ListReviewCase(ctx context.Context, req *v1.ListTestReviewCaseRequest) (*v1.ListTestReviewCaseReply, error) {
	parseIndex, err := strconv.ParseInt(req.GetPageToken(), 10, 64)
	if err != nil {
		return nil, err
	}
	testReviewCases, err := s.review.ListReviewCase(ctx, pagination.ListMeta{PageSize: int64(req.GetPageSize()), PageToken: parseIndex})
	if err != nil {
		return nil, err
	}
	testReviewCase := make([]*v1.TestReviewCase, 0, len(testReviewCases.Items))
	for _, item := range testReviewCases.Items {
		testReviewCase = append(testReviewCase, &v1.TestReviewCase{Id: item.Id, ReviewId: item.ReviewId, CaseId: item.CaseId, Status: item.Status})
	}
	return &v1.ListTestReviewCaseReply{Total: testReviewCases.TotalCount, Results: testReviewCase}, nil
}

func (s *ReviewService) GetReviewReport(ctx context.Context, req *v1.GetTestReviewReportRequest) (*v1.GetTestReviewReportReply, error) {
	testReviewReport, err := s.review.GetReviewReport(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetTestReviewReportReply{TestReviewReport: &v1.TestReviewReport{Id: testReviewReport.Id, ProjectId: testReviewReport.ProjectId, Name: testReviewReport.Name}}, nil
}
