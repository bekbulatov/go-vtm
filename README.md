# go-vtm [![Build Status](https://travis-ci.org/bekbulatov/go-vtm.svg?branch=master)](https://travis-ci.org/bekbulatov/go-vtm) [![Coverage Status](https://coveralls.io/repos/github/bekbulatov/go-vtm/badge.svg?branch=master)](https://coveralls.io/github/bekbulatov/go-vtm?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/bekbulatov/go-vtm)](https://goreportcard.com/report/github.com/bekbulatov/go-vtm)

Brocade VTM Rest Client

Currently supported objects are:
- Pools
- Monitors
    
## Code Examples

There is an examples directory in the source which contains some code of how to use it.
    
### Creating a client
    
    config := vtm.NewDefaultConfig()
    config.URL = "http://127.0.0.1:9070"

    client := vtm.NewClient(config)

    pools, err := client.ListPools()
    if err == nil {
      log.Printf("Found %d pools", len(pools))
    }

### Creating a new pool

    poolName := "my-pool"

    pool := vtm.NewPool()
    pool.AddNode(vtm.NodeItem{Node: "foo.bar.com:123", Weight: 1})
    _, err = client.CreatePool(poolName, pool)
    if err == nil {
      log.Printf("Successfully created the pool: %s", poolName)
    }
