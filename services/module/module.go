package module

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type typeModule string

const (
	ModuleID typeModule = "ID"
	ModuleTS typeModule = "TS"
	ModuleTR typeModule = "TR"
	ModuleCC typeModule = "CC"
	ModuleSS typeModule = "SS"
	ModuleDS typeModule = "DS"
)

// Type ...
type Type struct {
	Code        typeModule
	Name        string
	Description *string
}

// Types ...
func Types() []Type {
	return []Type{
		{Code: ModuleID, Name: "In-Depth Interview"},
		{Code: ModuleTS, Name: "Testing"},
		{Code: ModuleTR, Name: "Training"},
		{Code: ModuleCC, Name: "Co Creation"},
		{Code: ModuleSS, Name: "Scrum Sprint"},
		{Code: ModuleDS, Name: "Design Sprint"},
	}
}

type status string

const (
	ARCHIVED  status = "ARCHIVED"
	DRAFT     status = "DRAFT"
	PUBLISHED status = "PUBLISHED"
	POSTPONED status = "POSTPONED"
	ONGOING   status = "ONGOING"
	CLOSED    status = "CLOSED"
)

// Module ...
type Module struct {
	ID          int
	OrderID     *int
	ProjectID   int
	TypeID      typeModule
	Status      status
	Location    *string
	SuccessRate *float32
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CreateModuleInput ...
type CreateModuleInput struct {
	OrderID     *int
	ProjectID   int
	TypeID      typeModule
	Location    *string
	Description *string
}

// Validate ...
func (inpt CreateModuleInput) Validate() error {
	return validation.ValidateStruct(&inpt,
		validation.Field(&inpt.OrderID, validation.Required),
		validation.Field(&inpt.TypeID, validation.Required),
	)
}

// History ...
type History struct {
	ID        int
	ModuleID  int
	Status    status
	CreatedAt time.Time
}
