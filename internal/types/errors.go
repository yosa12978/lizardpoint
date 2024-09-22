package types

import "errors"

var (
	ErrInternalFailure  error = errors.New("internal failure")
	ErrValidationFailed error = errors.New("validation failed")
	ErrNotFound         error = errors.New("not found")
)

func NewErrInternalFailure(err error) error {
	return errors.Join(ErrInternalFailure, err)
}

func NewErrValidationFailed(err error) error {
	return errors.Join(ErrValidationFailed, err)
}

func NewErrNotFound(err error) error {
	return errors.Join(ErrNotFound, err)
}
