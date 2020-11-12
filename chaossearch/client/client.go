package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

type Client struct {
	config     *Configuration
	httpClient *http.Client
}

func NewClient(config *Configuration) *Client {
	return &Client{
		config:     config,
		httpClient: http.DefaultClient,
	}
}

func (client *Client) ListBuckets(ctx context.Context) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/V1/", client.config.URL)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	credentials := credentials.NewStaticCredentials(client.config.AccessKeyID, client.config.SecretAccessKey, "")
	_, err = v4.NewSigner(credentials).Sign(req, nil, client.config.AWSServiceName, client.config.AWSRegion, time.Now())
	if err != nil {
		return nil, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Got a %d response from %s for request %+v", resp.StatusCode, url, req)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Body: %v", body)
	result := []map[string]interface{}{
		{
			"name": "foo",
			"id":   123,
		},
	}

	return result, nil
}
