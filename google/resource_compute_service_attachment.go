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

func ResourceComputeServiceAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeServiceAttachmentCreate,
		Read:   resourceComputeServiceAttachmentRead,
		Update: resourceComputeServiceAttachmentUpdate,
		Delete: resourceComputeServiceAttachmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeServiceAttachmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"connection_preference": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The connection preference to use for this service attachment. Valid
values include "ACCEPT_AUTOMATIC", "ACCEPT_MANUAL".`,
			},
			"enable_proxy_protocol": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
				Description: `If true, enable the proxy protocol which is for supplying client TCP/IP
address data in TCP connections that traverse proxies on their way to
destination servers.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. The name must be 1-63 characters long, and
comply with RFC1035. Specifically, the name must be 1-63 characters
long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'
which means the first character must be a lowercase letter, and all
following characters must be a dash, lowercase letter, or digit,
except the last character, which cannot be a dash.`,
			},
			"nat_subnets": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `An array of subnets that is provided for NAT in this service attachment.`,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: compareSelfLinkOrResourceName,
				},
			},
			"target_service": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The URL of a forwarding rule that represents the service identified by
this service attachment.`,
			},
			"consumer_accept_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `An array of projects that are allowed to connect to this service
attachment.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_limit": {
							Type:     schema.TypeInt,
							Required: true,
							Description: `The number of consumer forwarding rules the consumer project can
create.`,
						},
						"project_id_or_num": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `A project that is allowed to connect to this service attachment.`,
						},
					},
				},
			},
			"consumer_reject_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `An array of projects that are not allowed to connect to this service
attachment.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of this resource.`,
			},
			"domain_names": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `If specified, the domain name will be used during the integration between
the PSC connected endpoints and the Cloud DNS. For example, this is a
valid domain name: "p.mycompany.com.". Current max number of domain names
supported is 1.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `URL of the region where the resource resides.`,
			},
			"connected_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `An array of the consumer forwarding rules connected to this service
