package vtm

import (
	"io"
	"io/ioutil"
	"net/http"
)

// Config holds the settings and options for the client
type Config struct {
	// URL is the url for stingray
	URL string
	// HTTPBasicAuthUser is the http basic auth
	HTTPBasicAuthUser string
	// HTTPBasicPassword is the http basic password
	HTTPBasicPassword string
	// LogOutput the output for debug log messages
	LogOutput io.Writer
	// HTTPClient is the http client
	HTTPClient *http.Client
}

// NewDefaultConfig create a default client config
func NewDefaultConfig() Config {
	return Config{
		URL:       "http://127.0.0.1:8080",
		LogOutput: ioutil.Discard,
	}
}
