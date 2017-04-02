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

type PoolBasic struct {
	BandwidthClass                string      `json:"bandwidth_class,omitempty"`
	FailurePool                   string      `json:"failure_pool,omitempty"`
	MaxConnectionAttempts         int         `json:"max_connection_attempts,omitempty"`
	MaxIdleConnectionsPernode     int         `json:"max_idle_connections_pernode,omitempty"`
	MaxTimedOutConnectionAttempts int         `json:"max_timed_out_connection_attempts,omitempty"`
	Monitors                      []string    `json:"monitors,omitempty,omitempty"`
	NodeCloseWithRst              bool        `json:"node_close_with_rst,omitempty"`
	NodeConnectionAttempts        int         `json:"node_connection_attempts,omitempty"`
	NodeDeleteBehavior            string      `json:"node_delete_behavior,omitempty"`
	NodeDrainToDeleteTimeout      int         `json:"node_drain_to_delete_timeout,omitempty"`
	NodesTable                    *[]NodeItem `json:"nodes_table,omitempty"`
	Note                          string      `json:"note,omitempty"`
	PassiveMonitoring             bool        `json:"passive_monitoring,omitempty"`
	PersistenceClass              string      `json:"persistence_class,omitempty"`
	Transparent                   bool        `json:"transparent,omitempty"`
}

type PoolAutoScaling struct {
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

type PoolConnection struct {
	MaxConnectTime        int `json:"max_connect_time,omitempty"`
	MaxConnectionsPerNode int `json:"max_connections_per_node,omitempty"`
	MaxQueueSize          int `json:"max_queue_size,omitempty"`
	MaxReplyTime          int `json:"max_reply_time,omitempty"`
	QueueTimeout          int `json:"queue_timeout,omitempty"`
}

type PoolDNSAutoscale struct {
	Enabled   bool     `json:"enabled,omitempty"`
	Hostnames []string `json:"hostnames,omitempty"`
	Port      int      `json:"port,omitempty"`
}
type PoolFTP struct {
	SupportRFC2428 bool `json:"support_rfc_2428,omitempty"`
}
type PoolHTTP struct {
	Keepalive              bool `json:"keepalive,omitempty"`
	KeepaliveNonIdempotent bool `json:"keepalive_non_idempotent,omitempty"`
}
type PoolKerberosProtocolTransition struct {
	Principal string `json:"principal,omitempty"`
	Target    string `json:"target,omitempty"`
}
type PoolLoadBalancing struct {
	Algorithm       string `json:"algorithm,omitempty"`
	PriorityEnabled bool   `json:"priority_enabled,omitempty"`
	PriorityNodes   int    `json:"priority_nodes,omitempty"`
}

type PoolNode struct {
	CloseOnDeath  bool `json:"close_on_death,omitempty"`
	RetryFailTime int  `json:"retry_fail_time,omitempty"`
}

type PoolSMTP struct {
	SendStarttls bool `json:"send_starttls,omitempty"`
}

type PoolSSL struct {
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

type PoolTCP struct {
	Nagle bool `json:"nagle,omitempty"`
}

type PoolUDP struct {
	AcceptFrom     string `json:"accept_from,omitempty"`
	AcceptFromMask string `json:"accept_from_mask,omitempty"`
}

// Pool is the definition for a pool in VTM
type Pool struct {
	Basic                      *PoolBasic                      `json:"basic,omitempty"`
	AutoScaling                *PoolAutoScaling                `json:"auto_scaling,omitempty"`
	Connection                 *PoolConnection                 `json:"connection,omitempty"`
	DNSAutoscale               *PoolDNSAutoscale               `json:"dns_autoscale,omitempty"`
	FTP                        *PoolFTP                        `json:"ftp,omitempty,omitempty"`
	HTTP                       *PoolHTTP                       `json:"http,omitempty"`
	KerberosProtocolTransition *PoolKerberosProtocolTransition `json:"kerberos_protocol_transition,omitempty"`
	LoadBalancing              *PoolLoadBalancing              `json:"load_balancing,omitempty"`
	Node                       *PoolNode                       `json:"node,omitempty"`
	SMTP                       *PoolSMTP                       `json:"smtp,omitempty"`
	SSL                        *PoolSSL                        `json:"ssl,omitempty"`
	TCP                        *PoolTCP                        `json:"tcp,omitempty"`
	UDP                        *PoolUDP                        `json:"udp,omitempty"`
}

// NewPool creates a default pool
func NewPool() *Pool {
	pool := new(Pool)
	return pool
}

// String returns the json representation of this pool
func (p *Pool) String() string {
	s, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error": "error decoding type into json: %s"}`, err)
	}
	return string(s)
}

// AddNode adds one or more nodes to the pool
//      arguments:  the pool(s) you are adding
func (p *Pool) AddNode(nodes ...NodeItem) *Pool {
	if p.Basic == nil || p.Basic.NodesTable == nil {
		p.EmptyNodes()
	}
	nodeTable := *p.Basic.NodesTable
	nodeTable = append(nodeTable, nodes...)
	p.Basic.NodesTable = &nodeTable
	return p
}

// EmptyNodes explicitly empties NodesTable
func (p *Pool) EmptyNodes() *Pool {
	if p.Basic == nil {
		p.Basic = &PoolBasic{}
	}
	p.Basic.NodesTable = &[]NodeItem{}
	return p
}

func (p *Pool) Nodes() *[]NodeItem {
	if p.Basic == nil {
		return nil
	}
	return p.Basic.NodesTable
}

type poolWrapper struct {
	Pool *Pool `json:"properties"`
}

// ListPools retrieves an array of the pool names currently registered in VTM
func (c *vtmClient) ListPools() ([]string, error) {
	var wrapper struct {
		Pools []struct {
			Name string `json:"name"`
			Href string `json:"href"`
		} `json:"children"`
	}
	if err := c.apiGet(vtmAPIPools, nil, &wrapper); err != nil {
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
func (c *vtmClient) Pool(name string) (*Pool, error) {
	result := new(poolWrapper)
	if err := c.apiGet(buildURI(vtmAPIPools, name), nil, &result); err != nil {
		return nil, err
	}
	return result.Pool, nil
}

// PoolExists checks if pool exists in VTM using HEAD request
// 		name: 		the id used to identify the pool
func (r *vtmClient) PoolExists(name string) (bool, error) {
	if err := r.apiHead(buildURI(vtmAPIPools, name), nil, nil); err != nil {
		return false, err
	}
	return true, nil
}

// CreatePool creates a new pool in VTM
// 		pool:		the structure holding the pool configuration
func (c *vtmClient) CreatePool(name string, pool *Pool) (*Pool, error) {
	wrapper := new(poolWrapper)
	wrapper.Pool = pool
	result := new(poolWrapper)
	if err := c.apiPut(buildURI(vtmAPIPools, name), wrapper, &result); err != nil {
		return nil, err
	}
	return result.Pool, nil
}

// DeletePool deletes an application from VTM
// 		name: 		the id used to identify the pool
func (c *vtmClient) DeletePool(name string) error {
	if err := c.apiDelete(buildURI(vtmAPIPools, name), nil, nil); err != nil {
		return err
	}
	return nil
}
