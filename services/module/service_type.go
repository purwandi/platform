package module

import (
	"context"
	"errors"
)

// FindAllTypes for get available module types
func (s *Service) FindAllTypes(ctx context.Context) ([]Type, error) {
	return Types(), nil
}

// FindTypeByCode ...
func (s *Service) FindTypeByCode(ctx context.Context, code string) (Type, error) {
	for _, item := range Types() {
		if string(item.Code) == code {
			return item, nil
		}
	}

	return Type{}, errors.New("Unrecognized type code")
}
