package vtm

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

const (
	fakePoolName    = "test-pool"
	fakeMonitorName = "test-monitor"
)

var (
	fakeResponses map[string][]indexedResponse
	once          sync.Once
)

type indexedResponse struct {
	Index   int    `yaml:"index,omitempty"`
	Content string `yaml:"content,omitempty"`
}

type responseIndices struct {
	sync.Mutex
	m map[string]int
}

func newResponseIndices() *responseIndices {
	return &responseIndices{m: map[string]int{}}
}

// restMethod represents an expected HTTP method and an associated fake response
type restMethod struct {
	// the uri of the method
	URI string `yaml:"uri,omitempty"`
	// the http method type (GET|PUT etc)
	Method string `yaml:"method,omitempty"`
	// the content i.e. response
	Content string `yaml:"content,omitempty"`
	// ContentSequence is a sequence of responses that are returned in order.
	ContentSequence []indexedResponse `yaml:"contentSequence,omitempty"`
	// the test scope
	Scope string `yaml:"scope,omitempty"`
}

// serverConfig holds the VTM server configuration
type serverConfig struct {
	// Username for basic auth
	username string
	// Password for basic auth
	password string
	// scope is an arbitrary test scope to distinguish fake responses from
	// otherwise equal HTTP methods and query strings.
	scope string
}

// configContainer holds both server and client VTM configuration
type configContainer struct {
	client *Config
	server *serverConfig
}

type fakeServer struct {
	io.Closer

	httpSrv         *httptest.Server
	fakeRespIndices *responseIndices
}

type endpoint struct {
	io.Closer

	Server fakeServer
	Client VTM
	URL    string
}

func getTestURL(urlString string) string {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		panic(fmt.Sprintf("failed to parse URL '%s': %s", urlString, err))
	}
	return fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)
}

func newFakeVTMEndpoint(t *testing.T, configs *configContainer) *endpoint {
	// step: read in the fake responses if required
	initFakeVTMResponses(t)

	// step: fill in the default if required
	defaultConfig := NewDefaultConfig()
	if configs == nil {
		configs = &configContainer{}
	}
	if configs.client == nil {
		configs.client = &defaultConfig
	}
	if configs.server == nil {
		configs.server = &serverConfig{}
	}

	fakeRespIndices := newResponseIndices()

	// step: create the HTTP router
	mux := http.NewServeMux()
	mux.HandleFunc("/", authMiddleware(configs.server, func(writer http.ResponseWriter, reader *http.Request) {
		respKey := fakeResponseMapKey(reader.Method, reader.RequestURI, configs.server.scope)
		fakeRespIndices.Lock()
		fakeRespIndex := fakeRespIndices.m[respKey]
		fakeRespIndices.m[respKey]++
		responses, found := fakeResponses[respKey]
		fakeRespIndices.Unlock()
		if found {
			for _, response := range responses {
				// Index < 0 indicates a static response.
				if response.Index < 0 || response.Index == fakeRespIndex {
					writer.Header().Add("Content-Type", "application/json")
					writer.Write([]byte(response.Content))
					return
				}
			}
		}

		http.Error(writer, `{error_id": "resource.not_found"}`, 404)
	}))

	// step: create HTTP test server
	httpSrv := httptest.NewServer(mux)

	if configs.client.URL == defaultConfig.URL {
		configs.client.URL = getTestURL(httpSrv.URL)
	}

	// step: create the client for the service
	client := NewClient(*configs.client)

	return &endpoint{
		Server: fakeServer{
			httpSrv:         httpSrv,
			fakeRespIndices: fakeRespIndices,
		},
		Client: client,
		URL:    configs.client.URL,
	}
}

// authMiddleware handles basic auth
func authMiddleware(server *serverConfig, next http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	unauthorized := `{"message": "invalid username or password"}`

	return func(w http.ResponseWriter, r *http.Request) {
		// step: is authentication required?

		if server.username != "" && server.password != "" {
			u, p, found := r.BasicAuth()
			// step: if no auth found, error it
			if !found {
				http.Error(w, unauthorized, 401)
				return
			}
			// step: if username and password don't match, error it
			if server.username != u || server.password != p {
				http.Error(w, unauthorized, 401)
				return
			}
		}

		next(w, r)
	}
}

// initFakeVTMResponses reads in the vtm fake responses from the yaml file
func initFakeVTMResponses(t *testing.T) {
	once.Do(func() {
		fakeResponses = make(map[string][]indexedResponse, 0)
		var methods []*restMethod

		// step: read in the test method specification
		methodSpec, err := ioutil.ReadFile("./tests/rest-api/methods.yml")
		if err != nil {
			t.Fatalf("failed to read in the fake yaml responses")
		}

		err = yaml.Unmarshal([]byte(methodSpec), &methods)
		if err != nil {
			t.Fatalf("failed to unmarshal the response")
		}
		for _, method := range methods {
			key := fakeResponseMapKey(method.Method, method.URI, method.Scope)
			switch {
			case method.Content != "" && len(method.ContentSequence) > 0:
				panic("content and contentSequence must not be provided simultaneously")
			case len(method.ContentSequence) > 0:
				fakeResponses[key] = method.ContentSequence
			default:
				// This combines the cases where static content was defined or not. The
				// latter models an empty response (via an empty content) that should
				// not result into a 404.
				fakeResponses[key] = []indexedResponse{
					indexedResponse{
						// Index -1 indicates a static response.
						Index:   -1,
						Content: method.Content,
					},
				}
			}
		}
	})
}

func fakeResponseMapKey(method, uri, scope string) string {
	return fmt.Sprintf("%s:%s:%s", method, uri, scope)
}

func (s *fakeServer) Close() error {
	s.httpSrv.Close()
	return nil
}

func (e *endpoint) Close() error {
	return e.Server.Close()
}
