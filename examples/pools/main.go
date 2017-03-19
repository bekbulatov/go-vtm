package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net"
	"net/http"
	"time"

	vtm "cx/go-vtm"
)

var (
	vtmURL       string
	authUser     string
	authPassword string
)

func init() {
	flag.StringVar(&vtmURL, "url", "https://127.0.0.1:9070", "the url for the VTM endpoint")
	flag.StringVar(&authUser, "user", "", "username for HTTP basic auth")
	flag.StringVar(&authPassword, "password", "", "password for HTTP basic auth")
}

func assert(err error) {
	if err != nil {
		log.Fatalf("Failed, error: %s", err)
	}
}

func main() {
	flag.Parse()

	config := vtm.NewDefaultConfig()
	config.URL = vtmURL
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
	config.HTTPBasicAuthUser = authUser
	config.HTTPBasicPassword = authPassword

	client := vtm.NewClient(config)

	pools, err := client.ListPools()
	assert(err)

	log.Printf("Found %d pools", len(pools))

	poolName := "my-pool"

	log.Printf("Creating the pool: %s", poolName)
	pool := vtm.NewPool()
	pool.AddNode(vtm.NodeItem{Node: "foo.bar.com:123", Weight: 1})
	_, err = client.CreatePool(poolName, pool)
	assert(err)
	log.Printf("Successfully created the pool: %s", poolName)

	pool, err = client.Pool("my-pool222")
	assert(err)
	log.Printf("Pool: %v", pool)

	log.Printf("Deleting the pool: %s", poolName)
	err = client.DeletePool(poolName)
	assert(err)
	log.Printf("Successfully deleted the pool: %s", poolName)

}
