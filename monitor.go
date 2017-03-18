package vtm

import (
	"encoding/json"
	"fmt"
)

type MonitorBasic struct {
	BackOff  bool   `json:"back_off"`
	Delay    int    `json:"delay"`
	Failures int    `json:"failures"`
	Machine  string `json:"machine"`
	Note     string `json:"note"`
	Scope    string `json:"scope"`
	Timeout  int    `json:"timeout"`
	Type     string `json:"type"`
	UseSsl   bool   `json:"use_ssl"`
	Verbose  bool   `json:"verbose"`
}
type MonitorHTTP struct {
	Authentication string `json:"authentication"`
	BodyRegex      string `json:"body_regex"`
	HostHeader     string `json:"host_header"`
	Path           string `json:"path"`
	StatusRegex    string `json:"status_regex"`
}
type MonitorRtsp struct {
	BodyRegex   string `json:"body_regex"`
	Path        string `json:"path"`
	StatusRegex string `json:"status_regex"`
}
type MonitorScript struct {
	Arguments []interface{} `json:"arguments"`
	Program   string        `json:"program"`
}
type MonitorSip struct {
	BodyRegex   string `json:"body_regex"`
	StatusRegex string `json:"status_regex"`
	Transport   string `json:"transport"`
}
type MonitorTCP struct {
	CloseString    string `json:"close_string"`
	MaxResponseLen int    `json:"max_response_len"`
	ResponseRegex  string `json:"response_regex"`
	WriteString    string `json:"write_string"`
}
type MonitorUDP struct {
	AcceptAll bool `json:"accept_all"`
}

// Monitor is the definition for a monitor in VTM
type Monitor struct {
	Basic  *MonitorBasic  `json:"basic"`
	HTTP   *MonitorHTTP   `json:"http"`
	Rtsp   *MonitorRtsp   `json:"rtsp"`
	Script *MonitorScript `json:"script"`
	Sip    *MonitorSip    `json:"sip"`
	TCP    *MonitorTCP    `json:"tcp"`
	UDP    *MonitorUDP    `json:"udp"`
}

// NewMonitor creates a default monitor
func NewMonitor() *Monitor {
	monitor := new(Monitor)
	return monitor
}

// String returns the json representation of this monitor
func (p *Monitor) String() string {
	s, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error": "error decoding type into json: %s"}`, err)
	}
	return string(s)
}

type monitorWrapper struct {
	Monitor *Monitor `json:"properties"`
}

// ListMonitors retrieves an array of the monitor names currently registered in VTM
func (c *vtmClient) ListMonitors() ([]string, error) {
	var wrapper struct {
		Monitors []struct {
			Name string `json:"name"`
			Href string `json:"href"`
		} `json:"children"`
	}
	if err := c.apiGet(vtmAPIMonitors, nil, &wrapper); err != nil {
		return nil, err
	}

	var list []string
	for _, monitor := range wrapper.Monitors {
		list = append(list, monitor.Name)
	}

	return list, nil
}

// Monitor retrieves the monitor configuration from VTM
// 		name: 		the id used to identify the monitor
func (c *vtmClient) Monitor(name string) (*Monitor, error) {
	result := new(monitorWrapper)
	if err := c.apiGet(buildURI(vtmAPIMonitors, name), nil, &result); err != nil {
		return nil, err
	}
	return result.Monitor, nil
}

// CreateMonitor creates a new monitor in VTM
// 		monitor:		the structure holding the monitor configuration
func (c *vtmClient) CreateMonitor(name string, monitor *Monitor) (*Monitor, error) {
	wrapper := new(monitorWrapper)
	wrapper.Monitor = monitor
	result := new(monitorWrapper)
	if err := c.apiPut(buildURI(vtmAPIMonitors, name), wrapper, &result); err != nil {
		return nil, err
	}
	return result.Monitor, nil
}

// DeleteMonitor deletes an application from VTM
// 		name: 		the id used to identify the monitor
func (c *vtmClient) DeleteMonitor(name string) error {
	if err := c.apiDelete(buildURI(vtmAPIMonitors, name), nil, nil); err != nil {
		return err
	}
	return nil
}
