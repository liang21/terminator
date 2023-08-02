package repo

import (
	"context"
	"github.com/liang21/terminator/internal/system/biz"
	"github.com/liang21/terminator/pkg/pagination"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type projectRepo struct {
	db  *xorm.Engine
	rdb *redis.Client
}

func NewProjectRepo(db *xorm.Engine, rdb *redis.Client) biz.ProjectRepo {
	return &projectRepo{
		db:  db,
		rdb: rdb,
	}
}

func (p *projectRepo) ListProject(ctx context.Context, meta pagination.ListMeta) (*biz.ProjectDTOList, error) {
	//TODO implement me
	panic("implement me")
}

func (p *projectRepo) GetProject(ctx context.Context, id int64) (*biz.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p *projectRepo) CreateProject(ctx context.Context, project *biz.Project) error {
	//TODO implement me
	panic("implement me")
}

func (p *projectRepo) UpdateProject(ctx context.Context, id int64, project *biz.Project) error {
	//TODO implement me
	panic("implement me")
}

func (p *projectRepo) DeleteProject(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
