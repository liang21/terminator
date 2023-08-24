package service

import (
	"context"
	project_v1 "github.com/liang21/terminator/api/system/v1"
	v1 "github.com/liang21/terminator/api/test/v1"
	"github.com/liang21/terminator/internal/test-service/biz"
	"github.com/liang21/terminator/pkg/pagination"
	"google.golang.org/grpc"
	"strconv"
)

type TestCaseService struct {
	v1.UnimplementedTestCaseServiceServer
	testcase      *biz.TestCaseUsecase
	moduleClient  v1.TestModuleServiceClient
	projectClient project_v1.ProjectServiceClient
}

func NewTestCaseService(cc *grpc.ClientConn, testCase *biz.TestCaseUsecase) *TestCaseService {
	return &TestCaseService{
		testcase:      testCase,
		moduleClient:  v1.NewTestModuleServiceClient(cc),
		projectClient: project_v1.NewProjectServiceClient(cc),
	}
}

func (s *TestCaseService) CreateTestCase(ctx context.Context, req *v1.CreateTestCaseRequest) (*v1.CreateTestCaseReply, error) {
	if _, err := s.projectClient.GetProject(ctx, &project_v1.GetProjectRequest{Id: req.GetProjectId()}); err != nil {
		return nil, err
	}
	if _, err := s.moduleClient.GetTestModule(ctx, &v1.GetTestModuleRequest{Id: req.GetModuleId()}); err != nil {
		return nil, err
	}
	if err := s.testcase.Create(ctx,
		&biz.TestCase{Name: req.GetName(), ModuleId: req.GetModuleId(), ProjectId: req.GetProjectId(), Type: req.GetType(),
			Maintainer: req.GetMaintainer(), Priority: req.GetPriority(), Method: req.GetMethod(), Prerequisite: req.GetPrerequisite(),
			Steps: req.GetSteps(), Remark: req.GetRemark(), Status: req.GetStatus(), StepDescription: req.GetStepDescription(),
			StepExpectedResult: req.GetStepExpectedResult(), CreateUser: req.CreateUser, OriginalStatus: req.GetOriginalStatus(), LastExecuteResult: req.GetLastExecutedResult()}); err != nil {
		return nil, err
	}
	return &v1.CreateTestCaseReply{}, nil
}

func (s *TestCaseService) UpdateTestCase(ctx context.Context, req *v1.UpdateTestCaseRequest) (*v1.UpdateTestCaseReply, error) {
	if _, err := s.projectClient.GetProject(ctx, &project_v1.GetProjectRequest{Id: req.GetProjectId()}); err != nil {
		return nil, err
	}
	if _, err := s.moduleClient.GetTestModule(ctx, &v1.GetTestModuleRequest{Id: req.GetModuleId()}); err != nil {
		return nil, err
	}
	if err := s.testcase.Update(ctx, req.GetId(),
		&biz.TestCase{Name: req.GetName(), ModuleId: req.GetModuleId(), ProjectId: req.GetProjectId(), Type: req.GetType(),
			Maintainer: req.GetMaintainer(), Priority: req.GetPriority(), Method: req.GetMethod(), Prerequisite: req.GetPrerequisite(),
			Steps: req.GetSteps(), Remark: req.GetRemark(), Status: req.GetStatus(), StepDescription: req.GetStepDescription(),
			StepExpectedResult: req.GetStepExpectedResult(), CreateUser: req.CreateUser, OriginalStatus: req.GetOriginalStatus(), LastExecuteResult: req.GetLastExecutedResult()}); err != nil {
		return nil, err
	}
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
	return &v1.GetTestCaseReply{TestCase: &v1.TestCase{Id: testcase.Id, ModuleId: testcase.ModuleId, ProjectId: testcase.ProjectId, Name: testcase.Name,
		Type: testcase.Type, Maintainer: testcase.Maintainer, Priority: testcase.Priority, Method: testcase.Method, Prerequisite: testcase.Prerequisite,
		Remark: testcase.Remark, Steps: testcase.Steps, ReviewStatus: testcase.ReviewStatus, Status: testcase.Status, StepDescription: testcase.StepDescription,
		StepExpectedResult: testcase.StepExpectedResult, CreateUser: testcase.CreateUser, OriginalStatus: testcase.OriginalStatus, LastExecutedResult: testcase.LastExecuteResult,
	}}, nil
}

func (s *TestCaseService) ListTestCase(ctx context.Context, req *v1.ListTestCaseRequest) (*v1.ListTestCaseReply, error) {
	parseIndex, err := strconv.ParseInt(req.GetPageToken(), 10, 64)
	if err != nil {
		return nil, err
	}
	testcases, err := s.testcase.List(ctx, pagination.ListMeta{
		PageSize:  int64(req.GetPageSize()),
		PageToken: parseIndex,
	})
	if err != nil {
		return nil, err
	}
	testcase := make([]*v1.TestCase, 0, len(testcases.Items))
	for _, item := range testcases.Items {
		testcase = append(testcase, &v1.TestCase{Id: item.Id, ModuleId: item.ModuleId, ProjectId: item.ProjectId, Name: item.Name,
			Type: item.Type, Maintainer: item.Maintainer, Priority: item.Priority, Method: item.Method, Prerequisite: item.Prerequisite,
			Remark: item.Remark, Steps: item.Steps, ReviewStatus: item.ReviewStatus, Status: item.Status, StepDescription: item.StepDescription,
			StepExpectedResult: item.StepExpectedResult, CreateUser: item.CreateUser, OriginalStatus: item.OriginalStatus, LastExecutedResult: item.LastExecuteResult,
		})
	}
	return &v1.ListTestCaseReply{Total: testcases.TotalCount, Results: testcase}, nil
}
