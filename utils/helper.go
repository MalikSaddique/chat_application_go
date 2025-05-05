package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func DecryptErrors(err error) map[string]string {
	var ve validator.ValidationErrors
	out := make(map[string]string)

	if errors.As(err, &ve) {
		for _, fe := range ve {
			field := strings.ToLower(fe.Field())
			tag := fe.Tag()
			param := fe.Param()

			msgMap := map[string]string{
				"required": fmt.Sprintf("%s is required", field),
				"email":    "invalid email format",
				"min":      fmt.Sprintf("%s must be at least %s characters long", field, param),
				"max":      fmt.Sprintf("%s must be at most %s characters long", field, param),
				"alphanum": fmt.Sprintf("%s must contain only letters and numbers", field),
			}

			if msg, ok := msgMap[tag]; ok {
				out[field] = msg
			} else {
				out[field] = fmt.Sprintf("%s is not valid", field)
			}
		}
	} else {
		out["error"] = "invalid request format"
	}

	return out
}
