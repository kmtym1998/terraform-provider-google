// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceApigeeAddonsConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeAddonsConfigCreate,
		Read:   resourceApigeeAddonsConfigRead,
		Update: resourceApigeeAddonsConfigUpdate,
		Delete: resourceApigeeAddonsConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeAddonsConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"org": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the Apigee organization.`,
			},
			"addons_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Addon configurations of the Apigee organization.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advanced_api_ops_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration for the Monetization add-on.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: `Flag that specifies whether the Advanced API Ops add-on is enabled.`,
									},
								},
							},
						},
						"api_security_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration for the Monetization add-on.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: `Flag that specifies whether the Advanced API Ops add-on is enabled.`,
									},
									"expires_at": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Flag that specifies whether the Advanced API Ops add-on is enabled.`,
									},
								},
							},
						},
						"connectors_platform_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration for the Monetization add-on.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: `Flag that specifies whether the Advanced API Ops add-on is enabled.`,
									},
									"expires_at": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Flag that specifies whether the Advanced API Ops add-on is enabled.`,
									},
								},
							},
						},
						"integration_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration for the Monetization add-on.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: `Flag that specifies whether the Advanced API Ops add-on is enabled.`,
									},
								},
							},
						},
						"monetization_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration for the Monetization add-on.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: `Flag that specifies whether the Advanced API Ops add-on is enabled.`,
									},
								},
							},
						},
					},
				},
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeAddonsConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	addonsConfigProp, err := expandApigeeAddonsConfigAddonsConfig(d.Get("addons_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("addons_config"); !isEmptyValue(reflect.ValueOf(addonsConfigProp)) && (ok || !reflect.DeepEqual(v, addonsConfigProp)) {
		obj["addonsConfig"] = addonsConfigProp
	}

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org}}:setAddons")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AddonsConfig: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating AddonsConfig: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "organizations/{{org}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ApigeeOperationWaitTime(
		config, res, "Creating AddonsConfig", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create AddonsConfig: %s", err)
	}

	log.Printf("[DEBUG] Finished creating AddonsConfig %q: %#v", d.Id(), res)

	return resourceApigeeAddonsConfigRead(d, meta)
}

func resourceApigeeAddonsConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ApigeeAddonsConfig %q", d.Id()))
	}

	if err := d.Set("addons_config", flattenApigeeAddonsConfigAddonsConfig(res["addonsConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading AddonsConfig: %s", err)
	}

	return nil
}

func resourceApigeeAddonsConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	addonsConfigProp, err := expandApigeeAddonsConfigAddonsConfig(d.Get("addons_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("addons_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, addonsConfigProp)) {
		obj["addonsConfig"] = addonsConfigProp
	}

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org}}:setAddons")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AddonsConfig %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating AddonsConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AddonsConfig %q: %#v", d.Id(), res)
	}

	err = ApigeeOperationWaitTime(
		config, res, "Updating AddonsConfig", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceApigeeAddonsConfigRead(d, meta)
}

func resourceApigeeAddonsConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org}}:setAddons")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting AddonsConfig %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "AddonsConfig")
	}

	err = ApigeeOperationWaitTime(
		config, res, "Deleting AddonsConfig", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting AddonsConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeAddonsConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := ParseImportId([]string{"(?P<org>.+)"}, d, config); err != nil {
		return nil, err
	}

	parts := strings.Split(d.Get("org").(string), "/")

	var projectId string
	switch len(parts) {
	case 1:
		projectId = parts[0]
	case 2:
		projectId = parts[1]
	default:
		return nil, fmt.Errorf(
			"Saw %s when the org is expected to have shape %s or %s",
			d.Get("org"),
			"{{org}}",
			"organizations/{{org}}",
		)
	}

	if err := d.Set("org", projectId); err != nil {
		return nil, fmt.Errorf("Error setting organization: %s", err)
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "organizations/{{org}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeAddonsConfigAddonsConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["advanced_api_ops_config"] =
		flattenApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig(original["advancedApiOpsConfig"], d, config)
	transformed["integration_config"] =
		flattenApigeeAddonsConfigAddonsConfigIntegrationConfig(original["integrationConfig"], d, config)
	transformed["monetization_config"] =
		flattenApigeeAddonsConfigAddonsConfigMonetizationConfig(original["monetizationConfig"], d, config)
	transformed["api_security_config"] =
		flattenApigeeAddonsConfigAddonsConfigApiSecurityConfig(original["apiSecurityConfig"], d, config)
	transformed["connectors_platform_config"] =
		flattenApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig(original["connectorsPlatformConfig"], d, config)
	return []interface{}{transformed}
}
func flattenApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enabled"] =
		flattenApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfigEnabled(original["enabled"], d, config)
	return []interface{}{transformed}
}
func flattenApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfigEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeAddonsConfigAddonsConfigIntegrationConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enabled"] =
		flattenApigeeAddonsConfigAddonsConfigIntegrationConfigEnabled(original["enabled"], d, config)
	return []interface{}{transformed}
}
func flattenApigeeAddonsConfigAddonsConfigIntegrationConfigEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeAddonsConfigAddonsConfigMonetizationConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enabled"] =
		flattenApigeeAddonsConfigAddonsConfigMonetizationConfigEnabled(original["enabled"], d, config)
	return []interface{}{transformed}
}
func flattenApigeeAddonsConfigAddonsConfigMonetizationConfigEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeAddonsConfigAddonsConfigApiSecurityConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enabled"] =
		flattenApigeeAddonsConfigAddonsConfigApiSecurityConfigEnabled(original["enabled"], d, config)
	transformed["expires_at"] =
		flattenApigeeAddonsConfigAddonsConfigApiSecurityConfigExpiresAt(original["expiresAt"], d, config)
	return []interface{}{transformed}
}
func flattenApigeeAddonsConfigAddonsConfigApiSecurityConfigEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeAddonsConfigAddonsConfigApiSecurityConfigExpiresAt(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enabled"] =
		flattenApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigEnabled(original["enabled"], d, config)
	transformed["expires_at"] =
		flattenApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigExpiresAt(original["expiresAt"], d, config)
	return []interface{}{transformed}
}
func flattenApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigExpiresAt(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApigeeAddonsConfigAddonsConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAdvancedApiOpsConfig, err := expandApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig(original["advanced_api_ops_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdvancedApiOpsConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["advancedApiOpsConfig"] = transformedAdvancedApiOpsConfig
	}

	transformedIntegrationConfig, err := expandApigeeAddonsConfigAddonsConfigIntegrationConfig(original["integration_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIntegrationConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["integrationConfig"] = transformedIntegrationConfig
	}

	transformedMonetizationConfig, err := expandApigeeAddonsConfigAddonsConfigMonetizationConfig(original["monetization_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonetizationConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["monetizationConfig"] = transformedMonetizationConfig
	}

	transformedApiSecurityConfig, err := expandApigeeAddonsConfigAddonsConfigApiSecurityConfig(original["api_security_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedApiSecurityConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["apiSecurityConfig"] = transformedApiSecurityConfig
	}

	transformedConnectorsPlatformConfig, err := expandApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig(original["connectors_platform_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConnectorsPlatformConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["connectorsPlatformConfig"] = transformedConnectorsPlatformConfig
	}

	return transformed, nil
}

func expandApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !isEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfigEnabled(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAddonsConfigAddonsConfigIntegrationConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandApigeeAddonsConfigAddonsConfigIntegrationConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !isEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandApigeeAddonsConfigAddonsConfigIntegrationConfigEnabled(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAddonsConfigAddonsConfigMonetizationConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandApigeeAddonsConfigAddonsConfigMonetizationConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !isEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandApigeeAddonsConfigAddonsConfigMonetizationConfigEnabled(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAddonsConfigAddonsConfigApiSecurityConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandApigeeAddonsConfigAddonsConfigApiSecurityConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !isEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedExpiresAt, err := expandApigeeAddonsConfigAddonsConfigApiSecurityConfigExpiresAt(original["expires_at"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpiresAt); val.IsValid() && !isEmptyValue(val) {
		transformed["expiresAt"] = transformedExpiresAt
	}

	return transformed, nil
}

func expandApigeeAddonsConfigAddonsConfigApiSecurityConfigEnabled(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAddonsConfigAddonsConfigApiSecurityConfigExpiresAt(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !isEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedExpiresAt, err := expandApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigExpiresAt(original["expires_at"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpiresAt); val.IsValid() && !isEmptyValue(val) {
		transformed["expiresAt"] = transformedExpiresAt
	}

	return transformed, nil
}

func expandApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigEnabled(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigExpiresAt(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
