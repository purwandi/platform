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
	t := Type{}

	for _, item := range Types() {
		if string(item.Code) == code {
			t = item
			break
		}
	}

	if (Type{} == t) {
		return t, errors.New("Unrecognized type code")
	}

	return t, nil
}
