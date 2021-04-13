package chaossearch

import (
	"context"
	"terraform-provider-chaossearch/chaossearch/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIndexingState() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIndexingStateCreate,
		ReadContext:   resourceIndexingStateRead,
		UpdateContext: resourceIndexingStateUpdate,
		DeleteContext: resourceIndexingStateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"objectGroupName": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"active": {
				Type:        schema.TypeBool,
				Description: "Whether the live indexing should be running or not",
				Required:    true,
				ForceNew:    false,
			},
		},
	}
}

func resourceIndexingStateCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*ProviderMeta).Client

	setActiveRequest := &client.SetActiveRequest{
		ObjectGroupName: data.Get("objectGroupName").(string),
		Active:          data.Get("active").(bool),
	}

	if err := c.UpdateIndexingState(ctx, setActiveRequest); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(data.Get("objectGroupName").(string))

	return resourceIndexingStateRead(ctx, data, meta)
}

func resourceIndexingStateRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diags := diag.Diagnostics{}

	// There is no request that is similar to getting an indexing state in ChaosSearch API,
	//  so we cannot retrieve it, therefore this function is kept for the sake of being
	//  compliant with Terraform providers standard, but it doesn't do anything

	return diags
}

func resourceIndexingStateUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*ProviderMeta).Client

	setActiveRequest := &client.SetActiveRequest{
		ObjectGroupName: data.Get("objectGroupName").(string),
		Active:          data.Get("active").(bool),
	}
	if err := c.UpdateIndexingState(ctx, setActiveRequest); err != nil {
		return diag.FromErr(err)
	}

	return resourceIndexingStateRead(ctx, data, meta)
}

func resourceIndexingStateDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*ProviderMeta).Client

	stopIndexingRequest := &client.SetActiveRequest{
		ObjectGroupName: data.Get("objectGroupName").(string),
		Active:          false,
	}
	if err := c.UpdateIndexingState(ctx, stopIndexingRequest); err != nil {
		return diag.FromErr(err)
	}

	data.SetId("")

	return nil
}
