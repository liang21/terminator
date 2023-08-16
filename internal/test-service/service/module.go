package service

import (
	"context"
	"errors"
	project_v1 "github.com/liang21/terminator/api/system/v1"
	v1 "github.com/liang21/terminator/api/test/v1"
	"github.com/liang21/terminator/internal/test-service/biz"
	"github.com/liang21/terminator/pkg/pagination"
	"google.golang.org/grpc"
	"strconv"
)

type TestModuleService struct {
	v1.UnimplementedTestModuleServiceServer
	projectClient project_v1.ProjectServiceClient
	module        *biz.TestModuleUsecase
}

func NewTestModuleService(cc *grpc.ClientConn, module *biz.TestModuleUsecase) *TestModuleService {
	return &TestModuleService{module: module, projectClient: project_v1.NewProjectServiceClient(cc)}
}

func (s *TestModuleService) CreateTestModule(ctx context.Context, req *v1.CreateTestModuleRequest) (*v1.CreateTestModuleReply, error) {
	_, err := s.projectClient.GetProject(ctx, &project_v1.GetProjectRequest{Id: req.GetProjectId()})
	if err != nil {
		return nil, errors.New("project not found")
	}
	if err := s.module.Create(ctx, &biz.TestModule{Name: req.GetName(), ProjectId: req.ProjectId}); err != nil {
		return nil, err
	}
	return &v1.CreateTestModuleReply{}, nil
}
func (s *TestModuleService) UpdateTestModule(ctx context.Context, req *v1.UpdateTestModuleRequest) (*v1.UpdateTestModuleReply, error) {
	if err := s.module.Update(ctx, req.GetId(), &biz.TestModule{Name: req.GetName(), ProjectId: req.ProjectId}); err != nil {
		return nil, err
	}
	return &v1.UpdateTestModuleReply{}, nil
}

func (s *TestModuleService) DeleteTestModule(ctx context.Context, req *v1.DeleteTestModuleRequest) (*v1.DeleteTestModuleReply, error) {
	if err := s.module.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &v1.DeleteTestModuleReply{}, nil
}

func (s *TestModuleService) GetTestModule(ctx context.Context, req *v1.GetTestModuleRequest) (*v1.GetTestModuleReply, error) {
	module, err := s.module.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetTestModuleReply{TestModule: &v1.TestModule{Id: module.Id, Name: module.Name, ProjectId: module.ProjectId}}, nil
}

func (s *TestModuleService) ListTestModule(ctx context.Context, req *v1.ListTestModuleRequest) (*v1.ListTestModuleReply, error) {
	parseIndex, err := strconv.ParseInt(req.GetPageToken(), 10, 64)
	if err != nil {
		return nil, err
	}
	modules, err := s.module.List(ctx, pagination.ListMeta{
		PageSize:  int64(req.GetPageSize()),
		PageToken: parseIndex,
	})
	if err != nil {
		return nil, err
	}
	module := make([]*v1.TestModule, 0, len(modules.Items))
	for _, item := range modules.Items {
		module = append(module, &v1.TestModule{Id: item.Id, Name: item.Name, ProjectId: item.ProjectId})
	}
	return &v1.ListTestModuleReply{Total: modules.TotalCount, Results: module}, nil
}
