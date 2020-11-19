package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-rel/rel"
	"github.com/purwandi/platform/pkg/flaw"
)

// UserService ...
type UserService struct {
	repository rel.Repository
}

// NewUserService is to register the service
func NewUserService(repo rel.Repository) *UserService {
	return &UserService{repository: repo}
}

// CreateUser ...
func (s *UserService) CreateUser(ctx context.Context, inpt RegisterInput) (User, error) {
	// Validate input
	if err := inpt.Validate(); err != nil {
		return User{}, flaw.Error(http.StatusBadRequest, err.Error())
	}

	u, err := CreateUser(inpt)
	if err != nil {
		return User{}, flaw.InternalError(err.Error())
	}

	// Persist data
	if err := s.repository.Insert(ctx, u); err != nil {
		return User{}, flaw.Error(http.StatusInternalServerError, err.Error())
	}

	return *u, nil
}

// Authenticate ...
func (s *UserService) Authenticate(ctx context.Context, inpt LoginInput) (User, error) {

	// Validate input
	if err := inpt.Validate(); err != nil {
		return User{}, flaw.BadRequest(err.Error())
	}

	// Process
	u, err := s.FindByUsername(ctx, inpt.Username)
	if err != nil {
		return User{}, err
	}

	// check if password is valid
	if _, err := (&u).IsPasswordIsValid(inpt.Password); err != nil {
		fmt.Println(err.Error(), " Helloo")
		return User{}, flaw.BadRequest("Password is not valid")
	}

	// Generate token
	if err := (&u).GenerateAccessToken(); err != nil {
		return User{}, flaw.InternalError("Unable to generate access token")
	}

	// Persist data
	changeset := rel.Set("access_token", u.GetAccessToken())
	if err := s.repository.Update(ctx, &u, changeset); err != nil {
		return User{}, flaw.InternalError("Unable to save data")
	}

	// Response
	return u, nil
}

// FindByUsername ...
func (s *UserService) FindByUsername(ctx context.Context, username string) (User, error) {
	// Var
	u := User{}

	// Process
	if err := s.repository.Find(ctx, &u, rel.Eq("username", username)); err != nil {
		return u, flaw.InternalError(err.Error())
	}

	// Response
	return u, nil
}

// FindUserByAccessToken ...
func (s *UserService) FindUserByAccessToken(ctx context.Context, token string) (User, error) {
	// Var
	u := User{}

	// Process
	if err := s.repository.Find(ctx, &u, rel.Eq("access_token", token)); err != nil {
		return u, flaw.InternalError(err.Error())
	}

	return u, nil
}

// AttachProjectRoles ...
func (s *UserService) AttachProjectRoles(ctx context.Context, u *User, roles []int) (User, error) {
	u.ProjectRoles = roles

	changeset := rel.Set("project_roles", roles)
	if err := s.repository.Update(ctx, u, changeset); err != nil {
		return *u, err
	}

	return User{}, nil
}
