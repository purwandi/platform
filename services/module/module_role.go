package module

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// ModuleRole ...
type ModuleRole struct {
	ID              int
	ModuleID        int
	RoleID          int
	Quantity        int
	NiceToHave      *string
	IsFixedQuantity *bool
	IsLead          *bool
	Duration        *int
	TotalHired      int
	TotalConfirmed  int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// CreateModuleRoleInput ...
type CreateModuleRoleInput struct {
	ModuleID        int
	RoleID          int
	Quantity        int
	NiceToHave      *string
	IsFixedQuantity *bool
	IsLead          *bool
	Duration        *int
}

// Validate for validate create moidule role input
func (inpt CreateModuleRoleInput) Validate() error {
	return validation.ValidateStruct(&inpt,
		validation.Field(&inpt.ModuleID, validation.Required),
		validation.Field(&inpt.RoleID, validation.Required),
		validation.Field(&inpt.Quantity, validation.Required),
	)
}
