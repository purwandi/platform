package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-rel/rel"
	"github.com/purwandi/platform/pkg/flaw"
)

// AuthService ...
type AuthService struct{}

// NewAuthService ...
func NewAuthService() *AuthService {
	return &AuthService{}
}

// User is to get current user ..
func (s *AuthService) User(ctx context.Context) (User, error) {
	u := ctx.Value("auth")
	if u == nil {
		return User{}, errors.New("Invalid user")
	}

	return u.(User), nil
}

// Service ...
type Service struct {
	repository rel.Repository
}

// NewService is to register the service
func NewService(repo rel.Repository) *Service {
	return &Service{repository: repo}
}

// CreateUser ...
func (s *Service) CreateUser(ctx context.Context, inpt RegisterInput) (User, error) {
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
func (s *Service) Authenticate(ctx context.Context, inpt LoginInput) (User, error) {

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
func (s *Service) FindByUsername(ctx context.Context, username string) (User, error) {
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
func (s *Service) FindUserByAccessToken(ctx context.Context, token string) (User, error) {
	// Var
	u := User{}

	// Process
	if err := s.repository.Find(ctx, &u, rel.Eq("access_token", token)); err != nil {
		return u, flaw.InternalError(err.Error())
	}

	return u, nil
}

// FindByUserID ...
func (s *Service) FindByUserID(ctx context.Context, id int) (User, error) {
	// Var
	u := User{}

	// Process
	if err := s.repository.Find(ctx, &u, rel.Eq("id", id)); err != nil {
		return u, flaw.InternalError(err.Error())
	}

	return u, nil
}

// UserExists ...
func (s *Service) UserExists(ctx context.Context, id int) bool {
	u, err := s.FindByUserID(ctx, id)
	if (err != nil) || (u.ID == 0) {
		return false
	}

	return true
}
