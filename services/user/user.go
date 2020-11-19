package user

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"github.com/purwandi/platform/pkg/cast"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID           int     `json:"id"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	Password     []byte  `json:"-"`
	AccessToken  *string `json:"-"`
	Fullname     *string `json:"fullname"`
	Callname     *string `json:"callname"`
	PhoneNumber  *string `json:"phone_number"`
	ProjectRoles []int
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

// CreateUser ...
func CreateUser(inpt RegisterInput) (*User, error) {
	// Generate password
	hash, err := bcrypt.GenerateFromPassword([]byte(inpt.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user domain
	return &User{
		Username:  inpt.Username,
		Email:     inpt.Email,
		Password:  hash,
		CreatedAt: time.Now(),
	}, nil
}

// GenerateAccessToken ...
func (u *User) GenerateAccessToken() error {
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	u.AccessToken = cast.String(uid.String())
	return nil
}

// GetAccessToken ...
func (u *User) GetAccessToken() string {
	return cast.StringValue(u.AccessToken)
}

// IsPasswordIsValid check if password is valid
func (u *User) IsPasswordIsValid(passwd string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(u.Password, []byte(passwd)); err != nil {
		return false, err
	}

	return true, nil
}

// RegisterInput ...
type RegisterInput struct {
	Username string
	Email    string
	Password string
}

// Validate ...
func (inpt RegisterInput) Validate() error {
	return validation.ValidateStruct(&inpt,
		validation.Field(&inpt.Username, validation.Required),
		validation.Field(&inpt.Email, validation.Required, is.Email),
		validation.Field(&inpt.Password, validation.Required),
	)
}

// LoginInput ...
type LoginInput struct {
	Username string
	Password string
}

// Validate ...
func (inpt LoginInput) Validate() error {
	return validation.ValidateStruct(&inpt,
		validation.Field(&inpt.Username, validation.Required),
		validation.Field(&inpt.Password, validation.Required),
	)
}
