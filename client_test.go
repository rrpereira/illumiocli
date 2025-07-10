package illumiocli

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	host := "https://test.example.com"
	user := "testuser"
	key := "testkey"
	version := "v1"

	client := NewClient(host, user, key, version)

	if client.Host != host {
		t.Errorf("Expected Host to be %s, got %s", host, client.Host)
	}
	if client.User != user {
		t.Errorf("Expected User to be %s, got %s", user, client.User)
	}
	if client.Key != key {
		t.Errorf("Expected Key to be %s, got %s", key, client.Key)
	}
	if client.Version != version {
		t.Errorf("Expected Version to be %s, got %s", version, client.Version)
	}
}
