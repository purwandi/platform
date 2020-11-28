package module

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// TypeID for module type id
type TypeID string

const (
	// ModuleID for in-depth interview
	ModuleID TypeID = "ID"
	// ModuleTS for testing
	ModuleTS TypeID = "TS"
	// ModuleTR for training
	ModuleTR TypeID = "TR"
	// ModuleCC for co-creation
	ModuleCC TypeID = "CC"
	// ModuleSS for scrum sprint
	ModuleSS TypeID = "SS"
	// ModuleDS for design sprint
	ModuleDS TypeID = "DS"
)

// Type ...
type Type struct {
	Code        TypeID
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

// Status flag module statuss
type Status string

const (
	// ARCHIVED when project is archived
	ARCHIVED Status = "ARCHIVED"
	// DRAFT when project is draft
	DRAFT Status = "DRAFT"
	// PUBLISHED when project is draft
	PUBLISHED Status = "PUBLISHED"
	// POSTPONED when project is postponed
	POSTPONED Status = "POSTPONED"
	// ONGOING when project is ongoing
	ONGOING Status = "ONGOING"
	// CLOSED when project is closed
	CLOSED Status = "CLOSED"
)

// Module ...
type Module struct {
	ID          int
	OrderID     *int
	ProjectID   int
	TypeID      TypeID
	Status      Status
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
	TypeID      TypeID
	Location    *string
	Description *string
}

// Validate ...
func (inpt CreateModuleInput) Validate() error {
	return validation.ValidateStruct(&inpt,
		validation.Field(&inpt.ProjectID, validation.Required),
		validation.Field(&inpt.TypeID, validation.Required),
	)
}

// History ...
type History struct {
	ID        int
	ModuleID  int
	Status    Status
	CreatedAt time.Time
}
