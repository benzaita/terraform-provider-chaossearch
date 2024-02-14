package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type appLogger struct{}

func (l appLogger) Log(args ...interface{}) {
	log.Printf("AWS: %+v", args...)
}

// BUG: ChaosSearch breaking API change now uses two representations of "regex", one as an object, the other as a string.
// The function `processFilterJSON` takes a JSON string as input, checks the structure of the "regex" key,
// and returns a modified JSON string according to the specified rules.
// For example:
//     `{"AND":[{"field":"key","regex":{"pattern":".*","strict":true}}]}`
// is converted back into
//     `{"AND":[{"field":"key","regex":".*"}]}`
func processFilterJSON(input string) (string, error) {
    // Unmarshal the input JSON string into a map structure.
    var data map[string]interface{}
    if err := json.Unmarshal([]byte(input), &data); err != nil {
        return "", err // Return an error if JSON is invalid.
    }

    // Iterate over the "AND" array in the JSON structure.
    for _, item := range data["AND"].([]interface{}) {
        entry := item.(map[string]interface{}) // Cast to a map.
        regex := entry["regex"]                // Access the "regex" key.

        // Check the type of the "regex" key.
        switch regex.(type) {
        case map[string]interface{}:
            // If "regex" is a map, set it to the value of the "pattern" key.
            regexMap := regex.(map[string]interface{})
            entry["regex"] = regexMap["pattern"]
        case string:
            // If "regex" is a string, no action is needed.
        default:
            // Return an error if "regex" is neither a string nor a map.
            return "", fmt.Errorf("unexpected type for regex key")
        }
    }

    // Marshal the modified map structure back into a JSON string.
    outputBytes, err := json.Marshal(data)
    if err != nil {
        return "", err // Return an error if marshaling fails.
    }

    // Return the modified JSON string.
    return string(outputBytes), nil
}

func (client *Client) ReadObjectGroup(ctx context.Context, req *ReadObjectGroupRequest) (*ReadObjectGroupResponse, error) {
	var resp ReadObjectGroupResponse

	if err := client.readAttributesFromBucketTagging(ctx, req, &resp); err != nil {
		return nil, err
	}

	if err := client.readAttributesFromDatasetEndpoint(ctx, req, &resp); err != nil {
		return nil, err
	}

	log.Printf("ReadObjectGroupResponse: %+v", resp)

	return &resp, nil
}

func (client *Client) readAttributesFromDatasetEndpoint(ctx context.Context, req *ReadObjectGroupRequest, resp *ReadObjectGroupResponse) error {
	method := "GET"
	url := fmt.Sprintf("%s/Bucket/dataset/name/%s", client.config.URL, req.ID)

	httpReq, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return fmt.Errorf("Failed to create request: %s", err)
	}

	httpResp, err := client.signAndDo(httpReq, nil)
	if err != nil {
		return fmt.Errorf("Failed to %s to %s: %s", method, url, err)
	}
	defer httpResp.Body.Close()

	var getDatasetResp struct {
		PartitionBy string `json:"partitionBy"`
		Options     struct {
			ColumnRenames   map[string]string        `json:"colRenames"`
			ColumnSelection []map[string]interface{} `json:"colSelection"`
			ColumnTypes     map[string]string        `json:"colTypes"`
		} `json:"options"`
	}
	if err := client.unmarshalJSONBody(httpResp.Body, &getDatasetResp); err != nil {
		return fmt.Errorf("Failed to unmarshal JSON response body: %s", err)
	}
	if len(getDatasetResp.Options.ColumnSelection) > 0 {
		if _, ok := getDatasetResp.Options.ColumnSelection[0]["include"]; !ok {
			// API doesn't return "include" field for type=whitelist|blacklist, but we set it in the state.
			// That results in a diff in configuration that causes an unchanged object group to be recreated.
			getDatasetResp.Options.ColumnSelection[0]["include"] = false
		}

	}

	resp.PartitionBy = getDatasetResp.PartitionBy
	resp.ColumnRenames = getDatasetResp.Options.ColumnRenames
	resp.ColumnSelection = getDatasetResp.Options.ColumnSelection
	resp.ColumnTypes = getDatasetResp.Options.ColumnTypes

	return nil
}

