package vtm

import (
	"encoding/json"
	"fmt"
)

type NodeItem struct {
	Node     string `json:"node,omitempty"`
	Priority int    `json:"priority,omitempty"`
	State    string `json:"state,omitempty"`
	Weight   int    `json:"weight,omitempty"`
}

type Basic struct {
	BandwidthClass                string     `json:"bandwidth_class,omitempty"`
	FailurePool                   string     `json:"failure_pool,omitempty"`
	MaxConnectionAttempts         int        `json:"max_connection_attempts,omitempty"`
	MaxIdleConnectionsPernode     int        `json:"max_idle_connections_pernode,omitempty"`
	MaxTimedOutConnectionAttempts int        `json:"max_timed_out_connection_attempts,omitempty"`
	Monitors                      []string   `json:"monitors,omitempty,omitempty"`
	NodeCloseWithRst              bool       `json:"node_close_with_rst,omitempty"`
	NodeConnectionAttempts        int        `json:"node_connection_attempts,omitempty"`
	NodeDeleteBehavior            string     `json:"node_delete_behavior,omitempty"`
	NodeDrainToDeleteTimeout      int        `json:"node_drain_to_delete_timeout,omitempty"`
	NodesTable                    []NodeItem `json:"nodes_table,omitempty"`
	Note                          string     `json:"note,omitempty"`
	PassiveMonitoring             bool       `json:"passive_monitoring,omitempty"`
	PersistenceClass              string     `json:"persistence_class,omitempty"`
	Transparent                   bool       `json:"transparent,omitempty"`
}

type AutoScaling struct {
	AddnodeDelaytime int      `json:"addnode_delaytime,omitempty"`
	CloudCredentials string   `json:"cloud_credentials,omitempty"`
	Cluster          string   `json:"cluster,omitempty"`
	DataCenter       string   `json:"data_center,omitempty"`
	DataStore        string   `json:"data_store,omitempty"`
	Enabled          bool     `json:"enabled,omitempty"`
	External         bool     `json:"external,omitempty"`
	Hysteresis       int      `json:"hysteresis,omitempty"`
	Imageid          string   `json:"imageid,omitempty"`
	IpsToUse         string   `json:"ips_to_use,omitempty"`
	LastNodeIdleTime int      `json:"last_node_idle_time,omitempty"`
	MaxNodes         int      `json:"max_nodes,omitempty"`
	MinNodes         int      `json:"min_nodes,omitempty"`
	Name             string   `json:"name,omitempty"`
	Port             int      `json:"port,omitempty"`
	Refractory       int      `json:"refractory,omitempty"`
	ResponseTime     int      `json:"response_time,omitempty"`
	ScaleDownLevel   int      `json:"scale_down_level,omitempty"`
	ScaleUpLevel     int      `json:"scale_up_level,omitempty"`
	Securitygroupids []string `json:"securitygroupids,omitempty"`
	SizeID           string   `json:"size_id,omitempty"`
	Subnetids        []string `json:"subnetids,omitempty"`
}

type Connection struct {
	MaxConnectTime        int `json:"max_connect_time,omitempty"`
	MaxConnectionsPerNode int `json:"max_connections_per_node,omitempty"`
	MaxQueueSize          int `json:"max_queue_size,omitempty"`
	MaxReplyTime          int `json:"max_reply_time,omitempty"`
	QueueTimeout          int `json:"queue_timeout,omitempty"`
}

type DNSAutoscale struct {
	Enabled   bool     `json:"enabled,omitempty"`
	Hostnames []string `json:"hostnames"`
	Port      int      `json:"port,omitempty"`
}
type FTP struct {
	SupportRFC2428 bool `json:"support_rfc_2428,omitempty"`
}
type HTTP struct {
	Keepalive              bool `json:"keepalive,omitempty"`
	KeepaliveNonIdempotent bool `json:"keepalive_non_idempotent,omitempty"`
}
type KerberosProtocolTransition struct {
	Principal string `json:"principal,omitempty"`
	Target    string `json:"target,omitempty"`
}
type LoadBalancing struct {
	Algorithm       string `json:"algorithm,omitempty"`
	PriorityEnabled bool   `json:"priority_enabled,omitempty"`
	PriorityNodes   int    `json:"priority_nodes,omitempty"`
}

type Node struct {
	CloseOnDeath  bool `json:"close_on_death,omitempty"`
	RetryFailTime int  `json:"retry_fail_time,omitempty"`
}

type SMTP struct {
	SendStarttls bool `json:"send_starttls,omitempty"`
}

