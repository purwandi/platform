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
	ID          int
	Code        typeModule
	Name        string
	Description *string
}

// Types ...
func Types() []Type {
	return []Type{
		{ID: 1, Code: ModuleID, Name: "In-Depth Interview"},
		{ID: 2, Code: ModuleTS, Name: "Testing"},
		{ID: 3, Code: ModuleTR, Name: "Training"},
		{ID: 4, Code: ModuleCC, Name: "Co Creation"},
		{ID: 5, Code: ModuleSS, Name: "Scrum Sprint"},
		{ID: 6, Code: ModuleDS, Name: "Design Sprint"},
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
