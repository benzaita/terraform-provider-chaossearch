package chaossearch

import (
	"context"
	"strconv"
	"terraform-provider-chaossearch/chaossearch/client"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceObjectGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceObjectGroupsRead,
		Schema: map[string]*schema.Schema{
			"object_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						// TODO BucketType (object group, view, native s3 bucket)
						// TODO Predicate
					},
				},
			},
		},
	}
}

func dataSourceObjectGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// client := &http.Client{Timeout: 10 * time.Second}

	config := client.NewConfiguration()
	config.URL = "https://klarna-eu-staging.chaossearch.io"

	client := client.NewClient(config)

	result, err := client.ListBuckets(context.Background())
	if err != nil {
		return diag.FromErr(err)
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees", "http://api:9090"), nil)
	// if err != nil {
	//   return diag.FromErr(err)
	// }

	// r, err := client.Do(req)
	// if err != nil {
	//   return diag.FromErr(err)
	// }
	// defer r.Body.Close()

	// objectGroups := make([]map[string]interface{}, 1)
	objectGroups := result

	// err = json.NewDecoder(r.Body).Decode(&coffees)
	// if err != nil {
	//   return diag.FromErr(err)
	// }

	if err := d.Set("object_groups", objectGroups); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
