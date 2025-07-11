package illumiocli

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetContainerWorkloadProfilesByContainerClusterHref_Success(t *testing.T) {
	mockResponse := `[
		{
			"href": "/orgs/10/container_clusters/11111111-1111-1111-1111-111111111111/container_workload_profiles/22222222-2222-2222-2222-222222222222",
			"namespace": "example-namespace",
			"description": "",
			"enforcement_mode": "full",
			"visibility_level": "flow_summary",
			"managed": true,
			"assign_labels": [
				{
					"href": "/orgs/10/labels/1001"
				},
				{
					"href": "/orgs/10/labels/1002"
				},
				{
					"href": "/orgs/10/labels/1003"
				},
				{
					"href": "/orgs/10/labels/1004"
				}
			],
			"labels": [
				{
					"key": "loc",
					"assignment": {
						"href": "/orgs/10/labels/1001",
						"value": "location"
					}
				},
				{
					"key": "env",
					"assignment": {
						"href": "/orgs/10/labels/1002",
						"value": "environment"
					}
				},
				{
					"key": "app",
					"assignment": {
						"href": "/orgs/10/labels/1003",
						"value": "application"
					}
				},
				{
					"key": "role",
					"assignment": {
						"href": "/orgs/10/labels/1004",
						"value": "role"
					}
				}
			],
			"linked": true,
			"created_at": "2024-09-16T13:27:58.826Z",
			"created_by": {
				"href": "/orgs/10/container_clusters/11111111-1111-1111-1111-111111111111"
			},
			"updated_at": "2025-06-10T13:07:29.634Z",
			"updated_by": {
				"href": "/users/888"
			}
		},
		{
			"href": "/orgs/10/container_clusters/11111111-1111-1111-1111-111111111111/container_workload_profiles/33333333-3333-3333-3333-333333333333",
			"namespace": "example-namespace",
			"description": "",
			"enforcement_mode": "visibility_only",
			"visibility_level": "flow_summary",
			"managed": true,
			"assign_labels": [
				{
					"href": "/orgs/10/labels/1001"
				},
				{
					"href": "/orgs/10/labels/1002"
				},
				{
					"href": "/orgs/10/labels/1003"
				},
				{
					"href": "/orgs/10/labels/1004"
				}
			],
			"labels": [
				{
					"key": "loc",
					"assignment": {
						"href": "/orgs/10/labels/1001",
						"value": "location"
					}
				},
				{
					"key": "env",
					"assignment": {
						"href": "/orgs/10/labels/1002",
						"value": "environment"
					}
				},
				{
					"key": "app",
					"assignment": {
						"href": "/orgs/10/labels/1003",
						"value": "application"
					}
				},
				{
					"key": "role",
					"assignment": {
						"href": "/orgs/10/labels/1004",
						"value": "role"
					}
				}
			],
			"linked": true,
			"created_at": "2024-09-16T13:27:58.759Z",
			"created_by": {
				"href": "/orgs/10/container_clusters/11111111-1111-1111-1111-111111111111"
			},
			"updated_at": "2025-06-10T13:07:49.078Z",
			"updated_by": {
				"href": "/users/888"
			}
		},
		{
			"href": "/orgs/10/container_clusters/11111111-1111-1111-1111-111111111111/container_workload_profiles/44444444-4444-4444-4444-444444444444",
			"namespace": "default",
			"description": "",
			"enforcement_mode": "full",
			"visibility_level": "flow_summary",
			"managed": true,
			"assign_labels": [
				{
					"href": "/orgs/10/labels/1001"
				},
				{
					"href": "/orgs/10/labels/1002"
				},
				{
					"href": "/orgs/10/labels/1003"
				},
				{
					"href": "/orgs/10/labels/1004"
				}
			],
			"labels": [
				{
					"key": "loc",
					"assignment": {
						"href": "/orgs/10/labels/1001",
						"value": "location"
					}
				},
				{
					"key": "env",
					"assignment": {
						"href": "/orgs/10/labels/1002",
						"value": "environment"
					}
				},
				{
					"key": "app",
					"assignment": {
						"href": "/orgs/10/labels/1003",
						"value": "application"
					}
				},
				{
					"key": "role",
					"assignment": {
						"href": "/orgs/10/labels/1004",
						"value": "role"
					}
				}
			],
			"linked": true,
			"created_at": "2024-09-16T13:27:58.627Z",
			"created_by": {
				"href": "/orgs/10/container_clusters/11111111-1111-1111-1111-111111111111"
			},
			"updated_at": "2025-06-10T13:07:29.982Z",
			"updated_by": {
				"href": "/users/888"
			}
		},
		{
			"href": "/orgs/10/container_clusters/11111111-1111-1111-1111-111111111111/container_workload_profiles/55555555-5555-5555-5555-555555555555",
			"namespace": "example-namespace",
			"description": "",
			"enforcement_mode": "full",
			"visibility_level": "flow_summary",
			"managed": true,
			"assign_labels": [
				{
					"href": "/orgs/10/labels/1001"
				},
				{
					"href": "/orgs/10/labels/1002"
				},
				{
					"href": "/orgs/10/labels/1003"
				},
				{
					"href": "/orgs/10/labels/1004"
				}
			],
			"labels": [
				{
					"key": "loc",
					"assignment": {
						"href": "/orgs/10/labels/1001",
						"value": "location"
					}
				},
				{
					"key": "env",
					"assignment": {
						"href": "/orgs/10/labels/1002",
						"value": "environment"
					}
				},
				{
					"key": "role",
					"assignment": {
						"href": "/orgs/10/labels/1003",
						"value": "role"
					}
				},
				{
					"key": "app",
					"assignment": {
						"href": "/orgs/10/labels/1004",
						"value": "application"
					}
				}
			],
			"linked": true,
			"created_at": "2024-09-16T13:27:58.617Z",
			"created_by": {
				"href": "/orgs/10/container_clusters/11111111-1111-1111-1111-111111111111"
			},
			"updated_at": "2025-06-10T13:07:30.616Z",
			"updated_by": {
				"href": "/users/888"
			}
		}
	]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}
		if r.URL.Path != "/api/v1/orgs/1/container_clusters/test-cluster-id/container_workload_profiles" {
			t.Errorf("Expected path /api/v1/orgs/1/container_clusters/test-cluster-id/container_workload_profiles, got %s", r.URL.Path)
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

	profiles, err := pce.GetContainerWorkloadProfilesByContainerClusterHref("/orgs/1/container_clusters/test-cluster-id")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if profiles == nil {
		t.Fatal("Expected profiles to not be nil")
	}
	if len(profiles) != 4 {
		t.Fatalf("Expected 4 profiles, got %d", len(profiles))
	}

	// Test first profile
	firstProfile := profiles[0]
	if firstProfile.Href != "/orgs/10/container_clusters/11111111-1111-1111-1111-111111111111/container_workload_profiles/22222222-2222-2222-2222-222222222222" {
		t.Errorf("Expected href '/orgs/10/container_clusters/11111111-1111-1111-1111-111111111111/container_workload_profiles/22222222-2222-2222-2222-222222222222', got %s", firstProfile.Href)
	}
	if firstProfile.Namespace != "example-namespace" {
		t.Errorf("Expected namespace 'example-namespace', got %s", firstProfile.Namespace)
	}
	if firstProfile.EnforcementMode != "full" {
		t.Errorf("Expected enforcement_mode 'full', got %s", firstProfile.EnforcementMode)
	}
	if firstProfile.VisibilityLevel != "flow_summary" {
		t.Errorf("Expected visibility_level 'flow_summary', got %s", firstProfile.VisibilityLevel)
	}
	if firstProfile.Managed != true {
		t.Errorf("Expected managed to be true, got %v", firstProfile.Managed)
	}
	if firstProfile.Linked != true {
		t.Errorf("Expected linked to be true, got %v", firstProfile.Linked)
	}
	if len(firstProfile.AssignLabels) != 4 {
		t.Errorf("Expected 4 assign labels, got %d", len(firstProfile.AssignLabels))
	}
	if len(firstProfile.Labels) != 4 {
		t.Errorf("Expected 4 labels, got %d", len(firstProfile.Labels))
	}

	// Test second profile (visibility_only mode)
	secondProfile := profiles[1]
	if secondProfile.EnforcementMode != "visibility_only" {
		t.Errorf("Expected enforcement_mode 'visibility_only', got %s", secondProfile.EnforcementMode)
	}

	// Test third profile (default namespace)
	thirdProfile := profiles[2]
	if thirdProfile.Namespace != "default" {
		t.Errorf("Expected namespace 'default', got %s", thirdProfile.Namespace)
	}
}

func TestGetContainerWorkloadProfilesByContainerClusterHref_Failure_403(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}
		if r.URL.Path != "/api/v1/orgs/1/container_clusters/test-cluster-id/container_workload_profiles" {
			t.Errorf("Expected path /api/v1/orgs/1/container_clusters/test-cluster-id/container_workload_profiles, got %s", r.URL.Path)
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

	profiles, err := pce.GetContainerWorkloadProfilesByContainerClusterHref("/orgs/1/container_clusters/test-cluster-id")

	if err == nil {
		t.Error("Expected an error due to 403 status, but got nil")
	}
	if profiles != nil {
		t.Errorf("Expected profiles to be nil, got %v", profiles)
	}
}

func TestPutContainerWorkloadProfilesByHref_Success(t *testing.T) {
	expectedRequestBody := `{
		"assign_labels": [
			{
				"href": "/orgs/10/labels/1001"
			},
			{
				"href": "/orgs/10/labels/1002"
			},
			{
				"href": "/orgs/10/labels/1003"
			},
			{
				"href": "/orgs/10/labels/1004"
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}
		if r.URL.Path != "/api/v1/orgs/10/container_clusters/11111111-1111-1111-1111-111111111111/container_workload_profiles/22222222-2222-2222-2222-222222222222" {
			t.Errorf("Expected path /api/v1/orgs/10/container_clusters/11111111-1111-1111-1111-111111111111/container_workload_profiles/22222222-2222-2222-2222-222222222222, got %s", r.URL.Path)
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

		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" && contentType != "application/json; charset=utf-8" {
			t.Errorf("Expected Content-Type to be application/json, got %s", contentType)
		}

		// Read and validate the request body
		body := make([]byte, r.ContentLength)
		_, err := r.Body.Read(body)
		if err != nil && err.Error() != "EOF" {
			t.Errorf("Error reading request body: %v", err)
		}

		// Compare the request body with expected (ignoring whitespace differences)
		var expectedJSON, actualJSON interface{}
		if err := json.Unmarshal([]byte(expectedRequestBody), &expectedJSON); err != nil {
			t.Errorf("Error unmarshaling expected JSON: %v", err)
		}
		if err := json.Unmarshal(body, &actualJSON); err != nil {
			t.Errorf("Error unmarshaling actual JSON: %v", err)
		}

		if !reflect.DeepEqual(expectedJSON, actualJSON) {
			t.Errorf("Request body mismatch.\nExpected: %s\nActual: %s", expectedRequestBody, string(body))
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	pce := &PCE{
		Host:    server.URL,
		User:    "testuser",
		Key:     "testkey",
		Version: "v1",
	}

	// Create the container workload profile with assign_labels
	cwp := &ContainerWorkloadProfile{
		AssignLabels: []Href{
			{Href: "/orgs/10/labels/1001"},
			{Href: "/orgs/10/labels/1002"},
			{Href: "/orgs/10/labels/1003"},
			{Href: "/orgs/10/labels/1004"},
		},
	}

	href := "/orgs/10/container_clusters/11111111-1111-1111-1111-111111111111/container_workload_profiles/22222222-2222-2222-2222-222222222222"

	result, err := pce.PutContainerWorkloadProfileByHref(href, cwp)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("Expected result to not be nil")
	}
}

func TestPutContainerWorkloadProfilesByHref_Failure_404_ResourceNotFound(t *testing.T) {

}

func TestPutContainerWorkloadProfilesByHref_Failure_406_InvalidUUID(t *testing.T) {

}