type SSL struct {
	ClientAuth          bool     `json:"client_auth,omitempty"`
	CommonNameMatch     []string `json:"common_name_match,omitempty"`
	EllipticCurves      []string `json:"elliptic_curves,omitempty"`
	Enable              bool     `json:"enable,omitempty"`
	Enhance             bool     `json:"enhance,omitempty"`
	SendCloseAlerts     bool     `json:"send_close_alerts,omitempty"`
	ServerName          bool     `json:"server_name,omitempty"`
	SignatureAlgorithms string   `json:"signature_algorithms,omitempty"`
	SSLCiphers          string   `json:"ssl_ciphers,omitempty"`
	SSLSupportSSL2      string   `json:"ssl_support_ssl2,omitempty"`
	SSLSupportSSL3      string   `json:"ssl_support_ssl3,omitempty"`
	SSLSupportTLS1      string   `json:"ssl_support_tls1,omitempty"`
	SSLSupportTLS11     string   `json:"ssl_support_tls1_1,omitempty"`
	SSLSupportTLS12     string   `json:"ssl_support_tls1_2,omitempty"`
	StrictVerify        bool     `json:"strict_verify,omitempty"`
}

type TCP struct {
	Nagle bool `json:"nagle,omitempty"`
}

type UDP struct {
	AcceptFrom     string `json:"accept_from,omitempty"`
	AcceptFromMask string `json:"accept_from_mask,omitempty"`
}

// Pool is the definition for a pool in VTM
type Pool struct {
	Basic                      *Basic                      `json:"basic,omitempty"`
	AutoScaling                *AutoScaling                `json:"auto_scaling,omitempty"`
	Connection                 *Connection                 `json:"connection,omitempty"`
	DNSAutoscale               *DNSAutoscale               `json:"dns_autoscale,omitempty"`
	FTP                        *FTP                        `json:"ftp,omitempty,omitempty"`
	HTTP                       *HTTP                       `json:"http,omitempty"`
	KerberosProtocolTransition *KerberosProtocolTransition `json:"kerberos_protocol_transition,omitempty"`
	LoadBalancing              *LoadBalancing              `json:"load_balancing,omitempty"`
	Node                       *Node                       `json:"node,omitempty"`
	SMTP                       *SMTP                       `json:"smtp,omitempty"`
	SSL                        *SSL                        `json:"ssl,omitempty"`
	TCP                        *TCP                        `json:"tcp,omitempty"`
	UDP                        *UDP                        `json:"udp,omitempty"`
}

type poolWrapper struct {
	Pool *Pool `json:"properties"`
}

// ListPools retrieves an array of the pool names currently registered in VTM
func (r *vtmClient) ListPools() ([]string, error) {
	var wrapper struct {
		Pools []struct {
			Name string `json:"name"`
			Href string `json:"href"`
		} `json:"children"`
	}
	if err := r.apiGet(vtmAPIPools, nil, &wrapper); err != nil {
		return nil, err
	}

	var list []string
	for _, pool := range wrapper.Pools {
		list = append(list, pool.Name)
	}

	return list, nil
}

// Pool retrieves the pool configuration from VTM
// 		name: 		the id used to identify the pool
func (r *vtmClient) Pool(name string) (*Pool, error) {
	result := new(poolWrapper)
	if err := r.apiGet(buildURI(name), nil, &result); err != nil {
		return nil, err
	}
	return result.Pool, nil
}

// NewDockerApplication creates a default docker application
// func NewDockerApplication() *Application {
// 	application := new(Application)
// 	return application
// }

// String returns the json representation of this pool
func (r *Pool) String() string {
	s, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error": "error decoding type into json: %s"}`, err)
	}
	return string(s)
}

// CreatePool creates a new pool in VTM
// 		pool:		the structure holding the pool configuration
func (r *vtmClient) CreatePool(name string, pool *Pool) (*Pool, error) {
	wrapper := new(poolWrapper)
	wrapper.Pool = pool
	result := new(poolWrapper)
	if err := r.apiPut(buildURI(name), wrapper, &result); err != nil {
		return nil, err
	}
	return result.Pool, nil
}

// DeletePool deletes an application from VTM
// 		name: 		the id used to identify the pool
func (r *vtmClient) DeletePool(name string) error {
	if err := r.apiDelete(buildURI(name), nil, nil); err != nil {
		return err
	}
	return nil
}

func buildURI(path string) string {
	return fmt.Sprintf("%s/%s", vtmAPIPools, path)
}
