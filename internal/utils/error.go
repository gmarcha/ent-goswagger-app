package utils

import (
	"net/http"

	"github.com/gmarcha/ent-goswagger-app/internal/goswagger/models"
)

// Err takes an error code and error message as parameters
// and returns an initialise Error type.
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
