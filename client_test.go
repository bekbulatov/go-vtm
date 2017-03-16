package vtm

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	config := Config{
		URL: "http://marathon",
	}
	cl := NewClient(config)

	conf := cl.(*vtmClient).config

	assert.Equal(t, conf.HTTPClient, http.DefaultClient)
}

func TestPing(t *testing.T) {
	endpoint := newFakeMarathonEndpoint(t, nil)
	defer endpoint.Close()

	pong, err := endpoint.Client.Ping()
	assert.NoError(t, err)
	assert.True(t, pong)
}

func TestAPIRequest(t *testing.T) {
	cases := []struct {
		Username       string
		Password       string
		ServerUsername string
		ServerPassword string
		Ok             bool
	}{
		{
			Username:       "should_pass",
			Password:       "",
			ServerUsername: "",
			ServerPassword: "",
			Ok:             true,
		},
		{
			Username:       "bad_username",
			Password:       "",
			ServerUsername: "test",
			ServerPassword: "password",
			Ok:             false,
		},
		{
			Username:       "test",
			Password:       "bad_password",
			ServerUsername: "test",
			ServerPassword: "password",
			Ok:             false,
		},
		{
			Username:       "",
			Password:       "",
			ServerUsername: "test",
			ServerPassword: "password",
			Ok:             false,
		},
		{
			Username:       "test",
			Password:       "password",
			ServerUsername: "test",
			ServerPassword: "password",
			Ok:             true,
		},
	}
	for i, x := range cases {
		var endpoint *endpoint

		config := NewDefaultConfig()
		config.HTTPBasicAuthUser = x.Username
		config.HTTPBasicPassword = x.Password

		endpoint = newFakeMarathonEndpoint(t, &configContainer{
			client: &config,
			server: &serverConfig{
				username: x.ServerUsername,
				password: x.ServerPassword,
			},
		})

		_, err := endpoint.Client.ListPools()

		if x.Ok && err != nil {
			t.Errorf("case %d, did not expect an error: %s", i, err)
		}
		if !x.Ok && err == nil {
			t.Errorf("case %d, expected to received an error", i)
		}

		endpoint.Close()
	}
}

func TestOneLogLine(t *testing.T) {
	in := `
	a
	b    c
	d\n
	  efgh
	i\r\n
	j\t
	{"json":  "works",
		"f o o": "ba    r"
	}
	`
	assert.Equal(t, `a\n b    c\n d\n\n efgh\n i\r\n\n j\t\n {"json":  "works",\n "f o o": "ba    r"\n }\n `, string(oneLogLine([]byte(in))))
}
