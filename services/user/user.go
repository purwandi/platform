package user

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// User ...
type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    []byte    `json:"-"`
	AccessToken *string   `json:"-"`
	Fullname    *string   `json:"fullname"`
	Callname    *string   `json:"callname"`
	PhoneNumber *string   `json:"phone_number"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

// RegisterInput ...
type RegisterInput struct {
	Username        string
	Email           string
	Password        string
	PasswordConfirm string
}

// Validate ...
func (inpt RegisterInput) Validate() error {
	return validation.ValidateStruct(&inpt,
		validation.Field(&inpt.Username, validation.Required),
		validation.Field(&inpt.Email, validation.Required, is.Email),
		validation.Field(&inpt.Password, validation.Required),
		validation.Field(&inpt.PasswordConfirm, validation.Required),
	)
}
