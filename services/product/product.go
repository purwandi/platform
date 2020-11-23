package product

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Product ...
type Product struct {
	ID          int
	UserID      int
	Code        *string
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CreateProductInput ...
type CreateProductInput struct {
	UserID      int
	Code        *string
	Name        string
	Description *string
}

// Validate ...
func (inpt CreateProductInput) Validate() error {
	return validation.ValidateStruct(&inpt,
		validation.Field(&inpt.UserID, validation.Required),
		validation.Field(&inpt.Name, validation.Required),
	)
}

// UpdateProductInput ...
type UpdateProductInput struct {
	ID          int
	Code        *string
	Name        string
	Description *string
}

// Validate ...
func (inpt UpdateProductInput) Validate() error {
	return validation.ValidateStruct(&inpt,
		validation.Field(&inpt.ID, validation.Required),
		validation.Field(&inpt.Name, validation.Required),
	)
}

// CreateProduct ...
func CreateProduct(inpt CreateProductInput) (*Product, error) {
	if err := inpt.Validate(); err != nil {
		return nil, err
	}

	return &Product{
		UserID:      inpt.UserID,
		Code:        inpt.Code,
		Name:        inpt.Name,
		Description: inpt.Description,
		CreatedAt:   time.Now(),
	}, nil
}
