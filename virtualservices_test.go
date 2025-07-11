package illumiocli

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVirtualServicesByParameter_Success(t *testing.T) {
	mockResponse := `[
		{
			"href": "/orgs/1/sec_policy/active/virtual_services/12345678-1234-1234-1234-123456789abc",
			"created_at": "2023-01-01T00:00:00.000Z",
			"updated_at": "2024-01-01T00:00:00.000Z",
			"deleted_at": null,
			"created_by": {
				"href": "/users/100"
			},
			"updated_by": {
				"href": "/users/101"
			},
			"deleted_by": null,
			"name": "example-service.local",
			"description": "Sample virtual service",
			"pce_fqdn": null,
			"service_ports": [
				{
					"port": 443,
					"proto": 6
				}
			],
			"labels": [
				{
					"href": "/orgs/1/labels/10001",
					"key": "app",
					"value": "demo"
				},
				{
					"href": "/orgs/1/labels/10002",
					"key": "loc",
					"value": "test"
				},
				{
					"href": "/orgs/1/labels/10003",
					"key": "env",
					"value": "dev"
				}
			],
			"ip_overrides": [
				"192.0.2.10"
			],
			"apply_to": "host_only",
			"caps": [
				"write",
				"provision",
				"delete"
			],
			"service_addresses": []
		},
		{
			"href": "/orgs/1/sec_policy/active/virtual_services/abcdef12-3456-7890-abcd-ef1234567890",
			"created_at": "2025-05-04T13:16:22.959Z",
			"updated_at": "2025-01-09T05:51:54.756Z",
			"deleted_at": null,
			"created_by": {
				"href": "/orgs/1/container_clusters/00000000-0000-0000-0000-000000000000"
			},
			"updated_by": {
				"href": "/orgs/1/service_accounts/11111111-1111-1111-1111-111111111111",
				"name": "service-account"
			},
			"deleted_by": null,
			"name": "example-virtual-service",
			"description": "Sample description",
			"external_data_set": "00000000-0000-0000-0000-000000000000|example",
			"external_data_reference": "00000000-0000-0000-0000-000000000000",
			"pce_fqdn": null,
			"service_ports": [
				{
					"port": 8443,
					"proto": 6
				}
			],
			"labels": [
				{
					"href": "/orgs/1/labels/10001",
					"key": "loc",
					"value": "test"
				},
				{
					"href": "/orgs/1/labels/10002",
					"key": "env",
					"value": "dev"
				},
				{
					"href": "/orgs/1/labels/10003",
					"key": "role",
					"value": "sample"
				},
				{
					"href": "/orgs/1/labels/10004",
					"key": "app",
					"value": "demo"
				}
			],
			"ip_overrides": [],
			"apply_to": "host_only",
			"caps": [],
			"k8s_service_info": {
				"service_type": "ClusterIP",
				"external_traffic_policy": null,
				"node_ports": []
			},
			"service_addresses": [
				{
					"ip": "192.0.2.100",
					"network": {
						"href": "/orgs/1/networks/00000000-0000-0000-0000-000000000000",
						"name": "example network"
					},
					"port": 8443
				}
			]
		},
		{
			"href": "/orgs/1/sec_policy/active/virtual_services/22345678-1234-1234-1234-123456789abc",
			"created_at": "2024-12-12T12:12:30.118Z",
			"updated_at": "2025-11-11T05:54:11.801Z",
			"deleted_at": null,
			"created_by": {
				"href": "/orgs/1/container_clusters/00000000-0000-0000-0000-000000000000"
			},
			"updated_by": {
				"href": "/orgs/1/service_accounts/11111111-1111-1111-1111-111111111111",
				"name": "service-account"
			},
			"deleted_by": null,
			"name": "example-virtual-service-2",
			"description": "Sample description",
			"external_data_set": "00000000-0000-0000-0000-000000000000|example",
			"external_data_reference": "00000000-0000-0000-0000-000000000000",
			"pce_fqdn": null,
			"service_ports": [
				{
					"port": 8080,
					"proto": 6
				},
				{
					"port": 9040,
					"proto": 6
				}
			],
			"labels": [
				{
					"href": "/orgs/1/labels/10001",
					"key": "loc",
					"value": "test"
				},
				{
					"href": "/orgs/1/labels/10002",
					"key": "env",
					"value": "dev"
				},
				{
					"href": "/orgs/1/labels/10003",
					"key": "role",
					"value": "sample"
				},
				{
					"href": "/orgs/1/labels/10004",
					"key": "app",
					"value": "demo"
				}
			],
			"ip_overrides": [],
			"apply_to": "host_only",
			"caps": [],
			"k8s_service_info": {
				"service_type": "ClusterIP",
				"external_traffic_policy": null,
				"node_ports": []
			},
			"service_addresses": [
				{
					"ip": "192.0.2.101",
					"network": {
						"href": "/orgs/1/networks/00000000-0000-0000-0000-000000000000",
						"name": "example network"
					}
				}
			]
		},
		{
			"href": "/orgs/1/sec_policy/active/virtual_services/32345678-1234-1234-1234-123456789abc",
			"created_at": "2024-02-11T15:02:51.659Z",
			"updated_at": "2025-04-05T01:51:54.841Z",
			"deleted_at": null,
			"created_by": {
				"href": "/orgs/1/container_clusters/11111111-2222-3333-4444-555555555555"
			},
			"updated_by": {
				"href": "/orgs/1/service_accounts/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
				"name": "test-account"
			},
			"deleted_by": null,
			"name": "demo-metrics-service-example",
			"description": null,
			"external_data_set": "11111111-2222-3333-4444-555555555555|example",
			"external_data_reference": "11111111-2222-3333-4444-555555555555",
			"pce_fqdn": null,
			"service_ports": [
				{
					"port": 8443,
					"proto": 6
				}
			],
			"labels": [
				{
					"href": "/orgs/1/labels/10001",
					"key": "loc",
					"value": "nyc"
				},
				{
					"href": "/orgs/1/labels/10002",
					"key": "env",
					"value": "test"
				},
				{
					"href": "/orgs/1/labels/10003",
					"key": "role",
					"value": "demo"
				},
				{
					"href": "/orgs/1/labels/10004",
					"key": "app",
					"value": "sample"
				}
			],
			"ip_overrides": [],
			"apply_to": "host_only",
			"caps": [],
			"k8s_service_info": {
				"service_type": "ClusterIP",
				"external_traffic_policy": null,
				"node_ports": []
			},
			"service_addresses": [
				{
					"ip": "10.0.0.1",
					"network": {
						"href": "/orgs/1/networks/11111111-2222-3333-4444-555555555555",
						"name": "example container network"
					},
					"port": 8443
				}
			]
		}
	]
	`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		expectedPath := "/api/v1/orgs/1/sec_policy/active/virtual_services"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}

		queryParams := r.URL.Query()
		if queryParams.Get("name") != "example" {
			t.Errorf("Expected query parameter name=example, got app=%s", queryParams.Get("name"))
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

	pce := &PCE{
		Host:    server.URL,
		User:    "testuser",
		Key:     "testkey",
		Version: "v1",
	}

	virtualServices, err := pce.GetVirtualServicesByParameter("name", "example", "active")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if virtualServices == nil {
		t.Fatal("Expected virtual services to not be nil")
	}

	if len(virtualServices) != 4 {
		t.Fatalf("Expected 4 virtual services, got %d", len(virtualServices))
	}

	firstVS := virtualServices[0]
	if firstVS.Href != "/orgs/1/sec_policy/active/virtual_services/12345678-1234-1234-1234-123456789abc" {
		t.Errorf("Expected href '/orgs/1/sec_policy/active/virtual_services/12345678-1234-1234-1234-123456789abc', got %s", firstVS.Href)
	}
	if firstVS.Name != "example-service.local" {
		t.Errorf("Expected name 'example-service.local', got %s", firstVS.Name)
	}
	if firstVS.Description != "Sample virtual service" {
		t.Errorf("Expected description 'Sample virtual service', got %s", firstVS.Description)
	}
	if firstVS.ApplyTo != "host_only" {
		t.Errorf("Expected apply_to 'host_only', got %s", firstVS.ApplyTo)
	}

	if len(firstVS.ServicePorts) != 1 {
		t.Fatalf("Expected 1 service port, got %d", len(firstVS.ServicePorts))
	}
	if firstVS.ServicePorts[0].Port != 443 {
		t.Errorf("Expected port 443, got %d", firstVS.ServicePorts[0].Port)
	}
	if firstVS.ServicePorts[0].Proto != 6 {
		t.Errorf("Expected protocol 6, got %d", firstVS.ServicePorts[0].Proto)
	}

	if len(firstVS.Labels) != 3 {
		t.Fatalf("Expected 3 labels, got %d", len(firstVS.Labels))
	}

	firstLabel := firstVS.Labels[0]
	if firstLabel.Key != "app" {
		t.Errorf("Expected label key 'app', got %s", firstLabel.Key)
	}
	if firstLabel.Value != "demo" {
		t.Errorf("Expected label value 'demo', got %s", firstLabel.Value)
	}
	if firstLabel.Href != "/orgs/1/labels/10001" {
		t.Errorf("Expected label href '/orgs/1/labels/10001', got %s", firstLabel.Href)
	}

	if len(firstVS.IpOverrides) != 1 {
		t.Fatalf("Expected 1 IP override, got %d", len(firstVS.IpOverrides))
	}
	if firstVS.IpOverrides[0] != "192.0.2.10" {
		t.Errorf("Expected IP override '192.0.2.10', got %s", firstVS.IpOverrides[0])
	}

	if len(firstVS.Caps) != 3 {
		t.Fatalf("Expected 3 capabilities, got %d", len(firstVS.Caps))
	}
	expectedCaps := []string{"write", "provision", "delete"}
	for i, cap := range firstVS.Caps {
		if cap != expectedCaps[i] {
			t.Errorf("Expected capability '%s', got '%s'", expectedCaps[i], cap)
		}
	}

	secondVS := virtualServices[1]
	if secondVS.Name != "example-virtual-service" {
		t.Errorf("Expected name 'example-virtual-service', got %s", secondVS.Name)
	}
	if secondVS.ExternalDataSet != "00000000-0000-0000-0000-000000000000|example" {
		t.Errorf("Expected external_data_set '00000000-0000-0000-0000-000000000000|example', got %s", secondVS.ExternalDataSet)
	}
	if len(secondVS.ServicePorts) != 1 {
		t.Fatalf("Expected 1 service port, got %d", len(secondVS.ServicePorts))
	}
	if secondVS.ServicePorts[0].Port != 8443 {
		t.Errorf("Expected port 8443, got %d", secondVS.ServicePorts[0].Port)
	}
	if len(secondVS.ServiceAddresses) != 1 {
		t.Fatalf("Expected 1 service address, got %d", len(secondVS.ServiceAddresses))
	}
}

