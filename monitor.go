package vtm

import (
	"encoding/json"
	"fmt"
)

// MonitorBasic contains basic properties
type MonitorBasic struct {
	BackOff  bool   `json:"back_off,omitempty"`
	Delay    int    `json:"delay,omitempty"`
	Failures int    `json:"failures,omitempty"`
	Machine  string `json:"machine,omitempty"`
	Note     string `json:"note,omitempty"`
	Scope    string `json:"scope,omitempty"`
	Timeout  int    `json:"timeout,omitempty"`
	Type     string `json:"type,omitempty"`
	UseSSL   bool   `json:"use_ssl,omitempty"`
	Verbose  bool   `json:"verbose,omitempty"`
}

// MonitorHTTP contains properties for the "http" section
type MonitorHTTP struct {
	Authentication string `json:"authentication,omitempty"`
	BodyRegex      string `json:"body_regex,omitempty"`
	HostHeader     string `json:"host_header,omitempty"`
	Path           string `json:"path,omitempty"`
	StatusRegex    string `json:"status_regex,omitempty"`
}

// MonitorRTSP contains properties for the "rtsp" section
type MonitorRTSP struct {
	BodyRegex   string `json:"body_regex,omitempty"`
	Path        string `json:"path,omitempty"`
	StatusRegex string `json:"status_regex,omitempty"`
}

// MonitorScript contains properties for the "script" section:
type MonitorScript struct {
	Arguments []interface{} `json:"arguments,omitempty"` // TODO (see page 80)
	Program   string        `json:"program,omitempty"`
}

// MonitorSIP contains properties for the "sip" section:
type MonitorSIP struct {
	BodyRegex   string `json:"body_regex,omitempty"`
	StatusRegex string `json:"status_regex,omitempty"`
	Transport   string `json:"transport,omitempty"`
}

// MonitorTCP contains properties for the "tcp" section:
type MonitorTCP struct {
	CloseString    string `json:"close_string,omitempty"`
	MaxResponseLen int    `json:"max_response_len,omitempty"`
	ResponseRegex  string `json:"response_regex,omitempty"`
	WriteString    string `json:"write_string,omitempty"`
}

// MonitorUDP contains properties for the "udp" section:
type MonitorUDP struct {
	AcceptAll bool `json:"accept_all,omitempty"`
}

// Monitor is the definition for a monitor in VTM
type Monitor struct {
	Basic  *MonitorBasic  `json:"basic,omitempty"`
	HTTP   *MonitorHTTP   `json:"http,omitempty"`
	RTSP   *MonitorRTSP   `json:"rtsp,omitempty"`
	Script *MonitorScript `json:"script,omitempty"`
	SIP    *MonitorSIP    `json:"sip,omitempty"`
	TCP    *MonitorTCP    `json:"tcp,omitempty"`
	UDP    *MonitorUDP    `json:"udp,omitempty"`
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

// MonitorExists checks if pool exists in VTM using HEAD request
// 		name: 		the id used to identify the monitor
func (c *vtmClient) MonitorExists(name string) (bool, error) {
	if err := c.apiHead(buildURI(vtmAPIMonitors, name), nil, nil); err != nil {
		return false, err
	}
	return true, nil
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
