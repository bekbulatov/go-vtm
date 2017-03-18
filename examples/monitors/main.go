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

	monitors, err := client.ListMonitors()
	assert(err)

	log.Printf("Found %d monitors", len(monitors))

	monitorName := "my-monitor"

	log.Printf("Creating the monitor: %s", monitorName)
	monitor := vtm.NewMonitor()
	_, err = client.CreateMonitor(monitorName, monitor)
	assert(err)
	log.Printf("Successfully created the monitor: %s", monitorName)

	monitor, err = client.Monitor(monitorName)
	assert(err)
	log.Printf("Monitor: %v", monitor)

	log.Printf("Deleting the monitor: %s", monitorName)
	err = client.DeleteMonitor(monitorName)
	assert(err)
	log.Printf("Successfully deleted the monitor: %s", monitorName)

}
