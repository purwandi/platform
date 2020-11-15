package user

import (
	"context"
	"net/http"
	"time"

	"github.com/go-rel/rel"
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
func (service *Service) CreateUser(ctx context.Context, inpt RegisterInput) (User, error) {
	// Validate input
	if err := inpt.Validate(); err != nil {
		return User{}, flaw.Error(http.StatusBadRequest, err.Error())
	}

	// Generate password
	hash, err := bcrypt.GenerateFromPassword([]byte(inpt.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, flaw.Error(http.StatusInternalServerError, err.Error())
	}

	// Create user domain
	user := User{
		Username:  inpt.Username,
		Email:     inpt.Email,
		Password:  hash,
		CreatedAt: time.Now(),
	}

	// Persist data
	if err := service.repository.Insert(ctx, &user); err != nil {
		return User{}, flaw.Error(http.StatusInternalServerError, err.Error())
	}

	return user, nil
}
