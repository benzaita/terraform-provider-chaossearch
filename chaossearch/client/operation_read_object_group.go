package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type appLogger struct{}

func (l appLogger) Log(args ...interface{}) {
	log.Printf("AWS: %+v", args...)
}

func (client *Client) ReadObjectGroup(ctx context.Context, req *ReadObjectGroupRequest) (*ReadObjectGroupResponse, error) {
	session, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(client.config.AccessKeyID, client.config.SecretAccessKey, ""),
		Endpoint:         aws.String(fmt.Sprintf("%s/V1", client.config.URL)),
		Region:           aws.String("eu-west-1"),
		S3ForcePathStyle: aws.Bool(true),
		LogLevel:         aws.LogLevel(aws.LogOff),
		Logger:           appLogger{},
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to create AWS session: %s", err)
	}

	svc := s3.New(session)
	input := &s3.GetBucketTaggingInput{
		Bucket: aws.String(req.ID),
	}

	tagging, err := svc.GetBucketTaggingWithContext(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("Failed to read bucket tagging: %s", err)
	}

	var resp ReadObjectGroupResponse
	if err := mapBucketTaggingToResponse(tagging, &resp); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal XML response body: %s", err)
	}

	log.Printf("ReadObjectGroupResponse: %+v", resp)

	return &resp, nil
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
		Type string `json:"_type"`
	}
	if err := readJSONTagValue(tagging, "cs3.dataset-format", &filterObject); err != nil {
		return err
	}
	v.Format = filterObject.Type

	if err := readStringTagValue(tagging, "cs3.predicate", &v.FilterJSON); err != nil {
		return err
	}

	log.Printf("WARNING - not reading PartitionBy")
	// v.PartitionBy = ""                // TODO where from?
	// log.Fatalf("Not implemented yet")
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
