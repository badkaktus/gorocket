package gorocket

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func NewTestServer(responseHandler http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(responseHandler)
}

func NewTestClientWithCustomHandler(t *testing.T, server *httptest.Server) *Client {
	client := NewClient(server.URL)

	if client.HTTPClient.Timeout != 5*time.Minute {
		t.Errorf("Expected timeout to be 5 minutes, got %v", client.HTTPClient.Timeout)
	}

	if client.baseURL != server.URL {
		t.Errorf("Expected base URL to be %s, got %s", server.URL, client.baseURL)
	}

	if client.apiVersion != "api/v1" {
		t.Errorf("Expected API version to be api/v1, got %s", client.apiVersion)
	}

	return client
}

func NewTestClient(t *testing.T) *Client {
	responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := NewTestServer(responseHandler)

	defer server.Close()

	client := NewClient(server.URL)

	if client.HTTPClient.Timeout != 5*time.Minute {
		t.Errorf("Expected timeout to be 5 minutes, got %v", client.HTTPClient.Timeout)
	}

	if client.baseURL != server.URL {
		t.Errorf("Expected base URL to be %s, got %s", server.URL, client.baseURL)
	}

	if client.apiVersion != "api/v1" {
		t.Errorf("Expected API version to be api/v1, got %s", client.apiVersion)
	}

	return client
}

func TestNewClient(t *testing.T) {
	NewTestClient(t)
}

func TestClientWithOptions(t *testing.T) {
	server := NewTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewWithOptions(
		server.URL,
		WithTimeout(5*time.Minute),
		WithUserID("user"),
		WithXToken("token"),
	)

	if client.HTTPClient.Timeout != 5*time.Minute {
		t.Errorf("Expected timeout to be 5 minutes, got %v", client.HTTPClient.Timeout)
	}

	if client.baseURL != server.URL {
		t.Errorf("Expected base URL to be %s, got %s", server.URL, client.baseURL)
	}

	if client.apiVersion != "api/v1" {
		t.Errorf("Expected API version to be api/v1, got %s", client.apiVersion)
	}
}
