package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func DataSourceDataprocMetastoreService() *schema.Resource {

	dsSchema := datasourceSchemaFromResourceSchema(ResourceDataprocMetastoreService().Schema)
	addRequiredFieldsToSchema(dsSchema, "service_id")
	addRequiredFieldsToSchema(dsSchema, "location")
	addOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceDataprocMetastoreServiceRead,
		Schema: dsSchema,
	}
}

func dataSourceDataprocMetastoreServiceRead(d *schema.ResourceData, meta interface{}) error {
	id, err := ReplaceVars(d, meta.(*transport_tpg.Config), "projects/{{project}}/locations/{{location}}/services/{{service_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	return resourceDataprocMetastoreServiceRead(d, meta)
}
