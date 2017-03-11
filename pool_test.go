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

// func TestApplicationDependsOn(t *testing.T) {
// 	app := NewDockerApplication()
// 	app.DependsOn("fake-app")
// 	app.DependsOn("fake-app1", "fake-app2")
// 	assert.Equal(t, 3, len(app.Dependencies))
// }
//
// func TestApplicationMemory(t *testing.T) {
// 	app := NewDockerApplication()
// 	app.Memory(50.0)
// 	assert.Equal(t, 50.0, *app.Mem)
// }
//
// func TestApplicationCount(t *testing.T) {
// 	app := NewDockerApplication()
// 	assert.Nil(t, app.Instances)
// 	app.Count(1)
// 	assert.Equal(t, 1, *app.Instances)
// }
//
// func TestApplicationStorage(t *testing.T) {
// 	app := NewDockerApplication()
// 	assert.Nil(t, app.Disk)
// 	app.Storage(0.10)
// 	assert.Equal(t, 0.10, *app.Disk)
// }
//
// func TestApplicationName(t *testing.T) {
// 	app := NewDockerApplication()
// 	assert.Equal(t, "", app.ID)
// 	app.Name(fakeAppName)
// 	assert.Equal(t, fakeAppName, app.ID)
// }
//
// func TestApplicationCommand(t *testing.T) {
// 	app := NewDockerApplication()
// 	assert.Equal(t, "", app.ID)
// 	app.Command("format C:")
// 	assert.Equal(t, "format C:", *app.Cmd)
// }
//
// func TestApplicationCPU(t *testing.T) {
// 	app := NewDockerApplication()
// 	assert.Equal(t, 0.0, app.CPUs)
// 	app.CPU(0.1)
// 	assert.Equal(t, 0.1, app.CPUs)
// }
//
// func TestApplicationArgs(t *testing.T) {
// 	app := NewDockerApplication()
// 	assert.Nil(t, app.Args)
// 	app.AddArgs("-p").AddArgs("option", "-v")
// 	assert.Equal(t, 3, len(*app.Args))
// 	assert.Equal(t, "-p", (*app.Args)[0])
// 	assert.Equal(t, "option", (*app.Args)[1])
// 	assert.Equal(t, "-v", (*app.Args)[2])
//
// 	app.EmptyArgs()
// 	assert.NotNil(t, app.Args)
// 	assert.Equal(t, 0, len(*app.Args))
// }
//
// func ExampleApplication_AddConstraint() {
// 	app := NewDockerApplication()
//
// 	// add two constraints
// 	app.AddConstraint("hostname", "UNIQUE").
// 		AddConstraint("rack_id", "CLUSTER", "rack-1")
// }
//
// func TestApplicationConstraints(t *testing.T) {
// 	app := NewDockerApplication()
// 	assert.Nil(t, app.Constraints)
// 	app.AddConstraint("hostname", "UNIQUE").
// 		AddConstraint("rack_id", "CLUSTER", "rack-1")
//
// 	assert.Equal(t, 2, len(*app.Constraints))
// 	assert.Equal(t, []string{"hostname", "UNIQUE"}, (*app.Constraints)[0])
// 	assert.Equal(t, []string{"rack_id", "CLUSTER", "rack-1"}, (*app.Constraints)[1])
//
// 	app.EmptyConstraints()
// 	assert.NotNil(t, app.Constraints)
// 	assert.Equal(t, 0, len(*app.Constraints))
// }
//
// func TestApplicationLabels(t *testing.T) {
// 	app := NewDockerApplication()
// 	assert.Nil(t, app.Labels)
//
// 	app.AddLabel("hello", "world").AddLabel("foo", "bar")
// 	assert.Equal(t, 2, len(*app.Labels))
// 	assert.Equal(t, "world", (*app.Labels)["hello"])
// 	assert.Equal(t, "bar", (*app.Labels)["foo"])
//
// 	app.EmptyLabels()
// 	assert.NotNil(t, app.Labels)
// 	assert.Equal(t, 0, len(*app.Labels))
// }
//
// func TestApplicationEnvs(t *testing.T) {
// 	app := NewDockerApplication()
// 	assert.Nil(t, app.Env)
//
// 	app.AddEnv("hello", "world").AddEnv("foo", "bar")
// 	assert.Equal(t, 2, len(*app.Env))
// 	assert.Equal(t, "world", (*app.Env)["hello"])
// 	assert.Equal(t, "bar", (*app.Env)["foo"])
//
// 	app.EmptyEnvs()
// 	assert.NotNil(t, app.Env)
// 	assert.Equal(t, 0, len(*app.Env))
// }
//
// func TestApplicationSetExecutor(t *testing.T) {
// 	app := NewDockerApplication()
// 	assert.Nil(t, app.Executor)
//
// 	app.SetExecutor("executor")
// 	assert.Equal(t, "executor", *app.Executor)
//
// 	app.SetExecutor("")
// 	assert.Equal(t, "", *app.Executor)
// }
//
// // func TestCreateApplication(t *testing.T) {
// // 	endpoint := newFakeMarathonEndpoint(t, nil)
// // 	defer endpoint.Close()
// //
// // 	application := NewDockerApplication()
// // 	application.Name(fakeAppName)
// // 	app, err := endpoint.Client.CreateApplication(application)
// // 	assert.NoError(t, err)
// // 	assert.NotNil(t, app)
// // 	assert.Equal(t, application.ID, fakeAppName)
// // 	assert.Equal(t, app.Deployments[0]["id"], "f44fd4fc-4330-4600-a68b-99c7bd33014a")
// // }
// //
// // func TestUpdateApplication(t *testing.T) {
// // 	for _, force := range []bool{false, true} {
// // 		endpoint := newFakeMarathonEndpoint(t, nil)
// // 		defer endpoint.Close()
// //
// // 		application := NewDockerApplication()
// // 		application.Name(fakeAppName)
// // 		id, err := endpoint.Client.UpdateApplication(application, force)
// // 		assert.NoError(t, err)
// // 		assert.Equal(t, id.DeploymentID, "83b215a6-4e26-4e44-9333-5c385eda6438")
// // 		assert.Equal(t, id.Version, "2014-08-26T07:37:50.462Z")
// // 	}
// // }
//
// // func TestApplications(t *testing.T) {
// // 	endpoint := newFakeMarathonEndpoint(t, nil)
// // 	defer endpoint.Close()
// //
// // 	applications, err := endpoint.Client.Applications(nil)
// // 	assert.NoError(t, err)
// // 	assert.NotNil(t, applications)
// // 	assert.Equal(t, len(applications.Apps), 2)
// //
// // 	v := url.Values{}
// // 	v.Set("cmd", "nginx")
// // 	applications, err = endpoint.Client.Applications(v)
// // 	assert.NoError(t, err)
// // 	assert.NotNil(t, applications)
// // 	assert.Equal(t, len(applications.Apps), 1)
// // }
//
func TestListPools(t *testing.T) {
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

	// endpoint := newFakeMarathonEndpoint(t, nil)
	// defer endpoint.Close()

	// applications, err := endpoint.Client.ListPools()
	// assert.NoError(t, err)
	assert.NotNil(t, pools)
	// assert.Equal(t, len(applications), 2)
	// assert.Equal(t, applications[0], fakeAppName)
	// assert.Equal(t, applications[1], fakeAppNameBroken)
}

