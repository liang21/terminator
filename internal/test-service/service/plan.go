package service

import (
	"context"
	project_v1 "github.com/liang21/terminator/api/system/v1"
	v1 "github.com/liang21/terminator/api/test/v1"
	"github.com/liang21/terminator/internal/test-service/biz"
	"github.com/liang21/terminator/pkg/pagination"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

type PlanService struct {
	v1.UnimplementedPlanServiceServer
	v1.UnimplementedTestPlanCaseServiceServer
	v1.UnimplementedTestPlanReportServiceServer
	plan           *biz.PlanUsecase
	projectClient  project_v1.ProjectServiceClient
	planClient     v1.PlanServiceClient
	testcaseClient v1.TestPlanCaseServiceClient
}

func NewPlanService(cc *grpc.ClientConn, plan *biz.PlanUsecase) *PlanService {
	return &PlanService{plan: plan}
}

func (s *PlanService) CreatePlan(ctx context.Context, req *v1.CreateTestPlanRequest) (*v1.CreateTestPlanReply, error) {
	if _, err := s.projectClient.GetProject(ctx, &project_v1.GetProjectRequest{Id: req.GetProjectId()}); err != nil {
		return nil, err
	}
	if err := s.plan.CreatePlan(ctx, &biz.TestPlan{Name: req.GetName(), ProjectId: req.GetProjectId(), Status: req.GetStatus(), Stage: req.GetStage(), Description: req.GetDescription(), PlanedStartAt: req.GetPlannedStartAt().AsTime(), PlanedEndAt: req.GetPlannedEndAt().AsTime(), ActualStartAt: req.GetActualStartAt().AsTime(), ActualEndAt: req.GetActualEndAt().AsTime()}); err != nil {
		return nil, err
	}
	return &v1.CreateTestPlanReply{}, nil
}

func (s *PlanService) UpdatePlan(ctx context.Context, req *v1.UpdateTestPlanRequest) (*v1.UpdateTestPlanReply, error) {
	if _, err := s.projectClient.GetProject(ctx, &project_v1.GetProjectRequest{Id: req.GetProjectId()}); err != nil {
		return nil, err
	}
	if err := s.plan.UpdatePlan(ctx, req.GetId(), &biz.TestPlan{Name: req.GetName(), ProjectId: req.GetProjectId(), Status: req.GetStatus(), Stage: req.GetStage(), Description: req.GetDescription(), PlanedStartAt: req.GetPlannedStartAt().AsTime(), PlanedEndAt: req.GetPlannedEndAt().AsTime(), ActualStartAt: req.GetActualStartAt().AsTime(), ActualEndAt: req.GetActualEndAt().AsTime()}); err != nil {
		return nil, err
	}
	return &v1.UpdateTestPlanReply{}, nil
}

func (s *PlanService) DeletePlan(ctx context.Context, req *v1.DeleteTestPlanRequest) (*v1.DeleteTestPlanReply, error) {
	if err := s.plan.DeletePlan(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &v1.DeleteTestPlanReply{}, nil
}

func (s *PlanService) GetPlan(ctx context.Context, req *v1.GetTestPlanRequest) (*v1.GetTestPlanReply, error) {
	testPlan, err := s.plan.GetPlanById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetTestPlanReply{TestPlan: &v1.TestPlan{Id: testPlan.Id, Name: testPlan.Name, ProjectId: testPlan.ProjectId, Status: testPlan.Status, Stage: testPlan.Stage, Description: testPlan.Description, ReviewTotal: testPlan.ReviewTotal, ReviewPassRule: testPlan.ReviewPassRule, PlannedStartAt: timestamppb.New(testPlan.PlanedStartAt.Local()), PlannedEndAt: timestamppb.New(testPlan.PlanedEndAt.Local()), ActualStartAt: timestamppb.New(testPlan.ActualStartAt.Local()), ActualEndAt: timestamppb.New(testPlan.ActualEndAt.Local())}}, nil
}

func (s *PlanService) ListPlan(ctx context.Context, req *v1.ListTestPlanRequest) (*v1.ListTestPlanReply, error) {
	parseIndex, err := strconv.ParseInt(req.GetPageToken(), 10, 64)
	if err != nil {
		return nil, err
	}
	testPlanLists, err := s.plan.ListPlan(ctx, pagination.ListMeta{PageSize: int64(req.GetPageSize()), PageToken: parseIndex})
	if err != nil {
		return nil, err
	}
	testPlanList := make([]*v1.TestPlan, 0, len(testPlanLists.Items))
	for _, item := range testPlanLists.Items {
		testPlanList = append(testPlanList, &v1.TestPlan{Id: item.Id, Name: item.Name, ProjectId: item.ProjectId, Status: item.Status, Stage: item.Stage, Description: item.Description, ReviewTotal: item.ReviewTotal, ReviewPassRule: item.ReviewPassRule, PlannedStartAt: timestamppb.New(item.PlanedStartAt.Local()), PlannedEndAt: timestamppb.New(item.PlanedEndAt.Local()), ActualStartAt: timestamppb.New(item.ActualStartAt.Local()), ActualEndAt: timestamppb.New(item.ActualEndAt.Local())})
	}
	return &v1.ListTestPlanReply{Total: testPlanLists.TotalCount, Results: testPlanList}, nil
}

func (s *PlanService) CreatePlanCase(ctx context.Context, req *v1.CreateTestPlanCaseRequest) (*v1.CreateTestPlanCaseReply, error) {
	if _, err := s.testcaseClient.GetTestPlanCase(ctx, &v1.GetTestPlanCaseRequest{Id: req.GetCaseId()}); err != nil {
		return nil, err
	}
	if _, err := s.planClient.GetTestPlan(ctx, &v1.GetTestPlanRequest{Id: req.GetPlanId()}); err != nil {
		return nil, err
	}

	if err := s.plan.CreatePlanCase(ctx, &biz.TestPlanCase{PlanId: req.GetPlanId(), CaseId: req.GetCaseId(), Status: req.GetStatus(), Result: req.GetResult()}); err != nil {
		return nil, err
	}
	return &v1.CreateTestPlanCaseReply{}, nil
}

func (s *PlanService) UpdatePlanCase(ctx context.Context, req *v1.UpdateTestPlanCaseRequest) (*v1.UpdateTestPlanCaseReply, error) {
	if _, err := s.testcaseClient.GetTestPlanCase(ctx, &v1.GetTestPlanCaseRequest{Id: req.GetCaseId()}); err != nil {
		return nil, err
	}
	if _, err := s.planClient.GetTestPlan(ctx, &v1.GetTestPlanRequest{Id: req.GetPlanId()}); err != nil {
		return nil, err
	}
	if err := s.plan.UpdatePlanCase(ctx, req.GetId(), &biz.TestPlanCase{PlanId: req.GetPlanId(), CaseId: req.GetCaseId(), Status: req.GetStatus(), Result: req.GetResult()}); err != nil {
		return nil, err
	}
	return &v1.UpdateTestPlanCaseReply{}, nil
}

func (s *PlanService) DeletePlanCase(ctx context.Context, req *v1.DeleteTestPlanCaseRequest) (*v1.DeleteTestPlanCaseReply, error) {
	if err := s.plan.DeletePlanCase(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &v1.DeleteTestPlanCaseReply{}, nil
}

func (s *PlanService) GetPlanCase(ctx context.Context, req *v1.GetTestPlanCaseRequest) (*v1.GetTestPlanCaseReply, error) {
	testPlanCase, err := s.plan.GetPlanCaseById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetTestPlanCaseReply{TestPlanCase: &v1.TestPlanCase{Id: testPlanCase.Id, PlanId: testPlanCase.PlanId, CaseId: testPlanCase.CaseId, Status: testPlanCase.Status, Result: testPlanCase.Result}}, nil
}

func (s *PlanService) ListPlanCase(ctx context.Context, req *v1.ListTestPlanCaseRequest) (*v1.ListTestPlanCaseReply, error) {
	parseIndex, err := strconv.ParseInt(req.GetPageToken(), 10, 64)
	if err != nil {
		return nil, err
	}
	testPlanCaseLists, err := s.plan.ListPlanCase(ctx, pagination.ListMeta{PageSize: int64(req.GetPageSize()), PageToken: parseIndex})
	if err != nil {
		return nil, err
	}
	testPlanCaseList := make([]*v1.TestPlanCase, 0, len(testPlanCaseLists.Items))
	for _, item := range testPlanCaseLists.Items {
		testPlanCaseList = append(testPlanCaseList, &v1.TestPlanCase{Id: item.Id, PlanId: item.PlanId, CaseId: item.CaseId, Status: item.Status, Result: item.Result})
	}
	return &v1.ListTestPlanCaseReply{Total: testPlanCaseLists.TotalCount, Results: testPlanCaseList}, nil
}

func (s *PlanService) GetPlanReport(ctx context.Context, req *v1.GetTestPlanReportRequest) (*v1.GetTestPlanReportReply, error) {
	return &v1.GetTestPlanReportReply{}, nil
}
