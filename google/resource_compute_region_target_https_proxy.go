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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceComputeRegionTargetHttpsProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionTargetHttpsProxyCreate,
		Read:   resourceComputeRegionTargetHttpsProxyRead,
		Update: resourceComputeRegionTargetHttpsProxyUpdate,
		Delete: resourceComputeRegionTargetHttpsProxyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionTargetHttpsProxyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"ssl_certificates": {
				Type:     schema.TypeList,
				Required: true,
				Description: `A list of RegionSslCertificate resources that are used to authenticate
connections between users and the load balancer. Currently, exactly
one SSL certificate must be specified.`,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: compareSelfLinkOrResourceName,
				},
			},
			"url_map": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `A reference to the RegionUrlMap resource that defines the mapping from URL
to the RegionBackendService.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The Region in which the created target https proxy should reside.
If it is not provided, the provider region is used.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"proxy_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The unique identifier for the resource.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeRegionTargetHttpsProxyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeRegionTargetHttpsProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeRegionTargetHttpsProxyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	sslCertificatesProp, err := expandComputeRegionTargetHttpsProxySslCertificates(d.Get("ssl_certificates"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_certificates"); !isEmptyValue(reflect.ValueOf(sslCertificatesProp)) && (ok || !reflect.DeepEqual(v, sslCertificatesProp)) {
		obj["sslCertificates"] = sslCertificatesProp
	}
	urlMapProp, err := expandComputeRegionTargetHttpsProxyUrlMap(d.Get("url_map"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("url_map"); !isEmptyValue(reflect.ValueOf(urlMapProp)) && (ok || !reflect.DeepEqual(v, urlMapProp)) {
		obj["urlMap"] = urlMapProp
	}
	regionProp, err := expandComputeRegionTargetHttpsProxyRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetHttpsProxies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionTargetHttpsProxy: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionTargetHttpsProxy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating RegionTargetHttpsProxy: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating RegionTargetHttpsProxy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionTargetHttpsProxy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating RegionTargetHttpsProxy %q: %#v", d.Id(), res)

	return resourceComputeRegionTargetHttpsProxyRead(d, meta)
}

func resourceComputeRegionTargetHttpsProxyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionTargetHttpsProxy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRegionTargetHttpsProxy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpsProxy: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeRegionTargetHttpsProxyCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpsProxy: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionTargetHttpsProxyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpsProxy: %s", err)
	}
	if err := d.Set("proxy_id", flattenComputeRegionTargetHttpsProxyProxyId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpsProxy: %s", err)
	}
	if err := d.Set("name", flattenComputeRegionTargetHttpsProxyName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpsProxy: %s", err)
	}
	if err := d.Set("ssl_certificates", flattenComputeRegionTargetHttpsProxySslCertificates(res["sslCertificates"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpsProxy: %s", err)
	}
	if err := d.Set("url_map", flattenComputeRegionTargetHttpsProxyUrlMap(res["urlMap"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpsProxy: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionTargetHttpsProxyRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpsProxy: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading RegionTargetHttpsProxy: %s", err)
	}

	return nil
}

func resourceComputeRegionTargetHttpsProxyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionTargetHttpsProxy: %s", err)
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("ssl_certificates") {
		obj := make(map[string]interface{})

		sslCertificatesProp, err := expandComputeRegionTargetHttpsProxySslCertificates(d.Get("ssl_certificates"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("ssl_certificates"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sslCertificatesProp)) {
			obj["sslCertificates"] = sslCertificatesProp
		}

		url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}/setSslCertificates")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating RegionTargetHttpsProxy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating RegionTargetHttpsProxy %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating RegionTargetHttpsProxy", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("url_map") {
		obj := make(map[string]interface{})

		urlMapProp, err := expandComputeRegionTargetHttpsProxyUrlMap(d.Get("url_map"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("url_map"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, urlMapProp)) {
			obj["urlMap"] = urlMapProp
		}

		url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}/setUrlMap")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating RegionTargetHttpsProxy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating RegionTargetHttpsProxy %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating RegionTargetHttpsProxy", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceComputeRegionTargetHttpsProxyRead(d, meta)
}

func resourceComputeRegionTargetHttpsProxyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionTargetHttpsProxy: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionTargetHttpsProxy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "RegionTargetHttpsProxy")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting RegionTargetHttpsProxy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionTargetHttpsProxy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionTargetHttpsProxyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/targetHttpsProxies/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionTargetHttpsProxyCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionTargetHttpsProxyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionTargetHttpsProxyProxyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeRegionTargetHttpsProxyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionTargetHttpsProxySslCertificates(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return convertAndMapStringArr(v.([]interface{}), ConvertSelfLinkToV1)
}

func flattenComputeRegionTargetHttpsProxyUrlMap(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRegionTargetHttpsProxyRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandComputeRegionTargetHttpsProxyDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionTargetHttpsProxyName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionTargetHttpsProxySslCertificates(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for ssl_certificates: nil")
		}
		f, err := parseRegionalFieldValue("sslCertificates", raw.(string), "project", "region", "zone", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for ssl_certificates: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeRegionTargetHttpsProxyUrlMap(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("urlMaps", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url_map: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionTargetHttpsProxyRegion(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
