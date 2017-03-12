package vtm

import (
	// "encoding/json"
	"fmt"
)

type Basic struct {
	BandwidthClass                string   `json:"bandwidth_class"`
	FailurePool                   string   `json:"failure_pool"`
	MaxConnectionAttempts         int      `json:"max_connection_attempts"`
	MaxIdleConnectionsPernode     int      `json:"max_idle_connections_pernode"`
	MaxTimedOutConnectionAttempts int      `json:"max_timed_out_connection_attempts"`
	Monitors                      []string `json:"monitors"`
	NodeCloseWithRst              bool     `json:"node_close_with_rst"`
	NodeConnectionAttempts        int      `json:"node_connection_attempts"`
	NodeDeleteBehavior            string   `json:"node_delete_behavior"`
	NodeDrainToDeleteTimeout      int      `json:"node_drain_to_delete_timeout"`
	NodesTable                    []struct {
		Node     string `json:"node"`
		Priority int    `json:"priority"`
		State    string `json:"state"`
		Weight   int    `json:"weight"`
	} `json:"nodes_table"`
	Note              string `json:"note"`
	PassiveMonitoring bool   `json:"passive_monitoring"`
	PersistenceClass  string `json:"persistence_class"`
	Transparent       bool   `json:"transparent"`
}

type AutoScaling struct {
	AddnodeDelaytime int           `json:"addnode_delaytime"`
	CloudCredentials string        `json:"cloud_credentials"`
	Cluster          string        `json:"cluster"`
	DataCenter       string        `json:"data_center"`
	DataStore        string        `json:"data_store"`
	Enabled          bool          `json:"enabled"`
	External         bool          `json:"external"`
	Hysteresis       int           `json:"hysteresis"`
	Imageid          string        `json:"imageid"`
	IpsToUse         string        `json:"ips_to_use"`
	LastNodeIdleTime int           `json:"last_node_idle_time"`
	MaxNodes         int           `json:"max_nodes"`
	MinNodes         int           `json:"min_nodes"`
	Name             string        `json:"name"`
	Port             int           `json:"port"`
	Refractory       int           `json:"refractory"`
	ResponseTime     int           `json:"response_time"`
	ScaleDownLevel   int           `json:"scale_down_level"`
	ScaleUpLevel     int           `json:"scale_up_level"`
	Securitygroupids []interface{} `json:"securitygroupids"`
	SizeID           string        `json:"size_id"`
	Subnetids        []interface{} `json:"subnetids"`
}

type Connection struct {
	MaxConnectTime        int `json:"max_connect_time"`
	MaxConnectionsPerNode int `json:"max_connections_per_node"`
	MaxQueueSize          int `json:"max_queue_size"`
	MaxReplyTime          int `json:"max_reply_time"`
	QueueTimeout          int `json:"queue_timeout"`
}

type DNSAutoscale struct {
	Enabled   bool          `json:"enabled"`
	Hostnames []interface{} `json:"hostnames"`
	Port      int           `json:"port"`
}
type FTP struct {
	SupportRfc2428 bool `json:"support_rfc_2428"`
}
type HTTP struct {
	Keepalive              bool `json:"keepalive"`
	KeepaliveNonIdempotent bool `json:"keepalive_non_idempotent"`
}
type KerberosProtocolTransition struct {
	Principal string `json:"principal"`
	Target    string `json:"target"`
}
type LoadBalancing struct {
	Algorithm       string `json:"algorithm"`
	PriorityEnabled bool   `json:"priority_enabled"`
	PriorityNodes   int    `json:"priority_nodes"`
}

type Node struct {
	CloseOnDeath  bool `json:"close_on_death"`
	RetryFailTime int  `json:"retry_fail_time"`
}

type PoolInfo struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

type SMTP struct {
	SendStarttls bool `json:"send_starttls"`
}

