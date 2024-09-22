package validation

import "context"

type Validatable[T any] interface {
	Validate(ctx context.Context) (T, map[string]string, bool)
}