attachment.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The URL of the consumer forwarding rule.`,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The status of the connection from the consumer forwarding rule to
this service attachment.`,
						},
					},
				},
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Fingerprint of this resource. This field is used internally during
updates of this resource.`,
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

func resourceComputeServiceAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeServiceAttachmentName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeServiceAttachmentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeServiceAttachmentFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	connectionPreferenceProp, err := expandComputeServiceAttachmentConnectionPreference(d.Get("connection_preference"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("connection_preference"); !isEmptyValue(reflect.ValueOf(connectionPreferenceProp)) && (ok || !reflect.DeepEqual(v, connectionPreferenceProp)) {
		obj["connectionPreference"] = connectionPreferenceProp
	}
	targetServiceProp, err := expandComputeServiceAttachmentTargetService(d.Get("target_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_service"); !isEmptyValue(reflect.ValueOf(targetServiceProp)) && (ok || !reflect.DeepEqual(v, targetServiceProp)) {
		obj["targetService"] = targetServiceProp
	}
	natSubnetsProp, err := expandComputeServiceAttachmentNatSubnets(d.Get("nat_subnets"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("nat_subnets"); ok || !reflect.DeepEqual(v, natSubnetsProp) {
		obj["natSubnets"] = natSubnetsProp
	}
	enableProxyProtocolProp, err := expandComputeServiceAttachmentEnableProxyProtocol(d.Get("enable_proxy_protocol"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_proxy_protocol"); !isEmptyValue(reflect.ValueOf(enableProxyProtocolProp)) && (ok || !reflect.DeepEqual(v, enableProxyProtocolProp)) {
		obj["enableProxyProtocol"] = enableProxyProtocolProp
	}
	domainNamesProp, err := expandComputeServiceAttachmentDomainNames(d.Get("domain_names"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("domain_names"); !isEmptyValue(reflect.ValueOf(domainNamesProp)) && (ok || !reflect.DeepEqual(v, domainNamesProp)) {
		obj["domainNames"] = domainNamesProp
	}
	consumerRejectListsProp, err := expandComputeServiceAttachmentConsumerRejectLists(d.Get("consumer_reject_lists"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("consumer_reject_lists"); ok || !reflect.DeepEqual(v, consumerRejectListsProp) {
		obj["consumerRejectLists"] = consumerRejectListsProp
	}
	consumerAcceptListsProp, err := expandComputeServiceAttachmentConsumerAcceptLists(d.Get("consumer_accept_lists"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("consumer_accept_lists"); ok || !reflect.DeepEqual(v, consumerAcceptListsProp) {
		obj["consumerAcceptLists"] = consumerAcceptListsProp
	}
	regionProp, err := expandComputeServiceAttachmentRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/serviceAttachments")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServiceAttachment: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceAttachment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ServiceAttachment: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating ServiceAttachment", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ServiceAttachment: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ServiceAttachment %q: %#v", d.Id(), res)

	return resourceComputeServiceAttachmentRead(d, meta)
}

func resourceComputeServiceAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceAttachment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeServiceAttachment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}

	if err := d.Set("name", flattenComputeServiceAttachmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}
	if err := d.Set("description", flattenComputeServiceAttachmentDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeServiceAttachmentFingerprint(res["fingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}
	if err := d.Set("connection_preference", flattenComputeServiceAttachmentConnectionPreference(res["connectionPreference"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}
	if err := d.Set("connected_endpoints", flattenComputeServiceAttachmentConnectedEndpoints(res["connectedEndpoints"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}
	if err := d.Set("target_service", flattenComputeServiceAttachmentTargetService(res["targetService"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}
	if err := d.Set("nat_subnets", flattenComputeServiceAttachmentNatSubnets(res["natSubnets"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}
	if err := d.Set("enable_proxy_protocol", flattenComputeServiceAttachmentEnableProxyProtocol(res["enableProxyProtocol"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}
	if err := d.Set("domain_names", flattenComputeServiceAttachmentDomainNames(res["domainNames"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}
	if err := d.Set("consumer_reject_lists", flattenComputeServiceAttachmentConsumerRejectLists(res["consumerRejectLists"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}
	if err := d.Set("consumer_accept_lists", flattenComputeServiceAttachmentConsumerAcceptLists(res["consumerAcceptLists"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}
	if err := d.Set("region", flattenComputeServiceAttachmentRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading ServiceAttachment: %s", err)
	}

	return nil
}

func resourceComputeServiceAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceAttachment: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeServiceAttachmentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeServiceAttachmentFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	connectionPreferenceProp, err := expandComputeServiceAttachmentConnectionPreference(d.Get("connection_preference"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("connection_preference"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, connectionPreferenceProp)) {
		obj["connectionPreference"] = connectionPreferenceProp
	}
	natSubnetsProp, err := expandComputeServiceAttachmentNatSubnets(d.Get("nat_subnets"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("nat_subnets"); ok || !reflect.DeepEqual(v, natSubnetsProp) {
		obj["natSubnets"] = natSubnetsProp
	}
	consumerRejectListsProp, err := expandComputeServiceAttachmentConsumerRejectLists(d.Get("consumer_reject_lists"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("consumer_reject_lists"); ok || !reflect.DeepEqual(v, consumerRejectListsProp) {
		obj["consumerRejectLists"] = consumerRejectListsProp
	}
	consumerAcceptListsProp, err := expandComputeServiceAttachmentConsumerAcceptLists(d.Get("consumer_accept_lists"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("consumer_accept_lists"); ok || !reflect.DeepEqual(v, consumerAcceptListsProp) {
		obj["consumerAcceptLists"] = consumerAcceptListsProp
	}

	obj, err = resourceComputeServiceAttachmentUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ServiceAttachment %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ServiceAttachment %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ServiceAttachment %q: %#v", d.Id(), res)
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Updating ServiceAttachment", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeServiceAttachmentRead(d, meta)
}

func resourceComputeServiceAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceAttachment: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ServiceAttachment %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ServiceAttachment")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting ServiceAttachment", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ServiceAttachment %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeServiceAttachmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/serviceAttachments/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeServiceAttachmentName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeServiceAttachmentDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeServiceAttachmentFingerprint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeServiceAttachmentConnectionPreference(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeServiceAttachmentConnectedEndpoints(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"endpoint": flattenComputeServiceAttachmentConnectedEndpointsEndpoint(original["endpoint"], d, config),
			"status":   flattenComputeServiceAttachmentConnectedEndpointsStatus(original["status"], d, config),
		})
	}
	return transformed
}
func flattenComputeServiceAttachmentConnectedEndpointsEndpoint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeServiceAttachmentConnectedEndpointsStatus(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeServiceAttachmentTargetService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeServiceAttachmentNatSubnets(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return convertAndMapStringArr(v.([]interface{}), ConvertSelfLinkToV1)
}

func flattenComputeServiceAttachmentEnableProxyProtocol(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeServiceAttachmentDomainNames(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeServiceAttachmentConsumerRejectLists(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeServiceAttachmentConsumerAcceptLists(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"project_id_or_num": flattenComputeServiceAttachmentConsumerAcceptListsProjectIdOrNum(original["projectIdOrNum"], d, config),
			"connection_limit":  flattenComputeServiceAttachmentConsumerAcceptListsConnectionLimit(original["connectionLimit"], d, config),
		})
	}
	return transformed
}
func flattenComputeServiceAttachmentConsumerAcceptListsProjectIdOrNum(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeServiceAttachmentConsumerAcceptListsConnectionLimit(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeServiceAttachmentRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeServiceAttachmentName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentFingerprint(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentConnectionPreference(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentTargetService(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("forwardingRules", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for target_service: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeServiceAttachmentNatSubnets(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for nat_subnets: nil")
		}
		f, err := parseRegionalFieldValue("subnetworks", raw.(string), "project", "region", "zone", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for nat_subnets: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeServiceAttachmentEnableProxyProtocol(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentDomainNames(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentConsumerRejectLists(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentConsumerAcceptLists(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedProjectIdOrNum, err := expandComputeServiceAttachmentConsumerAcceptListsProjectIdOrNum(original["project_id_or_num"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProjectIdOrNum); val.IsValid() && !isEmptyValue(val) {
			transformed["projectIdOrNum"] = transformedProjectIdOrNum
		}

		transformedConnectionLimit, err := expandComputeServiceAttachmentConsumerAcceptListsConnectionLimit(original["connection_limit"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedConnectionLimit); val.IsValid() && !isEmptyValue(val) {
			transformed["connectionLimit"] = transformedConnectionLimit
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeServiceAttachmentConsumerAcceptListsProjectIdOrNum(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentConsumerAcceptListsConnectionLimit(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentRegion(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}

func resourceComputeServiceAttachmentUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {

	// need to send value in PATCH due to validation bug on api b/198329756
	nameProp := d.Get("name")
	if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	// need to send value in PATCH due to validation bug on api b/198308475
	enableProxyProtocolProp := d.Get("enable_proxy_protocol")
	if v, ok := d.GetOkExists("enable_proxy_protocol"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableProxyProtocolProp)) {
		obj["enableProxyProtocol"] = enableProxyProtocolProp
	}

	return obj, nil
}
