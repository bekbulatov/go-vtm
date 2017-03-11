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

// Marathon is the interface to the marathon API
type Marathon interface {
	// -- POOLS ---
	ListPools() ([]string, error)
	Pool(string) (*Pool, error)
}

var (
	// ErrInvalidResponse is thrown when marathon responds with invalid or error response
	ErrInvalidResponse = errors.New("invalid response from Marathon")
	// ErrTimeoutError is thrown when the operation has timed out
	ErrTimeoutError = errors.New("the operation has timed out")
)

type marathonClient struct {
	sync.RWMutex
	// the configuration for the client
	config Config
	// the ip address of the client
	ipAddress string
	// the http server
	eventsHTTP *http.Server
	// the http client use for making requests
	httpClient *http.Client
	// a custom logger for debug log messages
	debugLog *log.Logger
}

// NewClient creates a new marathon client
//		config:			the configuration to use
func NewClient(config Config) (Marathon, error) {
	// step: if no http client, set to default
	if config.HTTPClient == nil {
		config.HTTPClient = http.DefaultClient
	}

	debugLogOutput := config.LogOutput
	if debugLogOutput == nil {
		debugLogOutput = ioutil.Discard
	}

	return &marathonClient{
		config:     config,
		httpClient: config.HTTPClient,
		debugLog:   log.New(debugLogOutput, "", 0),
	}, nil
}

// Ping pings the current marathon endpoint (note, this is not a ICMP ping, but a rest api call)
// func (r *marathonClient) Ping() (bool, error) {
// 	if err := r.apiGet(marathonAPIPing, nil, nil); err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

func (r *marathonClient) apiGet(uri string, post, result interface{}) error {
	return r.apiCall("GET", uri, post, result)
}

func (r *marathonClient) apiPut(uri string, post, result interface{}) error {
	return r.apiCall("PUT", uri, post, result)
}

func (r *marathonClient) apiPost(uri string, post, result interface{}) error {
	return r.apiCall("POST", uri, post, result)
}

func (r *marathonClient) apiDelete(uri string, post, result interface{}) error {
	return r.apiCall("DELETE", uri, post, result)
}

func (r *marathonClient) apiCall(method, url string, body, result interface{}) error {
	// step: marshall the request to json
	var requestBody []byte
	var err error
	if body != nil {
		if requestBody, err = json.Marshal(body); err != nil {
			return err
		}
	}

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

	if len(requestBody) > 0 {
		r.debugLog.Printf("apiCall(): %v %v %s returned %v %s\n", request.Method, request.URL.String(), requestBody, response.Status, oneLogLine(respBody))
	} else {
		r.debugLog.Printf("apiCall(): %v %v returned %v %s\n", request.Method, request.URL.String(), response.Status, oneLogLine(respBody))
	}

	// step: check for a successfull response
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		if result != nil {
			if err := json.Unmarshal(respBody, result); err != nil {
				r.debugLog.Printf("apiCall(): failed to unmarshall the response from marathon, error: %s\n", err)
				return ErrInvalidResponse
			}
		}
		return nil
	}

	return NewAPIError(response.StatusCode, respBody)
}

// buildAPIRequest creates a default API request
func (r *marathonClient) buildAPIRequest(method, uri string, reader io.Reader) (request *http.Request, err error) {
	// Create the endpoint URL
	url := fmt.Sprintf("%s/%s", r.config.URL, uri)

	// Make the http request to Marathon
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
