package service

import (
	"context"
	v1 "github.com/liang21/terminator/api/test/v1"
	"github.com/liang21/terminator/internal/test-service/biz"
)

type PlanService struct {
	v1.UnimplementedPlanServiceServer
	v1.UnimplementedTestPlanCaseServiceServer
	v1.UnimplementedTestPlanReportServiceServer
	plan *biz.PlanUsecase
}

func NewPlanService(plan *biz.PlanUsecase) *PlanService {
	return &PlanService{plan: plan}
}

func (s *PlanService) CreatePlan(ctx context.Context, req *v1.CreateTestPlanRequest) (*v1.CreateTestPlanReply, error) {
	return &v1.CreateTestPlanReply{}, nil
}

func (s *PlanService) UpdatePlan(ctx context.Context, req *v1.UpdateTestPlanRequest) (*v1.UpdateTestPlanReply, error) {
	return &v1.UpdateTestPlanReply{}, nil
}

func (s *PlanService) DeletePlan(ctx context.Context, req *v1.DeleteTestPlanRequest) (*v1.DeleteTestPlanReply, error) {
	return &v1.DeleteTestPlanReply{}, nil
}

func (s *PlanService) GetPlan(ctx context.Context, req *v1.GetTestPlanRequest) (*v1.GetTestPlanReply, error) {
	return &v1.GetTestPlanReply{}, nil
}

func (s *PlanService) ListPlan(ctx context.Context, req *v1.ListTestPlanRequest) (*v1.ListTestPlanReply, error) {
	return &v1.ListTestPlanReply{}, nil
}

func (s *PlanService) CreatePlanCase(ctx context.Context, req *v1.CreateTestPlanCaseRequest) (*v1.CreateTestPlanCaseReply, error) {
	return &v1.CreateTestPlanCaseReply{}, nil
}

func (s *PlanService) UpdatePlanCase(ctx context.Context, req *v1.UpdateTestPlanCaseRequest) (*v1.UpdateTestPlanCaseReply, error) {
	return &v1.UpdateTestPlanCaseReply{}, nil
}

func (s *PlanService) DeletePlanCase(ctx context.Context, req *v1.DeleteTestPlanCaseRequest) (*v1.DeleteTestPlanCaseReply, error) {
	return &v1.DeleteTestPlanCaseReply{}, nil
}

func (s *PlanService) GetPlanCase(ctx context.Context, req *v1.GetTestPlanCaseRequest) (*v1.GetTestPlanCaseReply, error) {
	return &v1.GetTestPlanCaseReply{}, nil
}

func (s *PlanService) ListPlanCase(ctx context.Context, req *v1.ListTestPlanCaseRequest) (*v1.ListTestPlanCaseReply, error) {
	return &v1.ListTestPlanCaseReply{}, nil
}

func (s *PlanService) GetPlanReport(ctx context.Context, req *v1.GetTestPlanReportRequest) (*v1.GetTestPlanReportReply, error) {
	return &v1.GetTestPlanReportReply{}, nil
}
