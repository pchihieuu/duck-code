// Package validator exposes the single go-playground/validator instance
// that also backs Gin's request binding, so custom validation rules
// (e.g. "strongpassword") are registered once and apply everywhere —
// both in c.ShouldBindJSON and in manual validation calls.
package validator

import (
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var (
	once sync.Once
	v    *validator.Validate
)

// Get returns the shared validator instance, registering custom rules
// on first use.
func Get() *validator.Validate {
	once.Do(func() {
		if engine, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v = engine
		} else {
			v = validator.New()
		}
		registerCustomRules(v)
	})
	return v
}

func registerCustomRules(v *validator.Validate) {
	// Example: password strength beyond min-length, used on auth.RegisterRequest
	// via `binding:"required,min=8,strongpassword"` if/when you want it enforced.
	_ = v.RegisterValidation("strongpassword", func(fl validator.FieldLevel) bool {
		s := fl.Field().String()
		var hasUpper, hasDigit bool
		for _, r := range s {
			switch {
			case r >= 'A' && r <= 'Z':
				hasUpper = true
			case r >= '0' && r <= '9':
				hasDigit = true
			}
		}
		return hasUpper && hasDigit
	})
}
