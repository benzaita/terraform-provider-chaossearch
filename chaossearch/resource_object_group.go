package chaossearch

import (
	"context"
	"log"
	"regexp"
	"strings"
	"terraform-provider-chaossearch/chaossearch/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

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
			},
			"pattern": {
				Type:     schema.TypeString,
				Default:  ".*",
				Optional: true,
				ForceNew: true,
			},
			"index_retention": {
				Type:        schema.TypeInt,
				Default:     14,
				Description: "Number of days to keep the data before deleting it",
				Optional:    true,
				ForceNew:    false,
			},
			"active": {
				Type:        schema.TypeBool,
				Description: "Whether the live indexing should be running or not",
				Required:    true,
				ForceNew:    false,
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

	createObjectGroupRequest := &client.CreateObjectGroupRequest{
		Name:             data.Get("name").(string),
		SourceBucket:     data.Get("source_bucket").(string),
		FilterJSON:       data.Get("filter_json").(string),
		Format:           data.Get("format").(string),
		Compression:      data.Get("compression").(string),
		LiveEventsSqsArn: data.Get("live_events_sqs_arn").(string),
		PartitionBy:      data.Get("partition_by").(string),
		Pattern:          data.Get("pattern").(string),
		IndexRetention:   data.Get("index_retention").(int),
	}

	if err := c.CreateObjectGroup(ctx, createObjectGroupRequest); err != nil {
		return diag.FromErr(err)
	}

	setActiveRequest := &client.SetActiveRequest{
		ObjectGroupName: data.Get("name").(string),
		Active:          data.Get("active").(bool),
	}

	if err := c.SetActive(ctx, setActiveRequest); err != nil {
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
	data.Set("index_retention", resp.IndexRetention)
	data.Set("active", resp.Active)

	// When the object in an Object Group use no compression, you need to create it with
	// `compression = ""`. However, when querying an Object Group whose object are not
	// compressed, the API returns `compression = "none"`. We coerce the "none" value to
	// an empty string in order not to confuse Terraform.
	compressionOrEmptyString := resp.Compression
	if strings.ToLower(compressionOrEmptyString) == "none" {
		compressionOrEmptyString = ""
	}
	data.Set("compression", compressionOrEmptyString)

	log.Printf("WARNING - not reading PartitionBy")
	// data.Set("partition_by", resp.PartitionBy)

	data.Set("source_bucket", resp.SourceBucket)

	return diags
}

func resourceObjectGroupUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*ProviderMeta).Client

	setActiveRequest := &client.SetActiveRequest{
		ObjectGroupName: data.Get("name").(string),
		Active:          data.Get("active").(bool),
	}
	if err := c.SetActive(ctx, setActiveRequest); err != nil {
		return diag.FromErr(err)
	}

	updateObjectGroupRequest := &client.UpdateObjectGroupRequest{
		Name:           data.Get("name").(string),
		IndexRetention: data.Get("index_retention").(int),
	}

	if err := c.UpdateObjectGroup(ctx, updateObjectGroupRequest); err != nil {
		return diag.FromErr(err)
	}

	return resourceObjectGroupRead(ctx, data, meta)
}

func resourceObjectGroupDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*ProviderMeta).Client

	stopIndexingRequest := &client.SetActiveRequest{
		ObjectGroupName: data.Get("name").(string),
		Active:          false,
	}
	if err := c.SetActive(ctx, stopIndexingRequest); err != nil {
		return diag.FromErr(err)
	}

	deleteObjectGroupRequest := &client.DeleteObjectGroupRequest{
		Name: data.Get("name").(string),
	}

	if err := c.DeleteObjectGroup(ctx, deleteObjectGroupRequest); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
