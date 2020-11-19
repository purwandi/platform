package user

import "time"

// ProjectRole ...
type ProjectRole struct {
	ID        int
	Name      string
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Table ...
func (p *ProjectRole) Table() string {
	return "project_roles"
}

type UserProjectRole struct {
	UserID        int
	ProjectRoleID int
}
