package usecase

import "errors"

// Common errors
var ErrInvalidPassword = errors.New("Errors.Common.InvalidPassword")
var ErrValidationFailed = errors.New("Errors.Common.ValidationFailed")
var ErrInternalDatabaseError = errors.New("Errors.Common.InternalDatabaseError")
