package client

import (
	"context"
	"encoding/xml"
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

func (client *Client) ListBuckets(ctx context.Context) (*ListBucketsResponseModel, error) {
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

	var responseModel ListBucketsResponseModel
	if err := xml.Unmarshal(body, &responseModel); err != nil {
		return nil, err
	}

	return &responseModel, nil
}

type ListBucketsResponseModel struct {
	BucketsCollection BucketCollectionModel `xml:"Buckets"`
}

type BucketCollectionModel struct {
	Buckets []BucketModel `xml:"Bucket"`
}

type BucketModel struct {
	Name         string `xml:"Name"`
	CreationDate string `xml:"CreationDate"`
}
