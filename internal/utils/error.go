package utils

import (
	"net/http"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

func Err(val int, err error) *models.Error {

	code := int64(val)
	status := http.StatusText(val)
	message := err.Error()

	return &models.Error{
		Code:    &code,
		Status:  &status,
		Message: &message,
	}
}
