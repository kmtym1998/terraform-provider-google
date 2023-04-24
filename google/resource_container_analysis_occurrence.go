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

func ResourceContainerAnalysisOccurrence() *schema.Resource {
	return &schema.Resource{
		Create: resourceContainerAnalysisOccurrenceCreate,
		Read:   resourceContainerAnalysisOccurrenceRead,
		Update: resourceContainerAnalysisOccurrenceUpdate,
		Delete: resourceContainerAnalysisOccurrenceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceContainerAnalysisOccurrenceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"attestation": {
				Type:     schema.TypeList,
				Required: true,
				Description: `Occurrence that represents a single "attestation". The authenticity
of an attestation can be verified using the attached signature.
If the verifier trusts the public key of the signer, then verifying
the signature is sufficient to establish trust. In this circumstance,
the authority to which this attestation is attached is primarily
useful for lookup (how to find this attestation if you already
know the authority and artifact to be verified) and intent (for
which authority this attestation was intended to sign.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"serialized_payload": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The serialized payload that is verified by one or
more signatures. A base64-encoded string.`,
						},
						"signatures": {
							Type:     schema.TypeSet,
							Required: true,
							Description: `One or more signatures over serializedPayload.
Verifier implementations should consider this attestation
message verified if at least one signature verifies
serializedPayload. See Signature in common.proto for more
details on signature structure and verification.`,
							Elem: containeranalysisOccurrenceAttestationSignaturesSchema(),
							// Default schema.HashSchema is used.
						},
					},
				},
			},
			"note_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The analysis note associated with this occurrence, in the form of
projects/[PROJECT]/notes/[NOTE_ID]. This field can be used as a
filter in list requests.`,
			},
			"resource_uri": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Required. Immutable. A URI that represents the resource for which
the occurrence applies. For example,
https://gcr.io/project/image@sha256:123abc for a Docker image.`,
			},
			"remediation": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A description of actions that can be taken to remedy the note.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the repository was created.`,
			},
			"kind": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The note kind which explicitly denotes which of the occurrence
details are specified. This field can be used as a filter in list
requests.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the occurrence.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the repository was last updated.`,
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

func containeranalysisOccurrenceAttestationSignaturesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"public_key_id": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The identifier for the public key that verifies this
signature. MUST be an RFC3986 conformant
URI. * When possible, the key id should be an
immutable reference, such as a cryptographic digest.
Examples of valid values:

* OpenPGP V4 public key fingerprint. See https://www.iana.org/assignments/uri-schemes/prov/openpgp4fpr
  for more details on this scheme.
    * 'openpgp4fpr:74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA'
* RFC6920 digest-named SubjectPublicKeyInfo (digest of the DER serialization):
    * "ni:///sha-256;cD9o9Cq6LG3jD0iKXqEi_vdjJGecm_iXkbqVoScViaU"`,
			},
			"signature": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The content of the signature, an opaque bytestring.
The payload that this signature verifies MUST be
unambiguously provided with the Signature during
verification. A wrapper message might provide the
payload explicitly. Alternatively, a message might
have a canonical serialization that can always be
unambiguously computed to derive the payload.`,
			},
		},
	}
}

func resourceContainerAnalysisOccurrenceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	resourceUriProp, err := expandContainerAnalysisOccurrenceResourceUri(d.Get("resource_uri"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("resource_uri"); !isEmptyValue(reflect.ValueOf(resourceUriProp)) && (ok || !reflect.DeepEqual(v, resourceUriProp)) {
		obj["resourceUri"] = resourceUriProp
	}
	noteNameProp, err := expandContainerAnalysisOccurrenceNoteName(d.Get("note_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("note_name"); !isEmptyValue(reflect.ValueOf(noteNameProp)) && (ok || !reflect.DeepEqual(v, noteNameProp)) {
		obj["noteName"] = noteNameProp
	}
	remediationProp, err := expandContainerAnalysisOccurrenceRemediation(d.Get("remediation"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("remediation"); !isEmptyValue(reflect.ValueOf(remediationProp)) && (ok || !reflect.DeepEqual(v, remediationProp)) {
		obj["remediation"] = remediationProp
	}
	attestationProp, err := expandContainerAnalysisOccurrenceAttestation(d.Get("attestation"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attestation"); !isEmptyValue(reflect.ValueOf(attestationProp)) && (ok || !reflect.DeepEqual(v, attestationProp)) {
		obj["attestation"] = attestationProp
	}

	obj, err = resourceContainerAnalysisOccurrenceEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := ReplaceVars(d, config, "{{note_name}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{ContainerAnalysisBasePath}}projects/{{project}}/occurrences")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Occurrence: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Occurrence: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Occurrence: %s", err)
	}
	if err := d.Set("name", flattenContainerAnalysisOccurrenceName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/occurrences/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Occurrence %q: %#v", d.Id(), res)

	return resourceContainerAnalysisOccurrenceRead(d, meta)
}

func resourceContainerAnalysisOccurrenceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ContainerAnalysisBasePath}}projects/{{project}}/occurrences/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Occurrence: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ContainerAnalysisOccurrence %q", d.Id()))
	}

	res, err = resourceContainerAnalysisOccurrenceDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ContainerAnalysisOccurrence because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Occurrence: %s", err)
	}

	if err := d.Set("name", flattenContainerAnalysisOccurrenceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Occurrence: %s", err)
	}
	if err := d.Set("resource_uri", flattenContainerAnalysisOccurrenceResourceUri(res["resourceUri"], d, config)); err != nil {
		return fmt.Errorf("Error reading Occurrence: %s", err)
	}
	if err := d.Set("note_name", flattenContainerAnalysisOccurrenceNoteName(res["noteName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Occurrence: %s", err)
	}
	if err := d.Set("kind", flattenContainerAnalysisOccurrenceKind(res["kind"], d, config)); err != nil {
		return fmt.Errorf("Error reading Occurrence: %s", err)
	}
	if err := d.Set("remediation", flattenContainerAnalysisOccurrenceRemediation(res["remediation"], d, config)); err != nil {
		return fmt.Errorf("Error reading Occurrence: %s", err)
	}
	if err := d.Set("create_time", flattenContainerAnalysisOccurrenceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Occurrence: %s", err)
	}
	if err := d.Set("update_time", flattenContainerAnalysisOccurrenceUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Occurrence: %s", err)
	}
	if err := d.Set("attestation", flattenContainerAnalysisOccurrenceAttestation(res["attestation"], d, config)); err != nil {
		return fmt.Errorf("Error reading Occurrence: %s", err)
	}

	return nil
}

func resourceContainerAnalysisOccurrenceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Occurrence: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	remediationProp, err := expandContainerAnalysisOccurrenceRemediation(d.Get("remediation"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("remediation"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, remediationProp)) {
		obj["remediation"] = remediationProp
	}
	attestationProp, err := expandContainerAnalysisOccurrenceAttestation(d.Get("attestation"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attestation"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, attestationProp)) {
		obj["attestation"] = attestationProp
	}

	obj, err = resourceContainerAnalysisOccurrenceUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := ReplaceVars(d, config, "{{note_name}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{ContainerAnalysisBasePath}}projects/{{project}}/occurrences/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Occurrence %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("remediation") {
		updateMask = append(updateMask, "remediation")
	}

	if d.HasChange("attestation") {
		updateMask = append(updateMask, "attestation")
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
		return fmt.Errorf("Error updating Occurrence %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Occurrence %q: %#v", d.Id(), res)
	}

	return resourceContainerAnalysisOccurrenceRead(d, meta)
}

func resourceContainerAnalysisOccurrenceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Occurrence: %s", err)
	}
	billingProject = project

	lockName, err := ReplaceVars(d, config, "{{note_name}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{ContainerAnalysisBasePath}}projects/{{project}}/occurrences/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Occurrence %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Occurrence")
	}

	log.Printf("[DEBUG] Finished deleting Occurrence %q: %#v", d.Id(), res)
	return nil
}

func resourceContainerAnalysisOccurrenceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/occurrences/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/occurrences/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenContainerAnalysisOccurrenceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenContainerAnalysisOccurrenceResourceUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenContainerAnalysisOccurrenceNoteName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenContainerAnalysisOccurrenceKind(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenContainerAnalysisOccurrenceRemediation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenContainerAnalysisOccurrenceCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenContainerAnalysisOccurrenceUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenContainerAnalysisOccurrenceAttestation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["serialized_payload"] =
		flattenContainerAnalysisOccurrenceAttestationSerializedPayload(original["serializedPayload"], d, config)
	transformed["signatures"] =
		flattenContainerAnalysisOccurrenceAttestationSignatures(original["signatures"], d, config)
	return []interface{}{transformed}
}
func flattenContainerAnalysisOccurrenceAttestationSerializedPayload(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenContainerAnalysisOccurrenceAttestationSignatures(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(containeranalysisOccurrenceAttestationSignaturesSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"signature":     flattenContainerAnalysisOccurrenceAttestationSignaturesSignature(original["signature"], d, config),
			"public_key_id": flattenContainerAnalysisOccurrenceAttestationSignaturesPublicKeyId(original["publicKeyId"], d, config),
		})
	}
	return transformed
}
func flattenContainerAnalysisOccurrenceAttestationSignaturesSignature(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenContainerAnalysisOccurrenceAttestationSignaturesPublicKeyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandContainerAnalysisOccurrenceResourceUri(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisOccurrenceNoteName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisOccurrenceRemediation(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisOccurrenceAttestation(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSerializedPayload, err := expandContainerAnalysisOccurrenceAttestationSerializedPayload(original["serialized_payload"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSerializedPayload); val.IsValid() && !isEmptyValue(val) {
		transformed["serializedPayload"] = transformedSerializedPayload
	}

	transformedSignatures, err := expandContainerAnalysisOccurrenceAttestationSignatures(original["signatures"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignatures); val.IsValid() && !isEmptyValue(val) {
		transformed["signatures"] = transformedSignatures
	}

	return transformed, nil
}

func expandContainerAnalysisOccurrenceAttestationSerializedPayload(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisOccurrenceAttestationSignatures(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedSignature, err := expandContainerAnalysisOccurrenceAttestationSignaturesSignature(original["signature"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSignature); val.IsValid() && !isEmptyValue(val) {
			transformed["signature"] = transformedSignature
		}

		transformedPublicKeyId, err := expandContainerAnalysisOccurrenceAttestationSignaturesPublicKeyId(original["public_key_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPublicKeyId); val.IsValid() && !isEmptyValue(val) {
			transformed["publicKeyId"] = transformedPublicKeyId
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandContainerAnalysisOccurrenceAttestationSignaturesSignature(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisOccurrenceAttestationSignaturesPublicKeyId(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceContainerAnalysisOccurrenceEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// encoder logic only in non-GA versions

	return obj, nil
}

func resourceContainerAnalysisOccurrenceUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Note is required, even for PATCH
	noteNameProp, err := expandContainerAnalysisOccurrenceNoteName(d.Get("note_name"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("note_name"); !isEmptyValue(reflect.ValueOf(noteNameProp)) && (ok || !reflect.DeepEqual(v, noteNameProp)) {
		obj["noteName"] = noteNameProp
	}

	return resourceContainerAnalysisOccurrenceEncoder(d, meta, obj)
}

func resourceContainerAnalysisOccurrenceDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// encoder logic only in non-GA version
	return res, nil
}
