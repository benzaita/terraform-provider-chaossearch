package chaossearch

import (
	"context"
	"terraform-provider-chaossearch/chaossearch/client"

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
	// request.Filter, err = mapFilterDataToRequest(data.Get("filter").(map[string]interface{}))

	err := c.CreateObjectGroup(ctx, request)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(data.Get("name").(string)) //TODO remove

	return resourceObjectGroupRead(ctx, data, meta)
}

// func mapFilterDataToRequest(filterData map[string]interface{}) (client.ObjectGroupFilter, error) {
// operator := filterData
// }

func resourceObjectGroupRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*ProviderMeta).Client

	req := &client.ReadObjectGroupRequest{
		Name: data.Get("name").(string),
	}

	resp, err := c.ReadObjectGroup(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("compression", resp.Compression); err != nil {
		return diag.FromErr(err)
	}

	return nil
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
