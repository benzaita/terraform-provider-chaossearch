package client

import (
	"context"
)

type Client struct {
	config *Configuration
}

func NewClient(config *Configuration) *Client {
	return &Client{
		config: config,
	}
}

func (client *Client) ListBuckets(ctx context.Context) ([]map[string]interface{}, error) {
	// path := "/V1/"

	result := []map[string]interface{}{
		{
			"name": "foo",
			"id":   123,
		},
	}

	return result, nil
}
