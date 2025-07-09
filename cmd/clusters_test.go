package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetClusters_Success(t *testing.T) {
	mockResponse := `[
		{
			"href": "/orgs/1/container_clusters/11111111-1111-1111-1111-111111111111",
			"pce_fqdn": null,
			"name": "demo-k8s-cluster",
			"description": "",
			"manager_type": "Kubernetes v1.22.2",
			"network_type": "host_switch_mode",
			"last_connected": "2025-03-09T18:45:29.258Z",
			"kubelink_version": "2.1.1.1847bc",
			"clas_mode": false,
			"cluster_mode": "legacy",
			"online": false,
			"nodes": [],
			"container_runtime": null,
			"errors": [
				{
					"error_type": "kubernetes_connection_failed"
				}
			],
			"caps": [
				"write"
			]
		},
		{
			"href": "/orgs/1/container_clusters/22222222-2222-2222-2222-222222222222",
			"pce_fqdn": null,
			"name": "demo-openshift-cluster",
			"description": "",
			"manager_type": "OpenShift v1.29.14+7cf4c05",
			"network_type": "host_switch_mode",
			"last_connected": "2025-03-09T18:45:29.258Z",
			"kubelink_version": "5.5.1-5",
			"clas_mode": false,
			"cluster_mode": "legacy",
			"online": true,
			"nodes": [
				{
					"name": "node001",
					"pod_subnet": "8.8.0.0/16"
				},
				{
					"name": "node002",
					"pod_subnet": "8.8.0.0/16"
				},
				{
					"name": "node003",
					"pod_subnet": "8.8.0.0/16"
				},
				{
					"name": "node004",
					"pod_subnet": "8.8.0.0/16"
				},
				{
					"name": "node005",
					"pod_subnet": "8.8.0.0/16"
				},
				{
					"name": "node006",
					"pod_subnet": "8.8.0.0/16"
				},
				{
					"name": "node007",
					"pod_subnet": "8.8.0.0/16"
				},
				{
					"name": "node008",
					"pod_subnet": "8.8.0.0/16"
				},
				{
					"name": "node009",
					"pod_subnet": "8.8.0.0/16"
				}
			],
			"container_runtime": "cri-o",
			"errors": [],
			"caps": [
				"write"
			]
		}
	]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}
		if r.URL.Path != "/orgs/1/container_clusters" {
			t.Errorf("Expected path /orgs/1/container_clusters, got %s", r.URL.Path)
		}

		username, password, ok := r.BasicAuth()
		if !ok {
			t.Error("Expected basic auth to be present")
		}
		if username != "testuser" {
			t.Errorf("Expected username 'testuser', got %s", username)
		}
		if password != "testkey" {
			t.Errorf("Expected password 'testkey', got %s", password)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := &Client{
		Host: server.URL,
		User: "testuser",
		Key:  "testkey",
	}

	clusters, err := client.GetClusters()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if clusters == nil {
		t.Fatal("Expected clusters to not be nil")
	}
	if len(clusters) != 2 {
		t.Fatalf("Expected 2 clusters, got %d", len(clusters))
	}

	k8sCluster := clusters[0]
	if k8sCluster.Href != "/orgs/1/container_clusters/11111111-1111-1111-1111-111111111111" {
		t.Errorf("Expected href '/orgs/1/container_clusters/11111111-1111-1111-1111-111111111111', got %s", k8sCluster.Href)
	}
	if k8sCluster.PceFqdn != nil {
		t.Errorf("Expected PceFqdn to be nil, got %v", k8sCluster.PceFqdn)
	}
	if k8sCluster.Name != "demo-k8s-cluster" {
		t.Errorf("Expected name 'demo-k8s-cluster', got %s", k8sCluster.Name)
	}
	if k8sCluster.ManagerType != "Kubernetes v1.22.2" {
		t.Errorf("Expected manager type 'Kubernetes v1.22.2', got %s", k8sCluster.ManagerType)
	}
	if k8sCluster.Online != false {
		t.Errorf("Expected online to be false, got %t", k8sCluster.Online)
	}
	if len(k8sCluster.Nodes) != 0 {
		t.Errorf("Expected 0 nodes, got %d", len(k8sCluster.Nodes))
	}
	if len(k8sCluster.Errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(k8sCluster.Errors))
	}
	if len(k8sCluster.Errors) > 0 && k8sCluster.Errors[0].ErrorType != "kubernetes_connection_failed" {
		t.Errorf("Expected error type 'kubernetes_connection_failed', got %s", k8sCluster.Errors[0].ErrorType)
	}

	openshiftCluster := clusters[1]
	if openshiftCluster.Name != "demo-openshift-cluster" {
		t.Errorf("Expected name 'demo-openshift-cluster', got %s", openshiftCluster.Name)
	}
	if openshiftCluster.ManagerType != "OpenShift v1.29.14+7cf4c05" {
		t.Errorf("Expected manager type 'OpenShift v1.29.14+7cf4c05', got %s", openshiftCluster.ManagerType)
	}
	if openshiftCluster.Online != true {
		t.Errorf("Expected online to be true, got %t", openshiftCluster.Online)
	}
	if len(openshiftCluster.Nodes) != 9 {
		t.Errorf("Expected 9 nodes, got %d", len(openshiftCluster.Nodes))
	}
	if openshiftCluster.ContainerRuntime == nil || *openshiftCluster.ContainerRuntime != "cri-o" {
		t.Errorf("Expected container runtime 'cri-o', got %v", openshiftCluster.ContainerRuntime)
	}
	if len(openshiftCluster.Errors) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(openshiftCluster.Errors))
	}

	for i, node := range openshiftCluster.Nodes {
		expectedName := fmt.Sprintf("node%03d", i+1)
		if node.Name != expectedName {
			t.Errorf("Expected node name %s, got %s", expectedName, node.Name)
		}
		if node.PodSubnet != "8.8.0.0/16" {
			t.Errorf("Expected pod subnet '8.8.0.0/16', got %s", node.PodSubnet)
		}
	}
}

func TestGetClusters_Failure_WrongOrg(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}
		if r.URL.Path != "/orgs/1/container_clusters" {
			t.Errorf("Expected path /orgs/1/container_clusters, got %s", r.URL.Path)
		}

		username, password, ok := r.BasicAuth()
		if !ok {
			t.Error("Expected basic auth to be present")
		}
		if username != "testuser" {
			t.Errorf("Expected username 'testuser', got %s", username)
		}
		if password != "testkey" {
			t.Errorf("Expected password 'testkey', got %s", password)
		}

		w.WriteHeader(http.StatusForbidden)
	}))
	defer server.Close()

	client := &Client{
		Host: server.URL,
		User: "testuser",
		Key:  "testkey",
	}

	clusters, err := client.GetClusters()

	if err == nil {
		t.Error("Expected an error due to 403 status, but got nil")
	}
	if clusters != nil {
		t.Errorf("Expected clusters to be nil, got %v", clusters)
	}
}
