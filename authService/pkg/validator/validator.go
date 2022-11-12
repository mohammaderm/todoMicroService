package validator

import (
	"context"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func AuthRequest(ctx context.Context, stru interface{}) error {
	return validate.StructCtx(ctx, stru)
}
