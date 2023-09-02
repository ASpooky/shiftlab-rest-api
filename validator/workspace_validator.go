package validator

import (
	"shiftlab-go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IWorkspaceValidator interface {
	WorkspaceValidate(task model.Workspace) error
}

type workspaceValidator struct{}

func NewWorkspaceValidator() IWorkspaceValidator {
	return &workspaceValidator{}
}

func (wv *workspaceValidator) WorkspaceValidate(workspace model.Workspace) error {
	return validation.ValidateStruct(&workspace,
		validation.Field(
			&workspace.Name,
			validation.Required.Error("name is required"),
			validation.RuneLength(1, 20).Error("limited max 20 char"),
		),
	)
}
