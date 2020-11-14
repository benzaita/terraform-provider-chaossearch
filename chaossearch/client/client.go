package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
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

func (client *Client) signAndDo(req *http.Request, bodyAsBytes []byte) (*http.Response, error) {
	var bodyReader io.ReadSeeker
	if bodyAsBytes == nil {
		bodyReader = nil
	} else {
		bodyReader = bytes.NewReader(bodyAsBytes)
	}

	credentials := credentials.NewStaticCredentials(client.config.AccessKeyID, client.config.SecretAccessKey, "")
	_, err := v4.NewSigner(credentials).Sign(req, bodyReader, client.config.AWSServiceName, client.config.AWSRegion, time.Now())
	if err != nil {
		return nil, fmt.Errorf("Failed to sign request: %s", err)
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute request: %s", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respAsBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Failed to read response body: %s", err)
		}
		return nil, fmt.Errorf(
			"Expected a 2xx status code, but got %d.\nMethod: %s\nURL: %s\nRequest body: %s\nResponse body: %s",
			resp.StatusCode, req.Method, req.URL, bodyAsBytes, respAsBytes)
	}

	return resp, nil
}

func (client *Client) unmarshalXMLBody(bodyReader io.Reader, v interface{}) error {
	bodyAsBytes, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return fmt.Errorf("Failed to read body: %s", err)
	}

	if err := xml.Unmarshal(bodyAsBytes, v); err != nil {
		return fmt.Errorf("Failed to unmarshal XML: %s", err)
	}

	return nil
}
