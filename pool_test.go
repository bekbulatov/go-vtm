package vtm

import (
	// "fmt"
	"crypto/tls"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// func TestCreateApplication(t *testing.T) {
// 	endpoint := newFakeMarathonEndpoint(t, nil)
// 	defer endpoint.Close()
//
// 	application := NewDockerApplication()
// 	application.Name(fakeAppName)
// 	app, err := endpoint.Client.CreateApplication(application)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, app)
// 	assert.Equal(t, application.ID, fakeAppName)
// 	assert.Equal(t, app.Deployments[0]["id"], "f44fd4fc-4330-4600-a68b-99c7bd33014a")
// }

func TestHttpClient(t *testing.T) {
	config := NewDefaultConfig()
	config.URL = "https://192.168.99.100:9070"
	config.HTTPClient = &http.Client{
		Timeout: (time.Duration(1) * time.Second),
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   500 * time.Millisecond,
				KeepAlive: 10 * time.Second,
			}).Dial,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	config.HTTPBasicAuthUser = "admin"
	config.HTTPBasicPassword = "bob"

	client, err := NewClient(config)
	assert.NoError(t, err)

	pools, err := client.ListPools()
	assert.NoError(t, err)

	assert.True(t, len(pools) > 0)
}

func TestListPools(t *testing.T) {
	endpoint := newFakeMarathonEndpoint(t, nil)
	defer endpoint.Close()

	pools, err := endpoint.Client.ListPools()
	assert.NoError(t, err)
	assert.NotNil(t, pools)
	assert.Equal(t, len(pools), 2)
	assert.Equal(t, pools[0], "foo")
	assert.Equal(t, pools[1], "bar")
}

func TestPool(t *testing.T) {
	endpoint := newFakeMarathonEndpoint(t, nil)
	defer endpoint.Close()

	pool, err := endpoint.Client.Pool(fakePoolName)
	assert.NoError(t, err)
	assert.NotNil(t, pool)
	assert.Equal(t, len(pool.Basic.NodesTable), 1)
	assert.Equal(t, pool.Basic.NodesTable[0].Node, "foo.bar.com:31199")

	_, err = endpoint.Client.Pool("no_such_pool")
	assert.Error(t, err)
	apiErr, ok := err.(*APIError)
	assert.True(t, ok)
	assert.Equal(t, ErrCodeNotFound, apiErr.ErrCode)

	config := NewDefaultConfig()
	config.URL = "http://non-existing-host.local:9070"
	// Reduce timeout to speed up test execution time.
	config.HTTPClient = &http.Client{
		Timeout: 100 * time.Millisecond,
	}
	endpoint = newFakeMarathonEndpoint(t, &configContainer{
		client: &config,
	})
	defer endpoint.Close()

	_, err = endpoint.Client.Pool(fakePoolName)
	assert.Error(t, err)
	_, ok = err.(*APIError)
	assert.False(t, ok)
}

