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
	// ErrCodeServer specifies a 500+ Server error.
	ErrCodeServer
	// ErrCodeUnknown specifies an unknown error.
	ErrCodeUnknown
)

// APIError represents a generic API error.
type APIError struct {
	ErrCode  int // ErrCode specifies the nature of the error.
	httpCode int
	message  string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("VTM API error %d: %s", e.httpCode, e.message)
}

// NewAPIError creates a new APIError instance from the given response code and content.
func NewAPIError(code int, content []byte) error {
	errDef := &errorDef{httpCode: code}
	switch {
	case code == http.StatusBadRequest:
		errDef.errCode = ErrCodeBadRequest
	case code == http.StatusUnauthorized:
		errDef.errCode = ErrCodeUnauthorized
	case code == http.StatusForbidden:
		errDef.errCode = ErrCodeForbidden
	case code == http.StatusNotFound:
		errDef.errCode = ErrCodeNotFound
	case code >= http.StatusInternalServerError:
		errDef.errCode = ErrCodeServer
	default:
		errDef.errCode = ErrCodeUnknown
	}
	return parseContent(errDef, content)
}

func parseContent(errDef *errorDef, content []byte) error {
	// If the content cannot be JSON-unmarshalled, we assume that it's not JSON
	// and encode it into the APIError instance as-is.
	errMessage := string(content)
	if err := json.Unmarshal(content, errDef); err == nil {
		errMessage = errDef.message()
	}

	return &APIError{ErrCode: errDef.errCode, httpCode: errDef.httpCode, message: errMessage}
}

type errorDef struct {
	ID       string                 `json:"error_id"`
	Message  string                 `json:"error_text"`
	Details  map[string]interface{} `json:"error_info"`
	httpCode int
	errCode  int
}

func (def *errorDef) message() string {
	message := def.Message
	if len(def.Details) > 0 {
		var details []string
		parseMap("", def.Details, &details)
		message = fmt.Sprintf("%s (%s)", message, strings.Join(details, ", "))
	}
	return message
}

func parseMap(prefix string, m map[string]interface{}, result *[]string) {
	for key, val := range m {
		switch val.(type) {
		case map[string]interface{}:
			// escape key
			if strings.Contains(key, ".") {
				key = fmt.Sprintf(`"%s"`, key)
			}
			// add key to prefix
			if prefix == "" {
				prefix = key
			} else {
				prefix = fmt.Sprintf("%s.%s", prefix, key)
			}

			if value, ok := val.(map[string]interface{}); ok {
				if errorText, ok := value["error_text"]; ok {
					*result = append(*result, fmt.Sprintf("%s: %s", prefix, errorText))
				} else {
					parseMap(prefix, value, result)
				}
			}
		}
	}
}
