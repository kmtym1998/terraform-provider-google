// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

var DataCatalogEntryGroupIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"region": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"entry_group": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type DataCatalogEntryGroupIamUpdater struct {
	project    string
	region     string
	entryGroup string
	d          TerraformResourceData
	Config     *transport_tpg.Config
}

func DataCatalogEntryGroupIamUpdaterProducer(d TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	region, _ := getRegion(d, config)
	if region != "" {
		if err := d.Set("region", region); err != nil {
			return nil, fmt.Errorf("Error setting region: %s", err)
		}
	}
	values["region"] = region
	if v, ok := d.GetOk("entry_group"); ok {
		values["entry_group"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/entryGroups/(?P<entry_group>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<entry_group>[^/]+)", "(?P<region>[^/]+)/(?P<entry_group>[^/]+)", "(?P<entry_group>[^/]+)"}, d, config, d.Get("entry_group").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataCatalogEntryGroupIamUpdater{
		project:    values["project"],
		region:     values["region"],
		entryGroup: values["entry_group"],
		d:          d,
		Config:     config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("region", u.region); err != nil {
		return nil, fmt.Errorf("Error setting region: %s", err)
	}
	if err := d.Set("entry_group", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting entry_group: %s", err)
	}

	return u, nil
}

func DataCatalogEntryGroupIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	region, _ := getRegion(d, config)
	if region != "" {
		values["region"] = region
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/entryGroups/(?P<entry_group>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<entry_group>[^/]+)", "(?P<region>[^/]+)/(?P<entry_group>[^/]+)", "(?P<entry_group>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataCatalogEntryGroupIamUpdater{
		project:    values["project"],
		region:     values["region"],
		entryGroup: values["entry_group"],
		d:          d,
		Config:     config,
	}
	if err := d.Set("entry_group", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting entry_group: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *DataCatalogEntryGroupIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyEntryGroupUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := SendRequest(u.Config, "POST", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *DataCatalogEntryGroupIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyEntryGroupUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = SendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *DataCatalogEntryGroupIamUpdater) qualifyEntryGroupUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{DataCatalogBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", u.project, u.region, u.entryGroup), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *DataCatalogEntryGroupIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", u.project, u.region, u.entryGroup)
}

func (u *DataCatalogEntryGroupIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-datacatalog-entrygroup-%s", u.GetResourceId())
}

func (u *DataCatalogEntryGroupIamUpdater) DescribeResource() string {
	return fmt.Sprintf("datacatalog entrygroup %q", u.GetResourceId())
}
