package repository

import (
	"fmt"
	"shiftlab-go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IShiftRepository interface {
	GetAllShifts(shifts *[]model.Shift, userId uint) error
	GetShiftByWorkspaceId(shift *[]model.Shift, workspaceId uint) error
	CreateShift(shift *model.Shift) error
	UpdateShift(shift *model.Shift, shiftId uint) error
	DeleteShift(shiftId uint) error
}

type shiftRepository struct {
	db *gorm.DB
}

func NewShiftRepository(db *gorm.DB) IShiftRepository {
	return &shiftRepository{db}
}

func (sr shiftRepository) GetAllShifts(shifts *[]model.Shift, userId uint) error {
	if err := sr.db.Joins("Workspace").Where("user_id=?", userId).Order("created_at").Find(shifts).Error; err != nil {
		return err
	}
	return nil
}

func (sr shiftRepository) GetShiftByWorkspaceId(shifts *[]model.Shift, workspaceId uint) error {
	if err := sr.db.Where("workspace_id=?", workspaceId).Order("created_at").Find(shifts).Error; err != nil {
		return err
	}
	return nil
}

func (sr shiftRepository) CreateShift(shift *model.Shift) error {
	if err := sr.db.Create(shift).Error; err != nil {
		return err
	}
	return nil
}

func (sr *shiftRepository) UpdateShift(shift *model.Shift, shiftId uint) error {
	result := sr.db.Model(shift).Clauses(clause.Returning{}).Where("id=?", shiftId).Updates(shift)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (sr *shiftRepository) DeleteShift(shiftId uint) error {
	result := sr.db.Where("id=?", shiftId).Delete(&model.Shift{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
