package repository

import (
	"fmt"
	"shiftlab-go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IWorkspaceRepository interface {
	GetAllWorkspaces(workspaces *[]model.Workspace, userId uint) error
	GetWorkspaceById(workspace model.Workspace, userId uint, workspaceId uint) error
	CreateWorkspace(workspace *model.Workspace) error
	UpdateWorkspace(workspace *model.Workspace, userId uint, workspaceId uint) error
	DeleteWorkspace(userId uint, workspaceId uint) error
}

type workspaceRepository struct {
	db *gorm.DB
}

func NewWorkspaceRepository(db *gorm.DB) IWorkspaceRepository {
	return &workspaceRepository{db}
}

func (wr *workspaceRepository) GetAllWorkspaces(workspaces *[]model.Workspace, userId uint) error {
	if err := wr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(workspaces).Error; err != nil {
		return err
	}
	return nil
}

func (wr *workspaceRepository) GetWorkspaceById(workspace model.Workspace, userId uint, workspaceId uint) error {
	if err := wr.db.Joins("User").Where("user_id=?", userId).First(workspace, workspaceId).Error; err != nil {
		return err
	}
	return nil
}

func (wr *workspaceRepository) CreateWorkspace(workspace *model.Workspace) error {
	if err := wr.db.Create(workspace).Error; err != nil {
		return err
	}
	return nil
}

func (wr *workspaceRepository) UpdateWorkspace(workspace *model.Workspace, userId uint, workspaceId uint) error {
	result := wr.db.Model(workspace).Clauses(clause.Returning{}).Where("id=? AND user_id=?", workspaceId, userId).Update("color", workspace.Color)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (wr *workspaceRepository) DeleteWorkspace(userId uint, workspaceId uint) error {
	result := wr.db.Where("id=? AND user_id=?", workspaceId, userId).Delete(&model.Workspace{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
