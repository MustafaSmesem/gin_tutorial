package validator

import (
	"github.com/go-playground/validator/v10"
	"strings"
)
import "golang.org/x/exp/slices"

func ValidateGoodName(filed validator.FieldLevel) bool {
	goodNames := []string{
		"mustafa", "mohammed", "ahmed",
	}
	if slices.Contains(goodNames, strings.ToLower(filed.Field().String())) {
		return true
	}
	return false
}
