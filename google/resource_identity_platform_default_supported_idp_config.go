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

func ResourceIdentityPlatformDefaultSupportedIdpConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityPlatformDefaultSupportedIdpConfigCreate,
		Read:   resourceIdentityPlatformDefaultSupportedIdpConfigRead,
		Update: resourceIdentityPlatformDefaultSupportedIdpConfigUpdate,
		Delete: resourceIdentityPlatformDefaultSupportedIdpConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIdentityPlatformDefaultSupportedIdpConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `OAuth client ID`,
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `OAuth client secret`,
			},
			"idp_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `ID of the IDP. Possible values include:

* 'apple.com'

* 'facebook.com'

* 'gc.apple.com'

* 'github.com'

* 'google.com'

* 'linkedin.com'

* 'microsoft.com'

* 'playgames.google.com'

* 'twitter.com'

* 'yahoo.com'`,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If this IDP allows the user to sign in`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the DefaultSupportedIdpConfig resource`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceIdentityPlatformDefaultSupportedIdpConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	clientIdProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_id"); !isEmptyValue(reflect.ValueOf(clientIdProp)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_secret"); !isEmptyValue(reflect.ValueOf(clientSecretProp)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}
	enabledProp, err := expandIdentityPlatformDefaultSupportedIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !isEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}

	url, err := ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/defaultSupportedIdpConfigs?idpId={{idp_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DefaultSupportedIdpConfig: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating DefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("name", flattenIdentityPlatformDefaultSupportedIdpConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating DefaultSupportedIdpConfig %q: %#v", d.Id(), res)

	return resourceIdentityPlatformDefaultSupportedIdpConfigRead(d, meta)
}

func resourceIdentityPlatformDefaultSupportedIdpConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("IdentityPlatformDefaultSupportedIdpConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}

	if err := d.Set("name", flattenIdentityPlatformDefaultSupportedIdpConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("client_id", flattenIdentityPlatformDefaultSupportedIdpConfigClientId(res["clientId"], d, config)); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("client_secret", flattenIdentityPlatformDefaultSupportedIdpConfigClientSecret(res["clientSecret"], d, config)); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("enabled", flattenIdentityPlatformDefaultSupportedIdpConfigEnabled(res["enabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}

	return nil
}

func resourceIdentityPlatformDefaultSupportedIdpConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	clientIdProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_id"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_secret"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}
	enabledProp, err := expandIdentityPlatformDefaultSupportedIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}

	url, err := ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating DefaultSupportedIdpConfig %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("client_id") {
		updateMask = append(updateMask, "clientId")
	}

	if d.HasChange("client_secret") {
		updateMask = append(updateMask, "clientSecret")
	}

	if d.HasChange("enabled") {
		updateMask = append(updateMask, "enabled")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating DefaultSupportedIdpConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating DefaultSupportedIdpConfig %q: %#v", d.Id(), res)
	}

	return resourceIdentityPlatformDefaultSupportedIdpConfigRead(d, meta)
}

func resourceIdentityPlatformDefaultSupportedIdpConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting DefaultSupportedIdpConfig %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "DefaultSupportedIdpConfig")
	}

	log.Printf("[DEBUG] Finished deleting DefaultSupportedIdpConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceIdentityPlatformDefaultSupportedIdpConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/defaultSupportedIdpConfigs/(?P<idp_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<idp_id>[^/]+)",
		"(?P<idp_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIdentityPlatformDefaultSupportedIdpConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformDefaultSupportedIdpConfigClientId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformDefaultSupportedIdpConfigClientSecret(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformDefaultSupportedIdpConfigEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIdentityPlatformDefaultSupportedIdpConfigClientId(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformDefaultSupportedIdpConfigClientSecret(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformDefaultSupportedIdpConfigEnabled(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
