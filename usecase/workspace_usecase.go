package usecase

import (
	"shiftlab-go-rest-api/model"
	"shiftlab-go-rest-api/repository"
	"shiftlab-go-rest-api/validator"
)

type IWorkspaceUsecase interface {
	GetAllWorkspaces(userId uint) ([]model.WorkspaceResponse, error)
	GetWorkspaceById(userId uint, workspaceId uint) (model.WorkspaceResponse, error)
	CreateWorkspace(workspace model.Workspace) (model.WorkspaceResponse, error)
	UpdateWorkspace(workspace model.Workspace, userId uint, workspaceId uint) (model.WorkspaceResponse, error)
	DeleteWorkspace(userId uint, workspaceId uint) error
}

type workspaceUsecase struct {
	wr repository.IWorkspaceRepository
	wv validator.IWorkspaceValidator
}

func NewWorkspaceUsecase(wr repository.IWorkspaceRepository, wv validator.IWorkspaceValidator) IWorkspaceUsecase {
	return &workspaceUsecase{wr, wv}
}

func (wu *workspaceUsecase) GetAllWorkspaces(userId uint) ([]model.WorkspaceResponse, error) {
	workspaces := []model.Workspace{}
	if err := wu.wr.GetAllWorkspaces(&workspaces, userId); err != nil {
		return nil, err
	}
	resWorkspaces := []model.WorkspaceResponse{}
	for _, v := range workspaces {
		t := model.WorkspaceResponse{
			ID:            v.ID,
			Name:          v.Name,
			Salary:        v.Salary,
			MorningSalary: v.MorningSalary,
			NightSalary:   v.NightSalary,
			Color:         v.Color,
			CreatedAt:     v.CreatedAt,
			UpdateAt:      v.UpdateAt,
		}
		resWorkspaces = append(resWorkspaces, t)
	}
	return resWorkspaces, nil
}
func (wu *workspaceUsecase) GetWorkspaceById(userId uint, workspaceId uint) (model.WorkspaceResponse, error) {
	workspace := model.Workspace{}
	if err := wu.wr.GetWorkspaceById(workspace, userId, workspaceId); err != nil {
		return model.WorkspaceResponse{}, err
	}
	resWorkspace := model.WorkspaceResponse{
		ID:            workspace.ID,
		Name:          workspace.Name,
		Salary:        workspace.Salary,
		MorningSalary: workspace.MorningSalary,
		NightSalary:   workspace.NightSalary,
		Color:         workspace.Color,
		CreatedAt:     workspace.CreatedAt,
		UpdateAt:      workspace.UpdateAt,
	}
	return resWorkspace, nil
}

func (wu *workspaceUsecase) CreateWorkspace(workspace model.Workspace) (model.WorkspaceResponse, error) {

	if err := wu.wv.WorkspaceValidate(workspace); err != nil {
		return model.WorkspaceResponse{}, err
	}

	if err := wu.wr.CreateWorkspace(&workspace); err != nil {
		return model.WorkspaceResponse{}, err
	}
	resWorkspace := model.WorkspaceResponse{
		ID:            workspace.ID,
		Name:          workspace.Name,
		Salary:        workspace.Salary,
		MorningSalary: workspace.MorningSalary,
		NightSalary:   workspace.NightSalary,
		Color:         workspace.Color,
		CreatedAt:     workspace.CreatedAt,
		UpdateAt:      workspace.UpdateAt,
	}
	return resWorkspace, nil
}

func (wu *workspaceUsecase) UpdateWorkspace(workspace model.Workspace, userId uint, workspaceId uint) (model.WorkspaceResponse, error) {

	if err := wu.wv.WorkspaceValidate(workspace); err != nil {
		return model.WorkspaceResponse{}, err
	}
	if err := wu.wr.UpdateWorkspace(&workspace, userId, workspaceId); err != nil {
		return model.WorkspaceResponse{}, err
	}

	resWorkspace := model.WorkspaceResponse{
		ID:            workspace.ID,
		Name:          workspace.Name,
		Salary:        workspace.Salary,
		MorningSalary: workspace.MorningSalary,
		NightSalary:   workspace.NightSalary,
		Color:         workspace.Color,
		CreatedAt:     workspace.CreatedAt,
		UpdateAt:      workspace.UpdateAt,
	}
	return resWorkspace, nil
}
func (wu *workspaceUsecase) DeleteWorkspace(userId uint, workspaceId uint) error {
	if err := wu.wr.DeleteWorkspace(userId, workspaceId); err != nil {
		return err
	}
	return nil
}
