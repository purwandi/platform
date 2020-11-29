package module

import "time"

// Role for role
type Role struct {
	ID          int
	Name        string
	Requirement string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