func TestPool(t *testing.T) {
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

	pool, err := client.Pool("foo-ob_cas_client-bar")
	assert.NoError(t, err)

	// endpoint := newFakeMarathonEndpoint(t, nil)
	// defer endpoint.Close()

	// applications, err := endpoint.Client.ListPools()
	// assert.NoError(t, err)
	assert.NotNil(t, pool)
	// assert.Equal(t, len(applications), 2)
	// assert.Equal(t, applications[0], fakeAppName)
	// assert.Equal(t, applications[1], fakeAppNameBroken)
}

// func TestApplicationVersions(t *testing.T) {
// 	endpoint := newFakeMarathonEndpoint(t, nil)
// 	defer endpoint.Close()
//
// 	versions, err := endpoint.Client.ApplicationVersions(fakeAppName)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, versions)
// 	assert.NotNil(t, versions.Versions)
// 	assert.Equal(t, len(versions.Versions), 1)
// 	assert.Equal(t, versions.Versions[0], "2014-04-04T06:25:31.399Z")
// 	/* check we get an error on app not there */
// 	versions, err = endpoint.Client.ApplicationVersions("/not/there")
// 	assert.Error(t, err)
// }
//
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
// // func TestDeleteApplication(t *testing.T) {
// // 	for _, force := range []bool{false, true} {
// // 		endpoint := newFakeMarathonEndpoint(t, nil)
// // 		defer endpoint.Close()
// // 		id, err := endpoint.Client.DeleteApplication(fakeAppName, force)
// // 		assert.NoError(t, err)
// // 		assert.NotNil(t, id)
// // 		assert.Equal(t, "83b215a6-4e26-4e44-9333-5c385eda6438", id.DeploymentID)
// // 		assert.Equal(t, "2014-08-26T07:37:50.462Z", id.Version)
// // 		id, err = endpoint.Client.DeleteApplication("no_such_app", force)
// // 		assert.Error(t, err)
// // 	}
// // }
//
// // func TestApplicationOK(t *testing.T) {
// // 	endpoint := newFakeMarathonEndpoint(t, nil)
// // 	defer endpoint.Close()
// //
// // 	ok, err := endpoint.Client.ApplicationOK(fakeAppName)
// // 	assert.NoError(t, err)
// // 	assert.True(t, ok)
// // 	ok, err = endpoint.Client.ApplicationOK(fakeAppNameBroken)
// // 	assert.NoError(t, err)
// // 	assert.False(t, ok)
// // 	ok, err = endpoint.Client.ApplicationOK(fakeAppNameUnhealthy)
// // 	assert.NoError(t, err)
// // 	assert.False(t, ok)
// // }
//
// func verifyApplication(application *Application, t *testing.T) {
// 	assert.NotNil(t, application)
// 	assert.Equal(t, application.ID, fakeAppName)
// 	assert.NotNil(t, application.HealthChecks)
// 	assert.Equal(t, len(*application.HealthChecks), 1)
// }
//
// func TestApplication(t *testing.T) {
// 	endpoint := newFakeMarathonEndpoint(t, nil)
// 	defer endpoint.Close()
//
// 	application, err := endpoint.Client.Application(fakeAppName)
// 	assert.NoError(t, err)
// 	verifyApplication(application, t)
//
// 	_, err = endpoint.Client.Application("no_such_app")
// 	assert.Error(t, err)
// 	apiErr, ok := err.(*APIError)
// 	assert.True(t, ok)
// 	assert.Equal(t, ErrCodeNotFound, apiErr.ErrCode)
//
// 	config := NewDefaultConfig()
// 	config.URL = "http://non-existing-marathon-host.local:5555"
// 	// Reduce timeout to speed up test execution time.
// 	config.HTTPClient = &http.Client{
// 		Timeout: 100 * time.Millisecond,
// 	}
// 	endpoint = newFakeMarathonEndpoint(t, &configContainer{
// 		client: &config,
// 	})
// 	defer endpoint.Close()
//
// 	_, err = endpoint.Client.Application(fakeAppName)
// 	assert.Error(t, err)
// 	_, ok = err.(*APIError)
// 	assert.False(t, ok)
// }
//
// func TestApplicationConfiguration(t *testing.T) {
// 	endpoint := newFakeMarathonEndpoint(t, nil)
// 	defer endpoint.Close()
//
// 	application, err := endpoint.Client.ApplicationByVersion(fakeAppName, "2014-09-12T23:28:21.737Z")
// 	assert.NoError(t, err)
// 	verifyApplication(application, t)
//
// 	_, err = endpoint.Client.ApplicationByVersion(fakeAppName, "no_such_version")
// 	assert.Error(t, err)
// 	apiErr, ok := err.(*APIError)
// 	assert.True(t, ok)
// 	assert.Equal(t, ErrCodeNotFound, apiErr.ErrCode)
//
// 	_, err = endpoint.Client.ApplicationByVersion("no_such_app", "latest")
// 	assert.Error(t, err)
// 	apiErr, ok = err.(*APIError)
// 	assert.True(t, ok)
// 	assert.Equal(t, ErrCodeNotFound, apiErr.ErrCode)
// }
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
//
// func TestAppExistAndRunning(t *testing.T) {
// 	endpoint := newFakeMarathonEndpoint(t, nil)
// 	defer endpoint.Close()
// 	client := endpoint.Client.(*marathonClient)
// 	assert.True(t, client.appExistAndRunning(fakeAppName))
// 	assert.False(t, client.appExistAndRunning("no_such_app"))
// }
//
// func TestSetIPPerTask(t *testing.T) {
// 	app := Application{}
// 	app.Ports = append(app.Ports, 10)
// 	app.AddPortDefinition(PortDefinition{})
// 	assert.Nil(t, app.IPAddressPerTask)
// 	assert.Equal(t, 1, len(app.Ports))
// 	assert.Equal(t, 1, len(*app.PortDefinitions))
//
// 	app.SetIPAddressPerTask(IPAddressPerTask{})
// 	assert.NotNil(t, app.IPAddressPerTask)
// 	assert.Equal(t, 0, len(app.Ports))
// 	assert.Equal(t, 0, len(*app.PortDefinitions))
// }
//
// func TestIPAddressPerTask(t *testing.T) {
// 	ipPerTask := IPAddressPerTask{}
// 	assert.Nil(t, ipPerTask.Groups)
// 	assert.Nil(t, ipPerTask.Labels)
// 	assert.Nil(t, ipPerTask.Discovery)
//
// 	ipPerTask.
// 		AddGroup("label").
// 		AddLabel("key", "value").
// 		SetDiscovery(Discovery{})
//
// 	assert.Equal(t, 1, len(*ipPerTask.Groups))
// 	assert.Equal(t, "label", (*ipPerTask.Groups)[0])
// 	assert.Equal(t, "value", (*ipPerTask.Labels)["key"])
// 	assert.NotEmpty(t, ipPerTask.Discovery)
//
// 	ipPerTask.EmptyGroups()
// 	assert.Equal(t, 0, len(*ipPerTask.Groups))
//
// 	ipPerTask.EmptyLabels()
// 	assert.Equal(t, 0, len(*ipPerTask.Labels))
//
// }
//
// func TestIPAddressPerTaskDiscovery(t *testing.T) {
// 	disc := Discovery{}
// 	assert.Nil(t, disc.Ports)
//
// 	disc.AddPort(Port{})
// 	assert.NotNil(t, disc.Ports)
// 	assert.Equal(t, 1, len(*disc.Ports))
//
// 	disc.EmptyPorts()
// 	assert.NotNil(t, disc.Ports)
// 	assert.Equal(t, 0, len(*disc.Ports))
//
// }
//
// func TestUpgradeStrategy(t *testing.T) {
// 	app := Application{}
// 	assert.Nil(t, app.UpgradeStrategy)
// 	app.SetUpgradeStrategy(UpgradeStrategy{}.SetMinimumHealthCapacity(1.0).SetMaximumOverCapacity(0.0))
// 	us := app.UpgradeStrategy
// 	assert.Equal(t, 1.0, *us.MinimumHealthCapacity)
// 	assert.Equal(t, 0.0, *us.MaximumOverCapacity)
//
// 	app.EmptyUpgradeStrategy()
// 	us = app.UpgradeStrategy
// 	assert.NotNil(t, us)
// 	assert.Nil(t, us.MinimumHealthCapacity)
// 	assert.Nil(t, us.MaximumOverCapacity)
// }