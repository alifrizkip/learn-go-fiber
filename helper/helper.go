package helper

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

// Response struct
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// Meta struct
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Success bool   `json:"success"`
}

// ErrorsData ...
func ErrorsData(data interface{}) interface{} {
	return map[string]interface{}{
		"errors": data,
	}
}

// APIResponse ...
func APIResponse(message string, code int, isSuccess bool, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Success: isSuccess,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

// ValidateRequest ...
func ValidateRequest(req interface{}) []string {
	var errors []string

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(
				errors,
				fmt.Sprintf("`%v` with value `%v` doesn't satisfy the `%v` constraint", err.Field(), err.Value(), err.Tag()),
			)
		}
	}
	return errors
}

// SendAPIResponse ...
func SendAPIResponse(c *fiber.Ctx) func(string, int, bool, interface{}) error {
	return func(message string, code int, isSuccess bool, data interface{}) error {
		response := APIResponse(message, code, isSuccess, data)

		return c.Status(response.Meta.Code).JSON(response)
	}
}
