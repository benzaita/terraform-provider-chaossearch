package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
)


type readBucketMetadataResponse struct {
	Metadata map[string]bucketMetadata `json:"Metadata"`
}

// There are other nested structures that we don't marshal here because we don't need to.
//  Feel free to add them in the future when needed
type bucketMetadata struct {
	Bucket string `json:"Bucket"`
	State  string `json:"State"`
}

// For documentation see: https://docs.chaossearch.io/reference#bucketmodel
func (client *Client) ReadIndexingState(ctx context.Context, req *ReadIndexingStateRequest) (*IndexingState, error) {
	method := "GET"

	responseBody, err := makeGetBucketMetadataRequest(method, client, ctx)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}

	var response readBucketMetadataResponse
	err = client.unmarshalJSONBody(responseBody, response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %s", err)
	}

	if _, ok := response.Metadata[req.ObjectGroupName]; !ok {
		return nil, fmt.Errorf("response doesn't contain the needed Bucket metadata: %s", req.ObjectGroupName)
	}

	bucketMetadata := &IndexingState{
		ObjectGroupName: req.ObjectGroupName,
		Active: false,
	}

	if response.Metadata[req.ObjectGroupName].State == "Active" {
		bucketMetadata.Active = true
	}

	return bucketMetadata, nil
}

func makeGetBucketMetadataRequest(method string, client *Client, ctx context.Context) (io.ReadCloser, error) {
	url := fmt.Sprintf("%s/Bucket/metadata", client.config.URL)

	httpReq, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}

	httpResp, err := client.signAndDo(httpReq, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to %s to %s: %s", method, url, err)
	}
	defer httpResp.Body.Close()

	return httpResp.Body, nil
}