package chaossearch

import (
	"context"
	"log"
	"fmt"
	"regexp"
	"terraform-provider-chaossearch/chaossearch/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/gijsbers/go-pcre"
)

func stringIsValidPcre(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
		return warnings, errors
	}

	if _, err := pcre.Compile(v, 0); err != nil {
		errors = append(errors, fmt.Errorf("%q: %s", k, err))
	}

	return warnings, errors
}

func resourceObjectGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceObjectGroupCreate,
		ReadContext:   resourceObjectGroupRead,
		UpdateContext: resourceObjectGroupUpdate,
		DeleteContext: resourceObjectGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"format": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"JSON", "LOG"}, false),
			},
			"filter_json": {
				Type:         schema.TypeString,
				Default:      `[{"field":"key","regex":".*"}]`,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsJSON,
			},
			"compression": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "",
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"", "gzip", "zlib", "snappy"}, true),
			},
			"live_events_sqs_arn": {
				Type:         schema.TypeString,
				Default:      "",
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^arn:aws:sqs:.*`), `must be an SQS ARN: "arn:aws:sqs:..."`),
			},
			"partition_by": {
				Type:         schema.TypeString,
				Default:      "",
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsValidRegExp,
			},
			"pattern": {
				Type:         schema.TypeString,
				Default:      ".*",
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: stringIsValidPcre,
			},

			// Workaround. Otherwise Terraform fails with "All fields are ForceNew or Computed w/out Optional, Update is superfluous"
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceObjectGroupCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*ProviderMeta).Client

	request := &client.CreateObjectGroupRequest{
		Name:             data.Get("name").(string),
		SourceBucket:     data.Get("source_bucket").(string),
		FilterJSON:       data.Get("filter_json").(string),
		Format:           data.Get("format").(string),
		Compression:      data.Get("compression").(string),
		LiveEventsSqsArn: data.Get("live_events_sqs_arn").(string),
		PartitionBy:      data.Get("partition_by").(string),
		Pattern:          data.Get("pattern").(string),
	}

	err := c.CreateObjectGroup(ctx, request)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(data.Get("name").(string))

	return resourceObjectGroupRead(ctx, data, meta)
}

func resourceObjectGroupRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diags := diag.Diagnostics{}
	c := meta.(*ProviderMeta).Client

	req := &client.ReadObjectGroupRequest{
		ID: data.Id(),
	}

	resp, err := c.ReadObjectGroup(ctx, req)
	if err != nil {
		return diag.Errorf("Failed to read object group: %s", err)
	}

	data.Set("name", data.Id())
	data.Set("filter_json", resp.FilterJSON)
	data.Set("format", resp.Format)
	data.Set("live_events_sqs_arn", resp.LiveEventsSqsArn)
	data.Set("compression", resp.Compression)

	log.Printf("WARNING - not reading PartitionBy")
	// data.Set("partition_by", resp.PartitionBy)

	data.Set("source_bucket", resp.SourceBucket)

	return diags
}

func resourceObjectGroupUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*ProviderMeta).Client

	req := &client.UpdateObjectGroupRequest{
		Name:           data.Get("name").(string),
		IndexRetention: data.Get("index_retention").(int),
	}

	err := c.UpdateObjectGroup(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceObjectGroupRead(ctx, data, meta)
}

func resourceObjectGroupDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*ProviderMeta).Client

	req := &client.DeleteObjectGroupRequest{
		Name: data.Get("name").(string),
	}

	err := c.DeleteObjectGroup(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
