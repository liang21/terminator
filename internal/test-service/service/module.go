package service

import (
	"context"
	v1 "github.com/liang21/terminator/api/test/v1"
	"github.com/liang21/terminator/internal/test-service/biz"
)

type TestModuleService struct {
	v1.UnimplementedTestModuleServiceServer
	module *biz.TestModuleUsecase
}

func NewTestModuleService(module *biz.TestModuleUsecase) *TestModuleService {
	return &TestModuleService{module: module}
}

func (s *TestModuleService) CreateTestModule(ctx context.Context, req *v1.CreateTestModuleRequest) (*v1.CreateTestModuleReply, error) {
	return &v1.CreateTestModuleReply{}, nil
}
func (s *TestModuleService) UpdateTestModule(ctx context.Context, req *v1.UpdateTestModuleRequest) (*v1.UpdateTestModuleReply, error) {
	return &v1.UpdateTestModuleReply{}, nil
}

func (s *TestModuleService) DeleteTestModule(ctx context.Context, req *v1.DeleteTestModuleRequest) (*v1.DeleteTestModuleReply, error) {
	return &v1.DeleteTestModuleReply{}, nil
}

func (s *TestModuleService) GetTestModule(ctx context.Context, req *v1.GetTestModuleRequest) (*v1.GetTestModuleReply, error) {
	return &v1.GetTestModuleReply{}, nil
}

func (s *TestModuleService) ListTestModule(ctx context.Context, req *v1.ListTestModuleRequest) (*v1.ListTestModuleReply, error) {
	return &v1.ListTestModuleReply{}, nil
}
