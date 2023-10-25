package data

import "github.com/go-playground/validator/v10"

// Validation contains
type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()

	return &Validation{validate}
}
