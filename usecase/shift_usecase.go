package usecase

import (
	"shiftlab-go-rest-api/model"
	"shiftlab-go-rest-api/repository"
)

type IShiftUsecase interface {
	GetAllShifts(userId uint) ([]model.ShiftResponse, error)
	GetShiftByWorkspaceId(userId uint, workspaceId uint) ([]model.ShiftResponse, error)
	CreateShift(shift model.Shift) (model.ShiftResponse, error)
	UpdateShift(shift model.Shift, userId uint, shiftId uint) (model.ShiftResponse, error)
	DeleteShift(userId uint, shiftId uint) error
}

type shiftUsecase struct {
	sr repository.IShiftRepository
}

func NewShiftUsecase(sr repository.IShiftRepository) IShiftUsecase {
	return &shiftUsecase{sr}
}

func (su *shiftUsecase) GetAllShifts(userId uint) ([]model.ShiftResponse, error) {
	shifts := []model.Shift{}
	if err := su.sr.GetAllShifts(&shifts, userId); err != nil {
		return nil, err
	}
	resShifts := []model.ShiftResponse{}
	for _, v := range shifts {
		s := model.ShiftResponse{
			ID:          v.ID,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
			WorkspaceId: v.WorkspaceId,
			CreatedAt:   v.CreatedAt,
			UpdateAt:    v.CreatedAt,
		}
		resShifts = append(resShifts, s)
	}
	return resShifts, nil
}

func (su *shiftUsecase) GetShiftByWorkspaceId(userId uint, workspaceId uint) ([]model.ShiftResponse, error) {
	shifts := []model.Shift{}
	if err := su.sr.GetShiftByWorkspaceId(&shifts, userId, workspaceId); err != nil {
		return nil, err
	}
	resShifts := []model.ShiftResponse{}
	for _, v := range shifts {
		s := model.ShiftResponse{
			ID:          v.ID,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
			WorkspaceId: v.WorkspaceId,
			CreatedAt:   v.CreatedAt,
			UpdateAt:    v.CreatedAt,
		}
		resShifts = append(resShifts, s)
	}
	return resShifts, nil
}

func (su *shiftUsecase) CreateShift(shift model.Shift) (model.ShiftResponse, error) {

	if err := su.sr.CreateShift(&shift); err != nil {
		return model.ShiftResponse{}, err
	}
	resShift := model.ShiftResponse{
		ID:          shift.ID,
		StartTime:   shift.StartTime,
		EndTime:     shift.EndTime,
		WorkspaceId: shift.WorkspaceId,
		CreatedAt:   shift.CreatedAt,
		UpdateAt:    shift.CreatedAt,
	}
	return resShift, nil
}

func (su *shiftUsecase) UpdateShift(shift model.Shift, userId uint, shiftId uint) (model.ShiftResponse, error) {
	if err := su.sr.UpdateShift(&shift, userId, shiftId); err != nil {
		return model.ShiftResponse{}, err
	}
	resShift := model.ShiftResponse{
		ID:          shift.ID,
		StartTime:   shift.StartTime,
		EndTime:     shift.EndTime,
		WorkspaceId: shift.WorkspaceId,
		CreatedAt:   shift.CreatedAt,
		UpdateAt:    shift.CreatedAt,
	}
	return resShift, nil
}

func (su *shiftUsecase) DeleteShift(userId uint, shiftId uint) error {
	if err := su.sr.DeleteShift(userId, shiftId); err != nil {
		return err
	}
	return nil
}
