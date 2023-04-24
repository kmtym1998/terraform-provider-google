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

func ResourceApigeeEnvgroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeEnvgroupCreate,
		Read:   resourceApigeeEnvgroupRead,
		Update: resourceApigeeEnvgroupUpdate,
		Delete: resourceApigeeEnvgroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeEnvgroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource ID of the environment group.`,
			},
			"org_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The Apigee Organization associated with the Apigee environment group,
in the format 'organizations/{{org_name}}'.`,
			},
			"hostnames": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Hostnames of the environment group.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeEnvgroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandApigeeEnvgroupName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	hostnamesProp, err := expandApigeeEnvgroupHostnames(d.Get("hostnames"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("hostnames"); !isEmptyValue(reflect.ValueOf(hostnamesProp)) && (ok || !reflect.DeepEqual(v, hostnamesProp)) {
		obj["hostnames"] = hostnamesProp
	}

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/envgroups")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Envgroup: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Envgroup: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{org_id}}/envgroups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ApigeeOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating Envgroup", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Envgroup: %s", err)
	}

	if err := d.Set("name", flattenApigeeEnvgroupName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "{{org_id}}/envgroups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Envgroup %q: %#v", d.Id(), res)

	return resourceApigeeEnvgroupRead(d, meta)
}

func resourceApigeeEnvgroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/envgroups/{{name}}")
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
		return handleNotFoundError(err, d, fmt.Sprintf("ApigeeEnvgroup %q", d.Id()))
	}

	if err := d.Set("name", flattenApigeeEnvgroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Envgroup: %s", err)
	}
	if err := d.Set("hostnames", flattenApigeeEnvgroupHostnames(res["hostnames"], d, config)); err != nil {
		return fmt.Errorf("Error reading Envgroup: %s", err)
	}

	return nil
}

func resourceApigeeEnvgroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	hostnamesProp, err := expandApigeeEnvgroupHostnames(d.Get("hostnames"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("hostnames"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, hostnamesProp)) {
		obj["hostnames"] = hostnamesProp
	}

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/envgroups/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Envgroup %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("hostnames") {
		updateMask = append(updateMask, "hostnames")
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
		return fmt.Errorf("Error updating Envgroup %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Envgroup %q: %#v", d.Id(), res)
	}

	err = ApigeeOperationWaitTime(
		config, res, "Updating Envgroup", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceApigeeEnvgroupRead(d, meta)
}

func resourceApigeeEnvgroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/envgroups/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Envgroup %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Envgroup")
	}

	err = ApigeeOperationWaitTime(
		config, res, "Deleting Envgroup", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Envgroup %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeEnvgroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	nameParts := strings.Split(d.Get("name").(string), "/")
	if len(nameParts) == 4 {
		// `organizations/{{org_name}}/envgroups/{{name}}`
		orgId := fmt.Sprintf("organizations/%s", nameParts[1])
		if err := d.Set("org_id", orgId); err != nil {
			return nil, fmt.Errorf("Error setting org_id: %s", err)
		}
		if err := d.Set("name", nameParts[3]); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
	} else if len(nameParts) == 3 {
		// `organizations/{{org_name}}/{{name}}`
		orgId := fmt.Sprintf("organizations/%s", nameParts[1])
		if err := d.Set("org_id", orgId); err != nil {
			return nil, fmt.Errorf("Error setting org_id: %s", err)
		}
		if err := d.Set("name", nameParts[2]); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
	} else {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s or %s",
			d.Get("name"),
			"organizations/{{org_name}}/envgroups/{{name}}",
			"organizations/{{org_name}}/{{name}}")
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "{{org_id}}/envgroups/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeEnvgroupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeEnvgroupHostnames(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApigeeEnvgroupName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvgroupHostnames(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
