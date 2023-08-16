package biz

import (
	"context"
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
