package service

import (
	"context"
	v1 "github.com/liang21/terminator/api/test/v1"
	"github.com/liang21/terminator/internal/test-service/biz"
)

type TestCaseService struct {
	v1.UnimplementedTestCaseServiceServer
	testCase *biz.TestCaseUsecase
}

func NewTestCaseService(testCase *biz.TestCaseUsecase) *TestCaseService {
	return &TestCaseService{testCase: testCase}
}

func (s *TestCaseService) CreateTestCase(ctx context.Context, req *v1.CreateTestCaseRequest) (*v1.CreateTestCaseReply, error) {
	return &v1.CreateTestCaseReply{}, nil
}

func (s *TestCaseService) UpdateTestCase(ctx context.Context, req *v1.UpdateTestCaseRequest) (*v1.UpdateTestCaseReply, error) {
	return &v1.UpdateTestCaseReply{}, nil
}

func (s *TestCaseService) DeleteTestCase(ctx context.Context, req *v1.DeleteTestCaseRequest) (*v1.DeleteTestCaseReply, error) {
	return &v1.DeleteTestCaseReply{}, nil
}

func (s *TestCaseService) GetTestCase(ctx context.Context, req *v1.GetTestCaseRequest) (*v1.GetTestCaseReply, error) {
	return &v1.GetTestCaseReply{}, nil
}

func (s *TestCaseService) ListTestCase(ctx context.Context, req *v1.ListTestCaseRequest) (*v1.ListTestCaseReply, error) {
	return &v1.ListTestCaseReply{}, nil
}
