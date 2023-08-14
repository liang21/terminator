package service

import (
	"context"
	v1 "github.com/liang21/terminator/api/system/v1"
	"github.com/liang21/terminator/internal/system-service/biz"
	"github.com/liang21/terminator/pkg/pagination"
	"strconv"
)

type ProjectService struct {
	v1.UnimplementedProjectServiceServer
	project *biz.ProjectUsecase
}

func NewProjectService(project *biz.ProjectUsecase) *ProjectService {
	return &ProjectService{project: project}
}

func (u *ProjectService) ListProject(ctx context.Context, req *v1.ListProjectRequest) (*v1.ListProjectReply, error) {
	parseIndex, err := strconv.ParseInt(req.GetPageToken(), 10, 64)
	if err != nil {
		return nil, err
	}
	meta := pagination.ListMeta{
		PageSize:  int64(req.GetPageSize()),
		PageToken: parseIndex,
	}
	projects, err := u.project.List(ctx, meta)
	project := make([]*v1.Project, 0, len(projects.Items))
	for _, item := range projects.Items {
		project = append(project, &v1.Project{Id: item.Id, Name: item.Name, Description: item.Description, Members: item.Members})
	}
	if err != nil {
		return nil, err
	}
	return &v1.ListProjectReply{Total: projects.TotalCount, Results: project}, nil
}

func (u *ProjectService) GetProject(ctx context.Context, req *v1.GetProjectRequest) (*v1.GetProjectReply, error) {
	project, err := u.project.Get(ctx, req.GetId())
	return &v1.GetProjectReply{Project: &v1.Project{Id: project.Id, Name: project.Name, Description: project.Description, Members: project.Members}}, err
}

func (u *ProjectService) CreateProject(ctx context.Context, req *v1.CreateProjectRequest) (*v1.CreateProjectReply, error) {
	err := u.project.Create(ctx, &biz.Project{Name: req.GetName(), Description: req.GetDescription(), ProductId: req.GetProductId()})
	return &v1.CreateProjectReply{}, err
}

func (u *ProjectService) UpdateProject(ctx context.Context, req *v1.UpdateProjectRequest) (*v1.UpdateProjectReply, error) {
	err := u.project.Update(ctx, req.GetId(), &biz.Project{Name: req.GetName(), Description: req.GetDescription()})
	return &v1.UpdateProjectReply{}, err
}

func (u *ProjectService) DeleteProject(ctx context.Context, req *v1.DeleteProjectRequest) (*v1.DeleteProjectReply, error) {
	err := u.project.Delete(ctx, req.GetId())
	return &v1.DeleteProjectReply{}, err
}
