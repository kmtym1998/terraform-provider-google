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
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceVertexAIFeaturestoreEntitytype() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAIFeaturestoreEntitytypeCreate,
		Read:   resourceVertexAIFeaturestoreEntitytypeRead,
		Update: resourceVertexAIFeaturestoreEntitytypeUpdate,
		Delete: resourceVertexAIFeaturestoreEntitytypeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAIFeaturestoreEntitytypeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"featurestore": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Optional. Description of the EntityType.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `A set of key/value label pairs to assign to this EntityType.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"monitoring_config": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The default monitoring configuration for all Features under this EntityType.

If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"categorical_threshold_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Threshold for categorical features of anomaly detection. This is shared by all types of Featurestore Monitoring for categorical features (i.e. Features with type (Feature.ValueType) BOOL or STRING).`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type:        schema.TypeFloat,
										Required:    true,
										Description: `Specify a threshold value that can trigger the alert. For categorical feature, the distribution distance is calculated by L-inifinity norm. Each feature must have a non-zero threshold if they need to be monitored. Otherwise no alert will be triggered for that feature. The default value is 0.3.`,
									},
								},
							},
						},
						"import_features_analysis": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `The config for ImportFeatures Analysis Based Feature Monitoring.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"anomaly_detection_baseline": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `Defines the baseline to do anomaly detection for feature values imported by each [entityTypes.importFeatureValues][] operation. The value must be one of the values below:
* LATEST_STATS: Choose the later one statistics generated by either most recent snapshot analysis or previous import features analysis. If non of them exists, skip anomaly detection and only generate a statistics.
* MOST_RECENT_SNAPSHOT_STATS: Use the statistics generated by the most recent snapshot analysis if exists.
* PREVIOUS_IMPORT_FEATURES_STATS: Use the statistics generated by the previous import features analysis if exists.`,
									},
									"state": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `Whether to enable / disable / inherite default hebavior for import features analysis. The value must be one of the values below:
* DEFAULT: The default behavior of whether to enable the monitoring. EntityType-level config: disabled.
* ENABLED: Explicitly enables import features analysis. EntityType-level config: by default enables import features analysis for all Features under it.
* DISABLED: Explicitly disables import features analysis. EntityType-level config: by default disables import features analysis for all Features under it.`,
									},
								},
							},
						},
						"numerical_threshold_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Threshold for numerical features of anomaly detection. This is shared by all objectives of Featurestore Monitoring for numerical features (i.e. Features with type (Feature.ValueType) DOUBLE or INT64).`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type:        schema.TypeFloat,
										Required:    true,
										Description: `Specify a threshold value that can trigger the alert. For numerical feature, the distribution distance is calculated by Jensen–Shannon divergence. Each feature must have a non-zero threshold if they need to be monitored. Otherwise no alert will be triggered for that feature. The default value is 0.3.`,
									},
								},
							},
						},
						"snapshot_analysis": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `The config for Snapshot Analysis Based Feature Monitoring.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: `The monitoring schedule for snapshot analysis. For EntityType-level config: unset / disabled = true indicates disabled by default for Features under it; otherwise by default enable snapshot analysis monitoring with monitoringInterval for Features under it.`,
										Default:     false,
									},
									"monitoring_interval_days": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: `Configuration of the snapshot analysis based monitoring pipeline running interval. The value indicates number of days. The default value is 1.
