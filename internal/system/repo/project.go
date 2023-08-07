package repo

import (
	"context"
	"errors"
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
	var projects []*biz.Project
	offset := pagination.GetPageOffset(meta.PageSize, meta.PageToken)
	count, err := p.db.Limit(int(meta.PageSize), int(offset)).Desc("name").And("deleted = ?", 0).FindAndCount(&projects)
	if err != nil {
		return nil, err
	}
	projectDtoList := &biz.ProjectDTOList{}
	projectDtoList.TotalCount = count
	projectDtoList.Items = projects
	return projectDtoList, nil
}

func (p *projectRepo) GetProject(ctx context.Context, id int64) (*biz.Project, error) {
	project := &biz.Project{Deleted: 0}
	err := p.db.Where("id = ?", id).Find(project)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p *projectRepo) CreateProject(ctx context.Context, project *biz.Project) error {
	result, err := p.db.Insert(project)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("project create failed")
	}
	return nil
}

func (p *projectRepo) UpdateProject(ctx context.Context, id int64, project *biz.Project) error {
	result, err := p.db.Where("id = ?", id).Update(project)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("project update failed")
	}
	return nil
}

func (p *projectRepo) DeleteProject(ctx context.Context, id int64) error {
	result, err := p.db.Update(&biz.Project{Deleted: 1}, &biz.Project{Id: id})
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("project delete failed")
	}
	return nil
}