type SSL struct {
	ClientAuth           bool          `json:"client_auth"`
	CommonNameMatch      []interface{} `json:"common_name_match"`
	EllipticCurves       []interface{} `json:"elliptic_curves"`
	Enable               bool          `json:"enable"`
	Enhance              bool          `json:"enhance"`
	SendCloseAlerts      bool          `json:"send_close_alerts"`
	ServerName           bool          `json:"server_name"`
	Signature_algorithms string        `json:"signature_algorithms"`
	Ssl_ciphers          string        `json:"ssl_ciphers"`
	Ssl_support_ssl2     string        `json:"ssl_support_ssl2"`
	Ssl_support_ssl3     string        `json:"ssl_support_ssl3"`
	Ssl_support_tls1     string        `json:"ssl_support_tls1"`
	Ssl_support_tls1_1   string        `json:"ssl_support_tls1_1"`
	Ssl_support_tls1_2   string        `json:"ssl_support_tls1_2"`
	StrictVerify         bool          `json:"strict_verify"`
}

type TCP struct {
	Nagle bool `json:"nagle"`
}

type UDP struct {
	AcceptFrom     string `json:"accept_from"`
	AcceptFromMask string `json:"accept_from_mask"`
}

// Pool is the definition for a pool in stingray
type Pool struct {
	Basic *Basic `json:"basic"`

	AutoScaling *AutoScaling `json:"auto_scaling"`

	Connection                 *Connection                 `json:"connection"`
	DNSAutoscale               *DNSAutoscale               `json:"dns_autoscale"`
	FTP                        *FTP                        `json:"ftp"`
	HTTP                       *HTTP                       `json:"http"`
	KerberosProtocolTransition *KerberosProtocolTransition `json:"kerberos_protocol_transition"`
	LoadBalancing              *LoadBalancing              `json:"load_balancing"`
	Node                       *Node                       `json:"node"`
	SMTP                       *SMTP                       `json:"smtp"`
	SSL                        *SSL                        `json:"ssl"`
	TCP                        *TCP                        `json:"tcp"`
	UDP                        *UDP                        `json:"udp"`
}

// ListPools retrieves an array of the pool names currently registered in stingray
func (r *marathonClient) ListPools() ([]string, error) {
	var wrapper struct {
		Pools []PoolInfo `json:"children"`
	}
	if err := r.apiGet(marathonAPIPools, nil, &wrapper); err != nil {
		return nil, err
	}

	var list []string
	for _, pool := range wrapper.Pools {
		list = append(list, pool.Name)
	}

	return list, nil
}

// Application retrieves the application configuration from marathon
// 		name: 		the id used to identify the application
func (r *marathonClient) Pool(name string) (*Pool, error) {
	var wrapper struct {
		Pool *Pool `json:"properties"`
	}

	if err := r.apiGet(buildURI(name), nil, &wrapper); err != nil {
		return nil, err
	}

	return wrapper.Pool, nil
}

// NewDockerApplication creates a default docker application
// func NewDockerApplication() *Application {
// 	application := new(Application)
// 	return application
// }

// String returns the json representation of this application
// func (r *Application) String() string {
// 	s, err := json.MarshalIndent(r, "", "  ")
// 	if err != nil {
// 		return fmt.Sprintf(`{"error": "error decoding type into json: %s"}`, err)
// 	}
//
// 	return string(s)
// }

// CreateApplication creates a new application in Marathon
// 		application:		the structure holding the application configuration
// func (r *marathonClient) CreateApplication(application *Application) (*Application, error) {
// 	result := new(Application)
// 	if err := r.apiPost(marathonAPIApps, application, result); err != nil {
// 		return nil, err
// 	}
//
// 	return result, nil
// }

// DeleteApplication deletes an application from marathon
// 		name: 		the id used to identify the application
//		force:		used to force the delete operation in case of blocked deployment
// func (r *marathonClient) DeleteApplication(name string, force bool) (*DeploymentID, error) {
// 	uri := buildURIWithForceParam(name, force)
// 	// step: check of the application already exists
// 	deployID := new(DeploymentID)
// 	if err := r.apiDelete(uri, nil, deployID); err != nil {
// 		return nil, err
// 	}
//
// 	return deployID, nil
// }

// UpdateApplication updates an application in Marathon
// 		application:		the structure holding the application configuration
// func (r *marathonClient) UpdateApplication(application *Application, force bool) (*DeploymentID, error) {
// 	result := new(DeploymentID)
// 	uri := buildURIWithForceParam(application.ID, force)
// 	if err := r.apiPut(uri, application, result); err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

func buildURI(path string) string {
	return fmt.Sprintf("%s/%s", marathonAPIPools, trimRootPath(path))
}
