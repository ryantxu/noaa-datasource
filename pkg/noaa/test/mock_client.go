package test

import (
	"context"
	"net/http"
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/experimental"
)

type mockClient struct {
	filePath string
}

func NewMockClient(filePath string) experimental.Client {
	return &mockClient{filePath: filePath}
}

// Fetch performs an HTTP GET and returns the body as []byte to prep for marshalling.
func (c *mockClient) Fetch(ctx context.Context, uriPath, uriQuery string) ([]byte, error) {
	return os.ReadFile(c.filePath)
}

// Get performs an HTTP GET and returns the response.
// This can be used directly from resource calls that don't need to marshal the data
func (c *mockClient) Get(ctx context.Context, uriPath, uriQuery string) (*http.Response, error) {
	return nil, nil // NOT really supported!!!
}
