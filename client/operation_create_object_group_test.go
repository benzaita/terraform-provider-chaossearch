package client

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"testing"
)

// Test CreateObjectGroup request json - no filters
func TestMarshalCreateObjectGroupRequestNoFilters(t *testing.T) {
	array_flatten_depth := 0
	req := CreateObjectGroupRequest{
		Name:              "test",
		Compression:       "gzip",
		FilterJSON:        `{"AND":[{"field":"key","regex":".*"}]}`,
		Format:            "JSON",
		LiveEventsSqsArn:  "test-sqs-arn",
		PartitionBy:       "",
		SourceBucket:      "test-bucket",
		Pattern:           ".*",
		IndexRetention:    123,
		KeepOriginal:      true,
		ArrayFlattenDepth: &array_flatten_depth,
		ColumnRenames:     map[string]interface{}{},
		ColumnSelection:   map[string]interface{}{},
	}
	expected_res := map[string]interface{}{
		"bucket": "test",
		"filter": map[string]interface{}{"AND": []interface{}{map[string]interface{}{"field": "key", "regex": ".*"}}},
		"format": map[string]interface{}{
			"_type":             "JSON",
			"arrayFlattenDepth": float64(0),
			"horizontal":        true,
			"keepOriginal":      true,
			"stripPrefix":       true,
		},
		"indexRetention": float64(123),
		"interval":       map[string]interface{}{"column": float64(0), "mode": float64(0)},
		"liveEvents":     "test-sqs-arn",
		"options": map[string]interface{}{
			"compression":     "gzip",
			"ignoreIrregular": true,
		},
		"source": "test-bucket",
	}
	bytes, err := marshalCreateObjectGroupRequest(&req)
	if err != nil {
		t.Fatalf("Marshalling error: %s", err)
	}

	res := map[string]interface{}{}
	if err := json.Unmarshal(bytes, &res); err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err)
	}

	if !cmp.Equal(res, expected_res) {
		t.Fatal(cmp.Diff(res, expected_res))
	}
}

// Test CreateObjectGroup request json - column selection by whitelist
func TestMarshalCreateObjectGroupRequestColSelectionWhitelist(t *testing.T) {
	array_flatten_depth := 0
	req := CreateObjectGroupRequest{
		Name:              "test",
		Compression:       "gzip",
		FilterJSON:        `{"AND":[{"field":"key","regex":".*"}]}`,
		Format:            "JSON",
		LiveEventsSqsArn:  "test-sqs-arn",
		PartitionBy:       "",
		SourceBucket:      "test-bucket",
		Pattern:           ".*",
		IndexRetention:    123,
		KeepOriginal:      true,
		ArrayFlattenDepth: &array_flatten_depth,
		ColumnRenames:     map[string]interface{}{},
		ColumnSelection: map[string]interface{}{
			"type": "whitelist",
			"includes": []interface{}{
				"column1", "column2",
			},
		},
	}
	expected_res := map[string]interface{}{
		"bucket": "test",
		"filter": map[string]interface{}{"AND": []interface{}{map[string]interface{}{"field": "key", "regex": ".*"}}},
		"format": map[string]interface{}{
			"_type":             "JSON",
			"arrayFlattenDepth": float64(0),
			"horizontal":        true,
			"keepOriginal":      true,
			"stripPrefix":       true,
		},
		"indexRetention": float64(123),
		"interval":       map[string]interface{}{"column": float64(0), "mode": float64(0)},
		"liveEvents":     "test-sqs-arn",
		"options": map[string]interface{}{
			"colSelection":    []interface{}{map[string]interface{}{"includes": []interface{}{"column1", "column2"}, "type": string("whitelist")}},
			"compression":     "gzip",
			"ignoreIrregular": true,
		},
		"source": string("test-bucket"),
	}
	bytes, err := marshalCreateObjectGroupRequest(&req)
	if err != nil {
		t.Fatalf("Marshalling error: %s", err)
	}

	res := map[string]interface{}{}
	if err := json.Unmarshal(bytes, &res); err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err)
	}

	if !cmp.Equal(res, expected_res) {
		t.Fatal(cmp.Diff(res, expected_res))
	}
}

// Test CreateObjectGroup request json - column selection by regex
func TestMarshalCreateObjectGroupRequestColSelectionRegex(t *testing.T) {
	array_flatten_depth := 0
	req := CreateObjectGroupRequest{
		Name:              "test",
		Compression:       "gzip",
		FilterJSON:        `{"AND":[{"field":"key","regex":".*"}]}`,
		Format:            "JSON",
		LiveEventsSqsArn:  "test-sqs-arn",
		PartitionBy:       "",
		SourceBucket:      "test-bucket",
		Pattern:           ".*",
		IndexRetention:    123,
		KeepOriginal:      true,
		ArrayFlattenDepth: &array_flatten_depth,
		ColumnRenames:     map[string]interface{}{},
		ColumnSelection: map[string]interface{}{
			"type": "regex",
			"patterns": []interface{}{
				"^Timestamp$",
				"^line\\.meta\\..+$",
			},
			"include": true,
		},
	}
	expected_res := map[string]interface{}{
		"bucket": "test",
		"filter": map[string]interface{}{"AND": []interface{}{map[string]interface{}{"field": "key", "regex": ".*"}}},
		"format": map[string]interface{}{
			"_type":             "JSON",
			"arrayFlattenDepth": float64(0),
			"horizontal":        true,
			"keepOriginal":      true,
			"stripPrefix":       true,
		},
		"indexRetention": float64(123),
		"interval":       map[string]interface{}{"column": float64(0), "mode": float64(0)},
		"liveEvents":     "test-sqs-arn",
		"options": map[string]interface{}{
			"colSelection": []interface{}{map[string]interface{}{
				"type": "regex",
				"patterns": []interface{}{
					"^Timestamp$",
					"^line\\.meta\\..+$",
				},
				"include": true,
			}},
			"compression":     "gzip",
			"ignoreIrregular": true,
		},
		"source": string("test-bucket"),
	}
	bytes, err := marshalCreateObjectGroupRequest(&req)
	if err != nil {
		t.Fatalf("Marshalling error: %s", err)
	}

	res := map[string]interface{}{}
	if err := json.Unmarshal(bytes, &res); err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err)
	}

	if !cmp.Equal(res, expected_res) {
		t.Fatal(cmp.Diff(res, expected_res))
	}
}
