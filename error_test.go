package vtm

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		httpCode   int
		nameSuffix string
		errCode    int
		errText    string
		content    string
	}{
		// 400
		{
			httpCode: http.StatusBadRequest,
			errCode:  ErrCodeBadRequest,
			errText:  `VTM API error 400: The resource provided is invalid (basic.nodes_table."foo.bar.com:123".weight: Value '-1' must be within 1-100)`,
			content:  content400(),
		},
		// 401
		{
			httpCode: http.StatusUnauthorized,
			errCode:  ErrCodeUnauthorized,
			errText:  "VTM API error 401: User name or password was invalid",
			content:  `{"error_id":"auth.invalid","error_text":"User name or password was invalid"}`,
		},
		// 403
		{
			httpCode: http.StatusForbidden,
			errCode:  ErrCodeForbidden,
			errText:  "VTM API error 403: Not Authorized to perform this action!",
			content:  `{"error_text": "Not Authorized to perform this action!"}`,
		},
		// 404
		{
			httpCode: http.StatusNotFound,
			errCode:  ErrCodeNotFound,
			errText:  "VTM API error 404: Resource '/api/tm/3.5/config/active/pools/my-pool' does not exist",
			content:  `{"error_id":"resource.not_found","error_text":"Resource '/api/tm/3.5/config/active/pools/my-pool' does not exist"}`,
		},
		// 499 unknown error
		{
			httpCode:   499,
			nameSuffix: "unknown error",
			errCode:    ErrCodeUnknown,
			errText:    "VTM API error 499: unknown error",
			content:    `{"error_text": "unknown error"}`,
		},
		// 500
		{
			httpCode: http.StatusInternalServerError,
			errCode:  ErrCodeServer,
			errText:  "VTM API error 500: internal server error",
			content:  `{"error_text": "internal server error"}`,
		},
		// // 503 (no JSON)
		{
			httpCode:   http.StatusServiceUnavailable,
			nameSuffix: "no JSON",
			errCode:    ErrCodeServer,
			errText:    "VTM API error 503: No server is available to handle this request.",
			content:    `No server is available to handle this request.`,
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d", test.httpCode)
		if len(test.nameSuffix) > 0 {
			name = fmt.Sprintf("%s (%s)", name, test.nameSuffix)
		}
		apiErr := NewAPIError(test.httpCode, []byte(test.content))
		gotErrCode := apiErr.(*APIError).ErrCode
		assert.Equal(t, test.errCode, gotErrCode, fmt.Sprintf("HTTP code %s (error code): got %d, want %d", name, gotErrCode, test.errCode))
		assert.Equal(t, apiErr.Error(), test.errText, fmt.Sprintf("HTTP code %s (error text)", name))
	}
}

func content400() string {
	return `{
  "error_id":"resource.validation_error",
  "error_text":"The resource provided is invalid",
  "error_info":{
    "basic":{
      "nodes_table":{
        "foo.bar.com:123":{
          "weight":{
            "error_id":"num.out_of_range",
            "error_text":"Value '-1' must be within 1-100"
          }
        }
      }
    }
  }
}`
}
