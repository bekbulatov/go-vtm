package vtm

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPoolNodes(t *testing.T) {
	pool := NewPool()
	assert.Nil(t, pool.Nodes())

	pool.AddNode(NodeItem{Node: "foo.bar.com:8080"}).AddNode(NodeItem{Node: "foo.bar.com:8081"})
	assert.Equal(t, 2, len(*pool.Nodes()))
	assert.Equal(t, "foo.bar.com:8080", (*pool.Nodes())[0].Node)
	assert.Equal(t, "foo.bar.com:8081", (*pool.Nodes())[1].Node)

	pool.EmptyNodes()
	assert.NotNil(t, pool.Nodes())
	assert.Equal(t, 0, len(*pool.Nodes()))
}

func TestListPools(t *testing.T) {
	endpoint := newFakeVTMEndpoint(t, nil)
	defer endpoint.Close()

	pools, err := endpoint.Client.ListPools()
	assert.NoError(t, err)
	assert.NotNil(t, pools)
	assert.Equal(t, len(pools), 2)
	assert.Equal(t, pools[0], "foo")
	assert.Equal(t, pools[1], "bar")
}

func TestPool(t *testing.T) {
	endpoint := newFakeVTMEndpoint(t, nil)
	defer endpoint.Close()

	pool, err := endpoint.Client.Pool(fakePoolName)
	assert.NoError(t, err)
	assert.NotNil(t, pool)
	assert.NotNil(t, pool.Nodes())
	assert.Equal(t, len(*pool.Nodes()), 1)
	assert.Equal(t, (*pool.Nodes())[0].Node, "foo.bar.com:31199")

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
	endpoint = newFakeVTMEndpoint(t, &configContainer{
		client: &config,
	})
	defer endpoint.Close()

	_, err = endpoint.Client.Pool(fakePoolName)
	assert.Error(t, err)
	_, ok = err.(*APIError)
	assert.False(t, ok)
}

func TestCreatePool(t *testing.T) {
	endpoint := newFakeVTMEndpoint(t, nil)
	defer endpoint.Close()

	pool := NewPool()
	pool.AddNode(NodeItem{Node: "foo.bar.com:123"})
	p, err := endpoint.Client.CreatePool(fakePoolName, pool)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.NotNil(t, p.Nodes())
	assert.Equal(t, len(*p.Nodes()), 1)
	assert.Equal(t, (*p.Nodes())[0].Node, "foo.bar.com:123")
}

func TestDeletePool(t *testing.T) {
	endpoint := newFakeVTMEndpoint(t, nil)
	defer endpoint.Close()

	err := endpoint.Client.DeletePool(fakePoolName)
	assert.NoError(t, err)

	err = endpoint.Client.DeletePool("no_such_app")
	assert.Error(t, err)
}
