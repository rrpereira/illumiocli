package illumiocli

import (
	"fmt"
)

// Error represents an API error
type Error struct {
	ErrorType string `json:"error_type"`
}

// Node represents a cluster node
type Node struct {
	Name      string `json:"name"`
	PodSubnet string `json:"pod_subnet"`
}

// Cluster represents an Illumio container cluster
type Cluster struct {
	Href             string   `json:"href"`
	PceFqdn          *string  `json:"pce_fqdn"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	ManagerType      string   `json:"manager_type"`
	NetworkType      string   `json:"network_type"`
	LastConnected    string   `json:"last_connected"`
	KubelinkVersion  string   `json:"kubelink_version"`
	ClasMode         bool     `json:"clas_mode"`
	ClusterMode      string   `json:"cluster_mode"`
	Online           bool     `json:"online"`
	Nodes            []Node   `json:"nodes"`
	ContainerRuntime *string  `json:"container_runtime"`
	Errors           []Error  `json:"errors"`
	Caps             []string `json:"caps"`
}

func (c *PCE) GetClusters() ([]Cluster, error) {
	var clusters []Cluster

	r := c.newRestyClient()

	resp, err := r.R().
		SetResult(&clusters).
		Get("/orgs/1/container_clusters")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("API error: status %d - %s", resp.StatusCode(), resp.Status())
	}

	return clusters, nil
}
