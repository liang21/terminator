package biz

import (
	"context"
	"errors"
	"github.com/liang21/terminator/pkg/pagination"
	"time"
)

type TestModule struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	ProjectId int64     `json:"project_id"`
	Deleted   int64     `json:"deleted"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}
type TestModuleDTOList struct {
	TotalCount int64         `json:"totalCount,omitempty"` //总数
	Items      []*TestModule `json:"data"`                 //数据
}

func (m TestModule) TableName() string {
	return "test_module"
}

type TestModuleRepo interface {
	// ListTestModule db
	ListTestModule(ctx context.Context, meta pagination.ListMeta) (TestModuleDto *TestModuleDTOList, err error)
	GetTestModule(ctx context.Context, id int64) (*TestModule, error)
	CreateTestModule(ctx context.Context, TestModule *TestModule) error
	UpdateTestModule(ctx context.Context, id int64, TestModule *TestModule) error
	DeleteTestModule(ctx context.Context, id int64) error
}

type TestModuleUsecase struct {
	repo TestModuleRepo
}

func NewTestModuleUsecase(repo TestModuleRepo) *TestModuleUsecase {
	return &TestModuleUsecase{repo: repo}
}

func (uc *TestModuleUsecase) List(ctx context.Context, meta pagination.ListMeta) (module *TestModuleDTOList, err error) {
	module, err = uc.repo.ListTestModule(ctx, meta)
	if err != nil {
		return
	}
	return module, nil
}

func (uc *TestModuleUsecase) Get(ctx context.Context, id int64) (module *TestModule, err error) {
	if id == 0 {
		return nil, errors.New("id is empty")
	}
	module, err = uc.repo.GetTestModule(ctx, id)
	if err != nil {
		return
	}
	return module, nil
}

func (uc *TestModuleUsecase) Create(ctx context.Context, module *TestModule) (err error) {
	if err = uc.repo.CreateTestModule(ctx, module); err != nil {
		return err
	}
	return
}

func (uc *TestModuleUsecase) Update(ctx context.Context, id int64, module *TestModule) (err error) {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err = uc.repo.UpdateTestModule(ctx, id, module); err != nil {
		return err
	}
	return
}

func (uc *TestModuleUsecase) Delete(ctx context.Context, id int64) (err error) {
	if id == 0 {
		return errors.New("id is empty")
	}
	if err = uc.repo.DeleteTestModule(ctx, id); err != nil {
		return err
	}
	return
}
