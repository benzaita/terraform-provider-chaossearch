package chaossearch

import (
	"context"
	"terraform-provider-chaossearch/chaossearch/client"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceObjectGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceObjectGroupCreate,
		ReadContext:   resourceObjectGroupRead,
		UpdateContext: resourceObjectGroupUpdate,
		DeleteContext: resourceObjectGroupDelete,
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
			"logs_type": {
				Type:     schema.TypeString,
				Default:  "JSON",
				Optional: true,
				ForceNew: true,
				// TODO add validation
			},
			"horizontal": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
				ForceNew: true,
			},
			"strip_prefix": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
				ForceNew: true,
			},
			"index_retention": {
				Type:     schema.TypeInt,
				Default:  -1,
				Optional: true,
			},
			"ignore_irregular": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
				ForceNew: true,
			},
			"compression": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				// TODO add validation
			},
			"daily_interval": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"live_events_sqs_arn": {
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
				ForceNew: true,
			},
			"partition_by": {
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceObjectGroupCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*ProviderMeta).Client

	request := &client.CreateObjectGroupRequest{
		Name:         data.Get("name").(string),
		SourceBucket: data.Get("source_bucket").(string),
		Format: client.ObjectGroupFormat{
			Type:        data.Get("logs_type").(string),
			Horizontal:  data.Get("horizontal").(bool),
			StripPrefix: data.Get("strip_prefix").(bool),
		},
		IndexRetention: data.Get("index_retention").(int),
		Options: client.ObjectGroupOptions{
			IgnoreIrregular: data.Get("ignore_irregular").(bool),
			Compression:     data.Get("compression").(string),
		},
		DailyInterval:    data.Get("daily_interval").(bool),
		LiveEventsSqsArn: data.Get("live_events_sqs_arn").(string),
		PartitionBy:      data.Get("partition_by").(string),
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
		Id: data.Id(),
	}

	resp, err := c.ReadObjectGroup(ctx, req)
	if err != nil {
		return diag.Errorf("Failed to read object group: ", err)
	}

	log.Printf("WARNING - not reading Horizontal, DailyInterval, IndexRetention, IgnoreIrregular, PartitionBy")
	// data.Set("daily_interval", resp.DailyInterval)
	// data.Set("horizontal", resp.Format.Horizontal)
	data.Set("strip_prefix", resp.Format.StripPrefix)
	data.Set("logs_type", resp.Format.Type)
	// data.Set("index_retention", resp.IndexRetention)
	data.Set("live_events_sqs_arn", resp.LiveEventsSqsArn)
	data.Set("compression", resp.Options.Compression)
	// data.Set("ignore_irregular", resp.Options.IgnoreIrregular)
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
