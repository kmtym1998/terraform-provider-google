package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func DataSourceGoogleSourceRepoRepository() *schema.Resource {

	dsSchema := datasourceSchemaFromResourceSchema(ResourceSourceRepoRepository().Schema)

	addRequiredFieldsToSchema(dsSchema, "name")
	addOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceGoogleSourceRepoRepositoryRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleSourceRepoRepositoryRead(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*transport_tpg.Config)

	id, err := ReplaceVars(d, config, "projects/{{project}}/repos/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return resourceSourceRepoRepositoryRead(d, meta)
}