func TestGetVirtualServicesByParameter_Success_EmptyResponse(t *testing.T) {
	mockResponse := `[]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		expectedPath := "/api/v1/orgs/1/sec_policy/active/virtual_services"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}

		queryParams := r.URL.Query()
		if queryParams.Get("name") != "example" {
			t.Errorf("Expected query parameter name=example, got app=%s", queryParams.Get("name"))
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

	pce := &PCE{
		Host:    server.URL,
		User:    "testuser",
		Key:     "testkey",
		Version: "v1",
	}

	virtualServices, err := pce.GetVirtualServicesByParameter("name", "example", "active")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if virtualServices == nil {
		t.Fatal("Expected virtual services to not be nil")
	}
	if len(virtualServices) != 0 {
		t.Errorf("Expected 0 virtual services, got %d", len(virtualServices))
	}
}

func TestGetVirtualServicesByParameter_Failure_403(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		expectedPath := "/api/v1/orgs/1/sec_policy/active/virtual_services"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}

		queryParams := r.URL.Query()
		if queryParams.Get("name") != "example" {
			t.Errorf("Expected query parameter name=example, got name=%s", queryParams.Get("name"))
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

	pce := &PCE{
		Host:    server.URL,
		User:    "testuser",
		Key:     "testkey",
		Version: "v1",
	}

	virtualServices, err := pce.GetVirtualServicesByParameter("name", "example", "active")
	if err == nil {
		t.Fatal("Expected error for 403 response, got nil")
	}
	if virtualServices != nil {
		t.Errorf("Expected virtual services to be nil on error, got %v", virtualServices)
	}

	expectedError := "API error: status 403 - 403 Forbidden"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}
}

func TestGetVirtualServicesByParameter_Failure_400_WrongPversion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		expectedPath := "/api/v1/orgs/1/sec_policy/invalid/virtual_services"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}

		queryParams := r.URL.Query()
		if queryParams.Get("name") != "example" {
			t.Errorf("Expected query parameter name=example, got name=%s", queryParams.Get("name"))
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "pversion is invalid"}`))
	}))
	defer server.Close()

	pce := &PCE{
		Host:    server.URL,
		User:    "testuser",
		Key:     "testkey",
		Version: "v1",
	}

	virtualServices, err := pce.GetVirtualServicesByParameter("name", "example", "invalid")
	if err == nil {
		t.Fatal("Expected error for 400 response, got nil")
	}
	if virtualServices != nil {
		t.Errorf("Expected virtual services to be nil on error, got %v", virtualServices)
	}

	expectedError := `API error: status 400 - 400 Bad Request - {"error": "pversion is invalid"}`
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}
}
