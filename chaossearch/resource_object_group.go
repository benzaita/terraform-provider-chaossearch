package chaossearch

import (
	"context"
	"log"
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
			},
			"source_bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"logs_type": {
				Type:     schema.TypeString,
				Default:  "JSON",
				Optional: true,
				// TODO add validation
			},
			"horizontal": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"strip_prefix": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"retention_days": {
				Type:     schema.TypeInt,
				Default:  -1,
				Optional: true,
			},
			"ignore_irregular": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"compression": {
				Type:     schema.TypeString,
				Required: true,
				// TODO add validation
			},
			"daily_interval": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"live_events_sqs_arn": {
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
			},
			"partition_by": {
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
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
		RetentionDays: data.Get("retention_days").(int),
		Options: client.ObjectGroupOptions{
			IgnoreIrregular: data.Get("ignore_irregular").(bool),
			Compression:     data.Get("compression").(string),
		},
		DailyInterval:    data.Get("daily_interval").(bool),
		LiveEventsSqsArn: data.Get("live_events_sqs_arn").(string),
		PartitionBy:      data.Get("partition_by").(string),
	}
	// request.Filter, err = mapFilterDataToRequest(data.Get("filter").(map[string]interface{}))

	log.Printf("CreateObjectGroupRequest: %+v\n", request)

	err := c.CreateObjectGroup(ctx, request)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(data.Get("name").(string))

	return resourceObjectGroupRead(ctx, data, meta)
}

// func mapFilterDataToRequest(filterData map[string]interface{}) (client.ObjectGroupFilter, error) {
// operator := filterData
// }

func resourceObjectGroupRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ProviderMeta).Client

	log.Printf("ListBucketsRequest: {}\n")

	clientResponse, err := client.ListBuckets(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("ListBucketsResponse: %+v\n", clientResponse)
	found := false
	for _, bucket := range clientResponse.BucketsCollection.Buckets {
		if bucket.Name == data.Id() {
			if err := data.Set("name", bucket.Name); err != nil {
				return diag.FromErr(err)
			}
			found = true
		}
	}

	if !found {
		return diag.Errorf("Cannot find Bucket with ID '%v'", data.Id())
	}

	return nil
}

func resourceObjectGroupUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("Not implemented")
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
