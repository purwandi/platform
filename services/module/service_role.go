package module

import (
	"context"

	"github.com/go-rel/rel"
)

// FindAllRoles is to get all roles
func (s *Service) FindAllRoles(ctx context.Context) ([]Role, error) {
	roles := []Role{}

	err := s.repository.FindAll(ctx, &roles)
	if err != nil {
		return roles, err
	}

	return roles, nil
}

// FindRoleByID ...
func (s *Service) FindRoleByID(ctx context.Context, id int) (Role, error) {
	role := Role{}

	err := s.repository.Find(ctx, &role, rel.Eq("id", id))
	if err != nil {
		return role, err
	}

	return role, nil
}