// func TestRestartApplication(t *testing.T) {
// 	endpoint := newFakeMarathonEndpoint(t, nil)
// 	defer endpoint.Close()
//
// 	id, err := endpoint.Client.RestartApplication(fakeAppName, false)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, id)
// 	assert.Equal(t, "83b215a6-4e26-4e44-9333-5c385eda6438", id.DeploymentID)
// 	assert.Equal(t, "2014-08-26T07:37:50.462Z", id.Version)
// 	id, err = endpoint.Client.RestartApplication("/not/there", false)
// 	assert.Error(t, err)
// 	assert.Nil(t, id)
// }
//
// func TestApplicationUris(t *testing.T) {
// 	app := NewDockerApplication()
// 	assert.Nil(t, app.Uris)
// 	app.AddUris("file://uri1.tar.gz").AddUris("file://uri2.tar.gz", "file://uri3.tar.gz")
// 	assert.Equal(t, 3, len(*app.Uris))
// 	assert.Equal(t, "file://uri1.tar.gz", (*app.Uris)[0])
// 	assert.Equal(t, "file://uri2.tar.gz", (*app.Uris)[1])
// 	assert.Equal(t, "file://uri3.tar.gz", (*app.Uris)[2])
//
// 	app.EmptyUris()
// 	assert.NotNil(t, app.Uris)
// 	assert.Equal(t, 0, len(*app.Uris))
// }
//
// func TestApplicationFetchURIs(t *testing.T) {
// 	app := NewDockerApplication()
// 	assert.Nil(t, app.Fetch)
// 	app.AddFetchURIs(Fetch{URI: "file://uri1.tar.gz"}).
// 		AddFetchURIs(Fetch{URI: "file://uri2.tar.gz"}, Fetch{URI: "file://uri3.tar.gz"})
// 	assert.Equal(t, 3, len(*app.Fetch))
// 	assert.Equal(t, Fetch{URI: "file://uri1.tar.gz"}, (*app.Fetch)[0])
// 	assert.Equal(t, Fetch{URI: "file://uri2.tar.gz"}, (*app.Fetch)[1])
// 	assert.Equal(t, Fetch{URI: "file://uri3.tar.gz"}, (*app.Fetch)[2])
//
// 	app.EmptyUris()
// 	assert.NotNil(t, app.Uris)
// 	assert.Equal(t, 0, len(*app.Uris))
// }
//
// func TestDeleteApplication(t *testing.T) {
// 	for _, force := range []bool{false, true} {
// 		endpoint := newFakeMarathonEndpoint(t, nil)
// 		defer endpoint.Close()
// 		id, err := endpoint.Client.DeleteApplication(fakeAppName, force)
// 		assert.NoError(t, err)
// 		assert.NotNil(t, id)
// 		assert.Equal(t, "83b215a6-4e26-4e44-9333-5c385eda6438", id.DeploymentID)
// 		assert.Equal(t, "2014-08-26T07:37:50.462Z", id.Version)
// 		id, err = endpoint.Client.DeleteApplication("no_such_app", force)
// 		assert.Error(t, err)
// 	}
// }
//
// func verifyApplication(application *Application, t *testing.T) {
// 	assert.NotNil(t, application)
// 	assert.Equal(t, application.ID, fakeAppName)
// 	assert.NotNil(t, application.HealthChecks)
// 	assert.Equal(t, len(*application.HealthChecks), 1)
// }
//
//
// func TestWaitOnApplication(t *testing.T) {
// 	waitTime := 100 * time.Millisecond
//
// 	tests := []struct {
// 		desc          string
// 		timeout       time.Duration
// 		appName       string
// 		testScope     string
// 		shouldSucceed bool
// 	}{
// 		{
// 			desc:          "initially existing app",
// 			timeout:       0,
// 			appName:       fakeAppName,
// 			shouldSucceed: true,
// 		},
//
// 		{
// 			desc:          "delayed existing app | timeout > ticker",
// 			timeout:       200 * time.Millisecond,
// 			appName:       fakeAppName,
// 			testScope:     "wait-on-app",
// 			shouldSucceed: true,
// 		},
//
// 		{
// 			desc:          "delayed existing app | timeout < ticker",
// 			timeout:       50 * time.Millisecond,
// 			appName:       fakeAppName,
// 			testScope:     "wait-on-app",
// 			shouldSucceed: false,
// 		},
// 		{
// 			desc:          "missing app | timeout > ticker",
// 			timeout:       200 * time.Millisecond,
// 			appName:       "no_such_app",
// 			shouldSucceed: false,
// 		},
// 		{
// 			desc:          "missing app | timeout < ticker",
// 			timeout:       50 * time.Millisecond,
// 			appName:       "no_such_app",
// 			shouldSucceed: false,
// 		},
// 	}
//
// 	for _, test := range tests {
// 		defaultConfig := NewDefaultConfig()
// 		defaultConfig.PollingWaitTime = waitTime
// 		configs := &configContainer{
// 			client: &defaultConfig,
// 			server: &serverConfig{
// 				scope: test.testScope,
// 			},
// 		}
//
// 		endpoint := newFakeMarathonEndpoint(t, configs)
// 		defer endpoint.Close()
//
// 		errCh := make(chan error)
// 		go func() {
// 			errCh <- endpoint.Client.WaitOnApplication(test.appName, test.timeout)
// 		}()
//
// 		select {
// 		case <-time.After(400 * time.Millisecond):
// 			assert.Fail(t, fmt.Sprintf("%s: WaitOnApplication did not complete in time", test.desc))
// 		case err := <-errCh:
// 			if test.shouldSucceed {
// 				assert.NoError(t, err, test.desc)
// 			} else {
// 				assert.IsType(t, err, ErrTimeoutError, test.desc)
// 			}
// 		}
// 	}
// }
