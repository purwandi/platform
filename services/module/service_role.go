package module

import "context"

// FindAllRoles is to get all roles
func (s *Service) FindAllRoles(ctx context.Context) ([]Role, error) {
	roles := []Role{}

	err := s.repository.FindAll(ctx, &roles)
	if err != nil {
		return roles, err
	}

	return roles, nil
}