func (client *Client) readAttributesFromBucketTagging(ctx context.Context, req *ReadObjectGroupRequest, resp *ReadObjectGroupResponse) error {
	session, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(client.config.AccessKeyID, client.config.SecretAccessKey, ""),
		Endpoint:         aws.String(fmt.Sprintf("%s/V1", client.config.URL)),
		Region:           aws.String(client.config.Region),
		S3ForcePathStyle: aws.Bool(true),
		LogLevel:         aws.LogLevel(aws.LogOff),
		Logger:           appLogger{},
	})
	if err != nil {
		return fmt.Errorf("Failed to create AWS session: %s", err)
	}

	svc := s3.New(session)
	input := &s3.GetBucketTaggingInput{
		Bucket: aws.String(req.ID),
	}

	tagging, err := svc.GetBucketTaggingWithContext(ctx, input)
	if err != nil {
		return fmt.Errorf("Failed to read bucket tagging: %s", err)
	}

	if err := mapBucketTaggingToResponse(tagging, resp); err != nil {
		return fmt.Errorf("Failed to unmarshal XML response body: %s", err)
	}

	return nil
}

func mapBucketTaggingToResponse(tagging *s3.GetBucketTaggingOutput, v *ReadObjectGroupResponse) error {
	if err := readStringTagValue(tagging, "cs3.parent", &v.SourceBucket); err != nil {
		return err
	}

	if err := readStringTagValue(tagging, "cs3.compression", &v.Compression); err != nil {
		return err
	}

	if err := readStringTagValue(tagging, "cs3.live-sqs-arn", &v.LiveEventsSqsArn); err != nil {
		return err
	}

	var filterObject struct {
		Type              string `json:"_type"`
		Pattern           string `json:"pattern"`
		ArrayFlattenDepth *int   `json:"arrayFlattenDepth"`
		KeepOriginal      bool   `json:"keepOriginal"`
		Horizontal        bool   `json:"horizontal"`
	}
	if err := readJSONTagValue(tagging, "cs3.dataset-format", &filterObject); err != nil {
		return err
	}
	v.Format = filterObject.Type
	v.Pattern = filterObject.Pattern
	v.ArrayFlattenDepth = filterObject.ArrayFlattenDepth
	v.KeepOriginal = filterObject.KeepOriginal
	v.Horizontal = filterObject.Horizontal

	if err := readStringTagValue(tagging, "cs3.predicate", &v.FilterJSON); err != nil {
		return err
	}


	// NOTE: this is to work around a non-versioned breaking API change in ChaosSearch.
	// The API might return a slightly modified JSON string.
	inputFilterJSON := v.FilterJSON
	outputJSON, err := processFilterJSON(inputFilterJSON)
    if err != nil {
        return err
    } else {
        v.FilterJSON = outputJSON
    }

	var retentionObject struct {
		Overall int `json:"overall"`
	}
	if err := readJSONTagValue(tagging, "cs3.index-retention", &retentionObject); err != nil {
		return err
	}
	v.IndexRetention = retentionObject.Overall

	return nil
}

func readStringTagValue(tagging *s3.GetBucketTaggingOutput, key string, v *string) error {
	stringValue, err := findTagValue(tagging, key)
	if err != nil {
		return nil
	}

	*v = stringValue
	return nil
}

func readJSONTagValue(tagging *s3.GetBucketTaggingOutput, key string, v interface{}) error {
	valueAsBytes, err := findTagValue(tagging, key)
	if err != nil {
		return nil
	}

	return json.Unmarshal([]byte(valueAsBytes), v)
}

func findTagValue(tagging *s3.GetBucketTaggingOutput, key string) (string, error) {
	for _, tag := range tagging.TagSet {
		if *tag.Key == key {
			return *tag.Value, nil
		}
	}

	return "", fmt.Errorf("No tag found with key: %s", key)
}
