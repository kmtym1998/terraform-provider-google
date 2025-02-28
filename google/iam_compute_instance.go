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

var ComputeInstanceIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"zone": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"instance_name": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type ComputeInstanceIamUpdater struct {
	project      string
	zone         string
	instanceName string
	d            TerraformResourceData
	Config       *transport_tpg.Config
}

func ComputeInstanceIamUpdaterProducer(d TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	zone, _ := getZone(d, config)
	if zone != "" {
		if err := d.Set("zone", zone); err != nil {
			return nil, fmt.Errorf("Error setting zone: %s", err)
		}
	}
	values["zone"] = zone
	if v, ok := d.GetOk("instance_name"); ok {
		values["instance_name"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/instances/(?P<instance_name>[^/]+)", "(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<instance_name>[^/]+)", "(?P<zone>[^/]+)/(?P<instance_name>[^/]+)", "(?P<instance_name>[^/]+)"}, d, config, d.Get("instance_name").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ComputeInstanceIamUpdater{
		project:      values["project"],
		zone:         values["zone"],
		instanceName: values["instance_name"],
		d:            d,
		Config:       config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("zone", u.zone); err != nil {
		return nil, fmt.Errorf("Error setting zone: %s", err)
	}
	if err := d.Set("instance_name", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting instance_name: %s", err)
	}

	return u, nil
}

func ComputeInstanceIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	zone, _ := getZone(d, config)
	if zone != "" {
		values["zone"] = zone
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/instances/(?P<instance_name>[^/]+)", "(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<instance_name>[^/]+)", "(?P<zone>[^/]+)/(?P<instance_name>[^/]+)", "(?P<instance_name>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ComputeInstanceIamUpdater{
		project:      values["project"],
		zone:         values["zone"],
		instanceName: values["instance_name"],
		d:            d,
		Config:       config,
	}
	if err := d.Set("instance_name", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting instance_name: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *ComputeInstanceIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyInstanceUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}
	url, err = AddQueryParams(url, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", IamPolicyVersion)})
	if err != nil {
		return nil, err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := SendRequest(u.Config, "GET", project, url, userAgent, obj)
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

func (u *ComputeInstanceIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyInstanceUrl("setIamPolicy")
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

func (u *ComputeInstanceIamUpdater) qualifyInstanceUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{ComputeBasePath}}%s/%s", fmt.Sprintf("projects/%s/zones/%s/instances/%s", u.project, u.zone, u.instanceName), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *ComputeInstanceIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/zones/%s/instances/%s", u.project, u.zone, u.instanceName)
}

func (u *ComputeInstanceIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-compute-instance-%s", u.GetResourceId())
}

func (u *ComputeInstanceIamUpdater) DescribeResource() string {
	return fmt.Sprintf("compute instance %q", u.GetResourceId())
}
