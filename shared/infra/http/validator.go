package http

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
)

type RequestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator() *RequestValidator {
	return &RequestValidator{
		validator: validator.New(),
	}
}

// Validate validates the request body using the validator package and returns an error if validation fails
func (v *RequestValidator) Validate(req interface{}) error {
	slog.Info("running validation function")
	if err := v.validator.Struct(req); err != nil {
		slog.Error("[RequestValidator:Validate]", "error", err, "req", req)
		return err
	}

	return nil
}
