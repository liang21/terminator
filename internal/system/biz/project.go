package biz

import (
	"context"
	"time"
)

type Project struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Members     int64     `json:"members"`
	UserId      []int64   `json:"user_id"`
	ProductId   int64     `json:"product_id"`
	Delete      int64     `json:"delete"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

func (p Project) TableName() string {
	return "project"
}

type ProjectRepo interface {
	// db
	ListProject(ctx context.Context) ([]*Project, error)
	GetProject(ctx context.Context, id int64) (*Project, error)
	CreateProject(ctx context.Context, project *Project) error
	UpdateProject(ctx context.Context, id int64, project *Project) error
	DeleteProject(ctx context.Context, id int64) error
	// redis
	GetProjectLike(ctx context.Context, id int64) (rv int64, err error)
	IncProjectLike(ctx context.Context, id int64) error
}
