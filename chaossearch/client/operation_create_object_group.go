package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) CreateObjectGroup(ctx context.Context, req *CreateObjectGroupRequest) error {
	method := "POST"
	url := fmt.Sprintf("%s/Bucket/createObjectGroup", client.config.URL)

	bodyAsBytes, err := marshalCreateObjectGroupRequest(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(bodyAsBytes))
	if err != nil {
		return fmt.Errorf("Failed to create request: %s", err)
	}

	httpResp, err := client.signAndDo(httpReq, bodyAsBytes)
	if err != nil {
		return fmt.Errorf("Failed to %s to %s: %s", method, url, err)
	}
	defer httpResp.Body.Close()

	return nil
}

func marshalCreateObjectGroupRequest(req *CreateObjectGroupRequest) ([]byte, error) {
	body := map[string]interface{}{
		"bucket": req.Name,
		"source": req.SourceBucket,
		"format": map[string]interface{}{
			"_type":       req.Format.Type,
			"horizontal":  req.Format.Horizontal,
			"stripPrefix": req.Format.StripPrefix,
		},
		"indexRetention": req.RetentionDays,
		"options": map[string]interface{}{
			"ignoreIrregular": req.Options.IgnoreIrregular,
			"compression":     req.Options.Compression,
		},
	}

	// filterBody, err := mapObjectGroupFilterRequestToBody(&req.Filter)
	// if err != nil {
	// 	return nil, err
	// }
	body["filter"] = []map[string]interface{}{
		{
			"field": "key",
			"regex": ".*",
		},
	}

	if req.DailyInterval {
		body["interval"] = map[string]interface{}{
			"mode":   0,
			"column": 0,
		}
	}

	if req.LiveEventsSqsArn != "" {
		body["liveEvents"] = req.LiveEventsSqsArn
	}

	if req.PartitionBy != "" {
		body["partitionBy"] = req.PartitionBy
	}

	bodyAsBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bodyAsBytes, nil
}

func mapObjectGroupFilterRequest(filter *ObjectGroupFilter) (map[string]interface{}, error) {
	operandsListBody := make([]map[string]interface{}, len(filter.Operands))
	for i, operand := range filter.Operands {
		operandBody := make(map[string]interface{})
		operandBody["field"] = operand.FieldName
		if operand.ValuePrefix != "" && operand.ValueRegex != "" {
			return nil, fmt.Errorf("Expected only one of ValuePrefix or ValueRegex to be defined, but both were defined in %+v", operand)
		}
		if operand.ValuePrefix != "" {
			operandBody["prefix"] = operand.ValuePrefix
		}
		if operand.ValueRegex != "" {
			operandBody["regex"] = operand.ValueRegex
		}

		operandsListBody[i] = operandBody
	}

	filterBody := make(map[string]interface{})
	filterBody[filter.Operator] = operandsListBody
	return filterBody, nil
}
