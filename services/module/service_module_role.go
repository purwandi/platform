package module

import (
	"context"
	"time"
)

// CreateModuleRole ...
func (s *Service) CreateModuleRole(ctx context.Context, inpt CreateModuleRoleInput) (ModuleRole, error) {
	// Validate
	if err := inpt.Validate(); err != nil {
		return ModuleRole{}, err
	}

	m := ModuleRole{
		RoleID:          inpt.RoleID,
		ModuleID:        inpt.ModuleID,
		Quantity:        inpt.Quantity,
		NiceToHave:      inpt.NiceToHave,
		IsFixedQuantity: inpt.IsFixedQuantity,
		IsLead:          inpt.IsLead,
		Duration:        inpt.Duration,
		CreatedAt:       time.Now(),
	}

	// Persist
	if err := s.repository.Insert(ctx, &m); err != nil {
		return ModuleRole{}, err
	}

	return m, nil
}
