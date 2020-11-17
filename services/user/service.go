package user

import (
	"context"
	"net/http"
	"time"

	"github.com/go-rel/rel"
	"github.com/google/uuid"
	"github.com/purwandi/platform/pkg/cast"
	"github.com/purwandi/platform/pkg/flaw"
	"golang.org/x/crypto/bcrypt"
)

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

	// Generate password
	hash, err := bcrypt.GenerateFromPassword([]byte(inpt.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, flaw.InternalError(err.Error())
	}

	// Create user domain
	user := User{
		Username:  inpt.Username,
		Email:     inpt.Email,
		Password:  hash,
		CreatedAt: time.Now(),
	}

	// Persist data
	if err := s.repository.Insert(ctx, &user); err != nil {
		return User{}, flaw.Error(http.StatusInternalServerError, err.Error())
	}

	return user, nil
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

	// Compare password
	if err := bcrypt.CompareHashAndPassword(u.Password, []byte(inpt.Password)); err != nil {
		return User{}, flaw.BadRequest("Password is not valid")
	}

	// Update token
	u.AccessToken = cast.String(uuid.New().String())
	changeset := rel.Set("access_token", u.AccessToken)

	// Persist data
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
