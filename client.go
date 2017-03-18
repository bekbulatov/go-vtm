package vtm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sync"
)

// VTM is the interface to the VTM API
type VTM interface {
	// -- POOLS ---
	ListPools() ([]string, error)
	Pool(string) (*Pool, error)
	CreatePool(string, *Pool) (*Pool, error)
	DeletePool(string) error

	// --- MISC ---

	// ping the VTM
	Ping() (bool, error)
}

var (
	// ErrInvalidResponse is thrown when VTM responds with invalid or error response
	ErrInvalidResponse = errors.New("invalid response from VTM")
	// ErrTimeoutError is thrown when the operation has timed out
	ErrTimeoutError = errors.New("the operation has timed out")
)

type vtmClient struct {
	sync.RWMutex
	// the configuration for the client
	config Config
	// the http client use for making requests
	httpClient *http.Client
	// a custom logger for debug log messages
	debugLog *log.Logger
}

// NewClient creates a new vtmclient
//		config:			the configuration to use
func NewClient(config Config) VTM {
	// if no http client, set to default
	if config.HTTPClient == nil {
		config.HTTPClient = http.DefaultClient
	}

	debugLogOutput := config.LogOutput
	if debugLogOutput == nil {
		debugLogOutput = ioutil.Discard
	}

	return &vtmClient{
		config:     config,
		httpClient: config.HTTPClient,
		debugLog:   log.New(debugLogOutput, "", 0),
	}
}

// Ping pings the current VTM endpoint (note, this is not a ICMP ping, but a rest api call)
func (r *vtmClient) Ping() (bool, error) {
	if err := r.apiGet(vtmAPIPing, nil, nil); err != nil {
		return false, err
	}
	return true, nil
}

func (r *vtmClient) apiGet(uri string, post, result interface{}) error {
	return r.apiCall("GET", uri, post, result)
}

func (r *vtmClient) apiPut(uri string, post, result interface{}) error {
	return r.apiCall("PUT", uri, post, result)
}

func (r *vtmClient) apiDelete(uri string, post, result interface{}) error {
	return r.apiCall("DELETE", uri, post, result)
}

func (r *vtmClient) apiCall(method, url string, body, result interface{}) error {
	// step: marshall the request to json
	var requestBody []byte
	var err error
	if body != nil {
		if requestBody, err = json.Marshal(body); err != nil {
			return err
		}
	}

	fmt.Printf("Request: %s\n", requestBody)

	// step: create the api request
	request, err := r.buildAPIRequest(method, url, bytes.NewReader(requestBody))
	if err != nil {
		return err
	}
	response, err := r.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// step: read the response body
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Response: %d %s\n", response.StatusCode, respBody)

	if len(requestBody) > 0 {
		r.debugLog.Printf("apiCall(): %v %v %s returned %v %s\n", request.Method, request.URL.String(), requestBody, response.Status, oneLogLine(respBody))
	} else {
		r.debugLog.Printf("apiCall(): %v %v returned %v %s\n", request.Method, request.URL.String(), response.Status, oneLogLine(respBody))
	}

	// step: check for a successfull response
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		if result != nil {
			if err := json.Unmarshal(respBody, result); err != nil {
				r.debugLog.Printf("apiCall(): failed to unmarshall the response from VTM, error: %s\n", err)
				return ErrInvalidResponse
			}
		}
		return nil
	}

	return NewAPIError(response.StatusCode, respBody)
}

// buildAPIRequest creates a default API request
func (r *vtmClient) buildAPIRequest(method, uri string, reader io.Reader) (request *http.Request, err error) {
	// Create the endpoint URL
	url := fmt.Sprintf("%s/%s", r.config.URL, uri)

	// Make the http request to VTM
	request, err = http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}

	// Add any basic auth and the content headers
	if r.config.HTTPBasicAuthUser != "" && r.config.HTTPBasicPassword != "" {
		request.SetBasicAuth(r.config.HTTPBasicAuthUser, r.config.HTTPBasicPassword)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	return request, nil
}

var oneLogLineRegex = regexp.MustCompile(`(?m)^\s*`)

// oneLogLine removes indentation at the beginning of each line and
// escapes new line characters.
func oneLogLine(in []byte) []byte {
	return bytes.Replace(oneLogLineRegex.ReplaceAll(in, nil), []byte("\n"), []byte("\\n "), -1)
}
