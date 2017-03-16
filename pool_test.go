package vtm

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

	client := NewClient(config)

	// ListPools

	pools, err := client.ListPools()
	assert.NoError(t, err)
	assert.True(t, len(pools) > 0)

	// CreatePool

	pool := new(Pool)
	pool.Basic = &Basic{
		NodesTable: []NodeItem{
			NodeItem{Node: "foo.bar.com:123"},
		},
	}

	p, err := client.CreatePool(fakePoolName, pool)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.NotNil(t, p.Basic)
	assert.NotNil(t, p.Basic.NodesTable)

	// DeletePool

	fmt.Println("!!!")
	err = client.DeletePool(fakePoolName)
	assert.NoError(t, err)

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

func TestCreatePool(t *testing.T) {
	endpoint := newFakeMarathonEndpoint(t, nil)
	defer endpoint.Close()

	pool := new(Pool)
	// application.Name(fakeAppName)
	p, err := endpoint.Client.CreatePool(fakePoolName, pool)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	// assert.Equal(t, application.ID, fakeAppName)
	// assert.Equal(t, app.Deployments[0]["id"], "f44fd4fc-4330-4600-a68b-99c7bd33014a")
}

func TestDeletePool(t *testing.T) {
	endpoint := newFakeMarathonEndpoint(t, nil)
	defer endpoint.Close()

	err := endpoint.Client.DeletePool(fakePoolName)
	assert.NoError(t, err)

	err = endpoint.Client.DeletePool("no_such_app")
	assert.Error(t, err)
}
