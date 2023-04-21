package google

import (
	"fmt"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGoogleRuntimeconfigConfig() *schema.Resource {

	dsSchema := datasourceSchemaFromResourceSchema(ResourceRuntimeconfigConfig().Schema)
	addRequiredFieldsToSchema(dsSchema, "name")
	addOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceGoogleRuntimeconfigConfigRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleRuntimeconfigConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)

	id, err := ReplaceVars(d, config, "projects/{{project}}/configs/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	return resourceRuntimeconfigConfigRead(d, meta)
}