If both FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days and [FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval][] are set when creating/updating EntityTypes/Features, FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days will be used.`,
										Default: 1,
									},
									"staleness_days": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: `Customized export features time window for snapshot analysis. Unit is one day. The default value is 21 days. Minimum value is 1 day. Maximum value is 4000 days.`,
										Default:     21,
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Used to perform consistent read-modify-write updates.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The region of the EntityType.",
			},
		},
		UseJSONNumber: true,
	}
}

func resourceVertexAIFeaturestoreEntitytypeCreate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandVertexAIFeaturestoreEntitytypeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandVertexAIFeaturestoreEntitytypeLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	monitoringConfigProp, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfig(d.Get("monitoring_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("monitoring_config"); !isEmptyValue(reflect.ValueOf(monitoringConfigProp)) && (ok || !reflect.DeepEqual(v, monitoringConfigProp)) {
		obj["monitoringConfig"] = monitoringConfigProp
	}

	obj, err = resourceVertexAIFeaturestoreEntitytypeEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{VertexAIBasePath}}{{featurestore}}/entityTypes?entityTypeId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FeaturestoreEntitytype: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	if v, ok := d.GetOk("featurestore"); ok {
		re := regexp.MustCompile("projects/([a-zA-Z0-9-]*)/(?:locations|regions)/([a-zA-Z0-9-]*)")
		switch {
		case re.MatchString(v.(string)):
			if res := re.FindStringSubmatch(v.(string)); len(res) == 3 && res[1] != "" {
				project = res[1]
			}
		}
	}
	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating FeaturestoreEntitytype: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{featurestore}}/entityTypes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = VertexAIOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating FeaturestoreEntitytype", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create FeaturestoreEntitytype: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "{{featurestore}}/entityTypes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating FeaturestoreEntitytype %q: %#v", d.Id(), res)

	return resourceVertexAIFeaturestoreEntitytypeRead(d, meta)
}

func resourceVertexAIFeaturestoreEntitytypeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{VertexAIBasePath}}{{featurestore}}/entityTypes/{{name}}")
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
		return handleNotFoundError(err, d, fmt.Sprintf("VertexAIFeaturestoreEntitytype %q", d.Id()))
	}

	if err := d.Set("description", flattenVertexAIFeaturestoreEntitytypeDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytype: %s", err)
	}
	if err := d.Set("create_time", flattenVertexAIFeaturestoreEntitytypeCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytype: %s", err)
	}
	if err := d.Set("update_time", flattenVertexAIFeaturestoreEntitytypeUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytype: %s", err)
	}
	if err := d.Set("labels", flattenVertexAIFeaturestoreEntitytypeLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytype: %s", err)
	}
	if err := d.Set("monitoring_config", flattenVertexAIFeaturestoreEntitytypeMonitoringConfig(res["monitoringConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytype: %s", err)
	}

	return nil
}

func resourceVertexAIFeaturestoreEntitytypeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandVertexAIFeaturestoreEntitytypeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandVertexAIFeaturestoreEntitytypeLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	monitoringConfigProp, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfig(d.Get("monitoring_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("monitoring_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, monitoringConfigProp)) {
		obj["monitoringConfig"] = monitoringConfigProp
	}

	obj, err = resourceVertexAIFeaturestoreEntitytypeEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{VertexAIBasePath}}{{featurestore}}/entityTypes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FeaturestoreEntitytype %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("monitoring_config") {
		updateMask = append(updateMask, "monitoringConfig")
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
		return fmt.Errorf("Error updating FeaturestoreEntitytype %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating FeaturestoreEntitytype %q: %#v", d.Id(), res)
	}

	return resourceVertexAIFeaturestoreEntitytypeRead(d, meta)
}

func resourceVertexAIFeaturestoreEntitytypeDelete(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{VertexAIBasePath}}{{featurestore}}/entityTypes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	if v, ok := d.GetOk("featurestore"); ok {
		re := regexp.MustCompile("projects/([a-zA-Z0-9-]*)/(?:locations|regions)/([a-zA-Z0-9-]*)")
		switch {
		case re.MatchString(v.(string)):
			if res := re.FindStringSubmatch(v.(string)); len(res) == 3 && res[1] != "" {
				project = res[1]
			}
		}
	}
	log.Printf("[DEBUG] Deleting FeaturestoreEntitytype %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "FeaturestoreEntitytype")
	}

	err = VertexAIOperationWaitTime(
		config, res, project, "Deleting FeaturestoreEntitytype", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting FeaturestoreEntitytype %q: %#v", d.Id(), res)
	return nil
}

func resourceVertexAIFeaturestoreEntitytypeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"(?P<featurestore>.+)/entityTypes/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "{{featurestore}}/entityTypes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	featurestore := d.Get("featurestore").(string)

	re := regexp.MustCompile("^projects/(.+)/locations/(.+)/featurestores/(.+)$")
	if parts := re.FindStringSubmatch(featurestore); parts != nil {
		d.Set("region", parts[2])
	}

	return []*schema.ResourceData{d}, nil
}

func flattenVertexAIFeaturestoreEntitytypeDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeMonitoringConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["snapshot_analysis"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis(original["snapshotAnalysis"], d, config)
	transformed["import_features_analysis"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis(original["importFeaturesAnalysis"], d, config)
	transformed["numerical_threshold_config"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfig(original["numericalThresholdConfig"], d, config)
	transformed["categorical_threshold_config"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfig(original["categoricalThresholdConfig"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["disabled"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisDisabled(original["disabled"], d, config)
	transformed["monitoring_interval_days"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisMonitoringIntervalDays(original["monitoringIntervalDays"], d, config)
	transformed["staleness_days"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisStalenessDays(original["stalenessDays"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisDisabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisMonitoringIntervalDays(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisStalenessDays(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["state"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisState(original["state"], d, config)
	transformed["anomaly_detection_baseline"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisAnomalyDetectionBaseline(original["anomalyDetectionBaseline"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisAnomalyDetectionBaseline(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["value"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfigValue(original["value"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfigValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["value"] =
		flattenVertexAIFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfigValue(original["value"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfigValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandVertexAIFeaturestoreEntitytypeDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreEntitytypeLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSnapshotAnalysis, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis(original["snapshot_analysis"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSnapshotAnalysis); val.IsValid() && !isEmptyValue(val) {
		transformed["snapshotAnalysis"] = transformedSnapshotAnalysis
	}

	transformedImportFeaturesAnalysis, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis(original["import_features_analysis"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImportFeaturesAnalysis); val.IsValid() && !isEmptyValue(val) {
		transformed["importFeaturesAnalysis"] = transformedImportFeaturesAnalysis
	}

	transformedNumericalThresholdConfig, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfig(original["numerical_threshold_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNumericalThresholdConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["numericalThresholdConfig"] = transformedNumericalThresholdConfig
	}

	transformedCategoricalThresholdConfig, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfig(original["categorical_threshold_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCategoricalThresholdConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["categoricalThresholdConfig"] = transformedCategoricalThresholdConfig
	}

	return transformed, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDisabled, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisDisabled(original["disabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisabled); val.IsValid() && !isEmptyValue(val) {
		transformed["disabled"] = transformedDisabled
	}

	transformedMonitoringIntervalDays, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisMonitoringIntervalDays(original["monitoring_interval_days"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonitoringIntervalDays); val.IsValid() && !isEmptyValue(val) {
		transformed["monitoringIntervalDays"] = transformedMonitoringIntervalDays
	}

	transformedStalenessDays, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisStalenessDays(original["staleness_days"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStalenessDays); val.IsValid() && !isEmptyValue(val) {
		transformed["stalenessDays"] = transformedStalenessDays
	}

	return transformed, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisDisabled(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisMonitoringIntervalDays(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysisStalenessDays(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedState, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !isEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedAnomalyDetectionBaseline, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisAnomalyDetectionBaseline(original["anomaly_detection_baseline"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAnomalyDetectionBaseline); val.IsValid() && !isEmptyValue(val) {
		transformed["anomalyDetectionBaseline"] = transformedAnomalyDetectionBaseline
	}

	return transformed, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisState(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysisAnomalyDetectionBaseline(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedValue, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfigValue(original["value"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
		transformed["value"] = transformedValue
	}

	return transformed, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfigValue(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedValue, err := expandVertexAIFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfigValue(original["value"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
		transformed["value"] = transformedValue
	}

	return transformed, nil
}

func expandVertexAIFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfigValue(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceVertexAIFeaturestoreEntitytypeEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	if v, ok := d.GetOk("featurestore"); ok {
		re := regexp.MustCompile("projects/(.+)/locations/(.+)/featurestores/(.+)$")
		if parts := re.FindStringSubmatch(v.(string)); parts != nil {
			d.Set("region", parts[2])
		}
	}

	return obj, nil
}
