package vtm

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListMonitors(t *testing.T) {
	endpoint := newFakeVTMEndpoint(t, nil)
	defer endpoint.Close()

	monitors, err := endpoint.Client.ListMonitors()
	assert.NoError(t, err)
	assert.NotNil(t, monitors)
	assert.Equal(t, len(monitors), 2)
	assert.Equal(t, monitors[0], "foo")
	assert.Equal(t, monitors[1], "bar")
}

func TestMonitor(t *testing.T) {
	endpoint := newFakeVTMEndpoint(t, nil)
	defer endpoint.Close()

	monitor, err := endpoint.Client.Monitor(fakeMonitorName)
	assert.NoError(t, err)
	assert.NotNil(t, monitor)

	_, err = endpoint.Client.Monitor("no_such_monitor")
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
	endpoint = newFakeVTMEndpoint(t, &configContainer{
		client: &config,
	})
	defer endpoint.Close()

	_, err = endpoint.Client.Monitor(fakeMonitorName)
	assert.Error(t, err)
	_, ok = err.(*APIError)
	assert.False(t, ok)
}

func TestCreateMonitor(t *testing.T) {
	endpoint := newFakeVTMEndpoint(t, nil)
	defer endpoint.Close()

	monitor := NewMonitor()
	p, err := endpoint.Client.CreateMonitor(fakeMonitorName, monitor)
	assert.NoError(t, err)
	assert.NotNil(t, p)
}

func TestDeleteMonitor(t *testing.T) {
	endpoint := newFakeVTMEndpoint(t, nil)
	defer endpoint.Close()

	err := endpoint.Client.DeleteMonitor(fakeMonitorName)
	assert.NoError(t, err)

	err = endpoint.Client.DeleteMonitor("no_such_app")
	assert.Error(t, err)
}
