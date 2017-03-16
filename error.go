package vtm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	// ErrCodeBadRequest specifies a 400 Bad Request error.
	ErrCodeBadRequest = iota
	// ErrCodeUnauthorized specifies a 401 Unauthorized error.
	ErrCodeUnauthorized
	// ErrCodeForbidden specifies a 403 Forbidden error.
	ErrCodeForbidden
	// ErrCodeNotFound specifies a 404 Not Found error.
	ErrCodeNotFound
	// ErrCodeDuplicateID specifies a PUT 409 Conflict error.
	ErrCodeDuplicateID
	// ErrCodeAppLocked specifies a POST 409 Conflict error.
	ErrCodeAppLocked
	// ErrCodeInvalidBean specifies a 422 UnprocessableEntity error.
	ErrCodeInvalidBean
	// ErrCodeServer specifies a 500+ Server error.
	ErrCodeServer
	// ErrCodeUnknown specifies an unknown error.
	ErrCodeUnknown
)

// InvalidEndpointError indicates a endpoint error in the marathon urls
type InvalidEndpointError struct {
	message string
}

// Error returns the string message
func (e *InvalidEndpointError) Error() string {
	return e.message
}

// newInvalidEndpointError creates a new error
func newInvalidEndpointError(message string, args ...interface{}) error {
	return &InvalidEndpointError{message: fmt.Sprintf(message, args)}
}

// APIError represents a generic API error.
type APIError struct {
	// ErrCode specifies the nature of the error.
	ErrCode int
	message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Marathon API error: %s", e.message)
}

// NewAPIError creates a new APIError instance from the given response code and content.
func NewAPIError(code int, content []byte) error {
	var errDef errorDefinition
	switch {
	case code == http.StatusBadRequest:
		errDef = &badRequestDef{}
	case code == http.StatusUnauthorized:
		errDef = &simpleErrDef{code: ErrCodeUnauthorized}
	case code == http.StatusForbidden:
		errDef = &simpleErrDef{code: ErrCodeForbidden}
	case code == http.StatusNotFound:
		errDef = &simpleErrDef{code: ErrCodeNotFound}
	case code == 422:
		errDef = &unprocessableEntityDef{}
	case code >= http.StatusInternalServerError:
		errDef = &simpleErrDef{code: ErrCodeServer}
	default:
		errDef = &simpleErrDef{code: ErrCodeUnknown}
	}

	return parseContent(errDef, content)
}

type errorDefinition interface {
	message() string
	errCode() int
}

func parseContent(errDef errorDefinition, content []byte) error {
	// If the content cannot be JSON-unmarshalled, we assume that it's not JSON
	// and encode it into the APIError instance as-is.
	errMessage := string(content)
	if err := json.Unmarshal(content, errDef); err == nil {
		errMessage = errDef.message()
	}

	return &APIError{message: errMessage, ErrCode: errDef.errCode()}
}

type simpleErrDef struct {
	Message string `json:"message"`
	code    int
}

func (def *simpleErrDef) message() string {
	return def.Message
}

func (def *simpleErrDef) errCode() int {
	return def.code
}

type detailDescription struct {
	Path   string   `json:"path"`
	Errors []string `json:"errors"`
}

func (d detailDescription) String() string {
	return fmt.Sprintf("path: '%s' errors: %s", d.Path, strings.Join(d.Errors, ", "))
}

type badRequestDef struct {
	Id      string `json:"error_id"`
	Message string `json:"error_text"`
	// TODO: parse `json:"error_info"`
	// Details []detailDescription
}

func (def *badRequestDef) message() string {
	// var details []string
	// for _, detail := range def.Details {
	// 	details = append(details, detail.String())
	// }

	// return fmt.Sprintf("%s (%s)", def.Message, strings.Join(details, "; "))
	return def.Message
}

func (def *badRequestDef) errCode() int {
	return ErrCodeBadRequest
}

type unprocessableEntityDetails []struct {
	// Used in Marathon >= 1.0.0-RC1.
	detailDescription
	// Used in Marathon < 1.0.0-RC1.
	Attribute string `json:"attribute"`
	Error     string `json:"error"`
}

type unprocessableEntityDef struct {
	Message string `json:"message"`
	// Name used in Marathon >= 0.15.0.
	Details unprocessableEntityDetails `json:"details"`
	// Name used in Marathon < 0.15.0.
	Errors unprocessableEntityDetails `json:"errors"`
}

func (def *unprocessableEntityDef) message() string {
	joinDetails := func(details unprocessableEntityDetails) []string {
		var res []string
		for _, detail := range details {
			res = append(res, fmt.Sprintf("attribute '%s': %s", detail.Attribute, detail.Error))
		}
		return res
	}

	var details []string
	switch {
	case len(def.Errors) > 0:
		details = joinDetails(def.Errors)
	case len(def.Details) > 0 && len(def.Details[0].Attribute) > 0:
		details = joinDetails(def.Details)
	default:
		for _, detail := range def.Details {
			details = append(details, detail.detailDescription.String())
		}
	}

	return fmt.Sprintf("%s (%s)", def.Message, strings.Join(details, "; "))
}

func (def *unprocessableEntityDef) errCode() int {
	return ErrCodeInvalidBean
}
