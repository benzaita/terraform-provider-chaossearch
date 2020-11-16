package client

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func (client *Client) ReadObjectGroup(ctx context.Context, req *ReadObjectGroupRequest) (*ReadObjectGroupResponse, error) {
	method := "GET"
	safeObjectGroupName := url.PathEscape(req.Name)
	url := fmt.Sprintf("%s/V1/%s?tagging", client.config.URL, safeObjectGroupName)

	httpReq, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %s", err)
	}

	httpResp, err := client.signAndDo(httpReq, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to %s to %s: %s", method, url, err)
	}
	defer httpResp.Body.Close()

	var resp ReadObjectGroupResponse
	if err := unmarshalReadObjectGroupResponse(httpResp.Body, &resp); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal XML response body: %s", err)
	}

	log.Printf("ReadObjectGroupResponse: %+v", resp)

	return &resp, nil
}

type objectGroupTag struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

type objectGroupTagging struct {
	TagSet []objectGroupTag `xml:"TagSet>Tag"`
}

func unmarshalReadObjectGroupResponse(bodyReader io.Reader, v *ReadObjectGroupResponse) error {
	bodyAsBytes, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return fmt.Errorf("Failed to read body: %s", err)
	}

	log.Printf("bodyAsBytes: %s", bodyAsBytes)

	var tagging objectGroupTagging
	if err := xml.Unmarshal(bodyAsBytes, &tagging); err != nil {
		return fmt.Errorf("Failed to unmarshal XML: %s", err)
	}

	log.Printf("objectGroupTagging: %+v", tagging)

	if err := readStringTagValue(&tagging, "cs3.parent", &v.SourceBucket); err != nil {
		return err
	}

	if err := readStringTagValue(&tagging, "cs3.compression", &v.Options.Compression); err != nil {
		return err
	}

	if err := readStringTagValue(&tagging, "cs3.live-sqs-arn", &v.LiveEventsSqsArn); err != nil {
		return err
	}

	if err := readJSONTagValue(&tagging, "cs3.dataset-format", &v.Format); err != nil {
		return err
	}

	v.Format.Horizontal = false       // TODO where from?
	v.DailyInterval = false           // TODO where from?
	v.IndexRetention = -1             // TODO where from?
	v.Options.IgnoreIrregular = false // TODO where from?
	v.PartitionBy = ""                // TODO where from?
	return fmt.Errorf("Not implemented yet")
}

func readStringTagValue(tagging *objectGroupTagging, key string, v *string) error {
	stringValue, err := findTagValue(tagging, key)
	if err != nil {
		return nil
	}

	*v = stringValue
	return nil
}

func readJSONTagValue(tagging *objectGroupTagging, key string, v interface{}) error {
	valueAsBytes, err := findTagValue(tagging, key)
	if err != nil {
		return nil
	}

	return json.Unmarshal([]byte(valueAsBytes), v)
}

func findTagValue(tagging *objectGroupTagging, key string) (string, error) {
	for _, tag := range tagging.TagSet {
		if tag.Key == key {
			return tag.Value, nil
		}
	}

	return "", fmt.Errorf("No tag found with key: %s", key)
}
