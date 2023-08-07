package biz

import (
	"context"
	"errors"
	"github.com/liang21/terminator/pkg/pagination"
	"time"
)

type Project struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Members     int64     `json:"members"`
	UserId      int64     `json:"user_id"`
	ProductId   int64     `json:"product_id"`
	Deleted     int64     `json:"deleted"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

func (p Project) TableName() string {
	return "project"
}

type ProjectRepo interface {
	// db
	ListProject(ctx context.Context, meta pagination.ListMeta) (*ProjectDTOList, error)
	GetProject(ctx context.Context, id int64) (*Project, error)
	CreateProject(ctx context.Context, project *Project) error
	UpdateProject(ctx context.Context, id int64, project *Project) error
	DeleteProject(ctx context.Context, id int64) error
}

type ProjectUsecase struct {
	repo ProjectRepo
}

func NewProjectUsecase(repo ProjectRepo) *ProjectUsecase {
	return &ProjectUsecase{repo: repo}
}

type ProjectDTOList struct {
	TotalCount int64      `json:"total"` //总数
	Items      []*Project `json:"items"` //数据
}

func (pu *ProjectUsecase) List(ctx context.Context, meta pagination.ListMeta) (projectDtoList *ProjectDTOList, err error) {
	projectDtoList, err = pu.repo.ListProject(ctx, meta)
	if err != nil {
		return
	}
	return projectDtoList, nil
}

func (pu *ProjectUsecase) Get(ctx context.Context, id int64) (project *Project, err error) {
	if id == 0 {
		return nil, errors.New("id is empty")
	}
	project, err = pu.repo.GetProject(ctx, id)
	if err != nil {
		return
	}
	return project, nil
}

func (pu *ProjectUsecase) Create(ctx context.Context, project *Project) error {
	return pu.repo.CreateProject(ctx, project)
}

func (pu *ProjectUsecase) Update(ctx context.Context, id int64, project *Project) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	return pu.repo.UpdateProject(ctx, id, project)
}

func (pu *ProjectUsecase) Delete(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	return pu.repo.DeleteProject(ctx, id)
}
