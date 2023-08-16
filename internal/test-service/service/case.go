package service

import (
	"context"
	project_v1 "github.com/liang21/terminator/api/system/v1"
	v1 "github.com/liang21/terminator/api/test/v1"
	"github.com/liang21/terminator/internal/test-service/biz"
)

type TestCaseService struct {
	v1.UnimplementedTestCaseServiceServer
	testcase      *biz.TestCaseUsecase
	moduleClient  *v1.TestModuleServiceClient
	projectClient *project_v1.ProjectServiceClient
}

func NewTestCaseService(testCase *biz.TestCaseUsecase) *TestCaseService {
	return &TestCaseService{testcase: testCase}
}

func (s *TestCaseService) CreateTestCase(ctx context.Context, req *v1.CreateTestCaseRequest) (*v1.CreateTestCaseReply, error) {
	if err := s.testcase.Create(ctx,
		&biz.TestCase{Name: req.GetName(), ModuleId: req.GetModuleId()}); err != nil {
		return nil, err
	}
	return &v1.CreateTestCaseReply{}, nil
}

func (s *TestCaseService) UpdateTestCase(ctx context.Context, req *v1.UpdateTestCaseRequest) (*v1.UpdateTestCaseReply, error) {
	return &v1.UpdateTestCaseReply{}, nil
}

func (s *TestCaseService) DeleteTestCase(ctx context.Context, req *v1.DeleteTestCaseRequest) (*v1.DeleteTestCaseReply, error) {
	if err := s.testcase.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &v1.DeleteTestCaseReply{}, nil
}

func (s *TestCaseService) GetTestCase(ctx context.Context, req *v1.GetTestCaseRequest) (*v1.GetTestCaseReply, error) {
	testcase, err := s.testcase.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetTestCaseReply{TestCase: &v1.TestCase{Id: testcase.Id, ModuleId: testcase.ModuleId, ProjectId: testcase.ProjectId, Name: testcase.Name, Type: testcase.Type,
		Maintainer: testcase.Maintainer, Priority: testcase.Priority, Method: testcase.Method, Prerequisite: testcase.Prerequisite,
		Remark: testcase.Remark, Steps: testcase.Steps, ReviewStatus: testcase.ReviewStatus, Status: testcase.Status,
		StepDescription: testcase.StepDescription, StepExpectedResult: testcase.StepExpectedResult, CreateUser: testcase.CreateUser, OriginalStatus: testcase.OriginalStatus, LastExecutedResult: testcase.LastExecuteResult}}, nil
}

func (s *TestCaseService) ListTestCase(ctx context.Context, req *v1.ListTestCaseRequest) (*v1.ListTestCaseReply, error) {
	return &v1.ListTestCaseReply{}, nil
}
