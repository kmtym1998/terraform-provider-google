---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file in
#     .github/CONTRIBUTING.md.
#
# ----------------------------------------------------------------------------
subcategory: "Network services"
description: |-
  ServiceBinding is the resource that defines a Service Directory Service to be used in a 
  BackendService resource.
---

# google\_network\_services\_service\_binding

ServiceBinding is the resource that defines a Service Directory Service to be used in a 
BackendService resource.

~> **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
See [Provider Versions](https://terraform.io/docs/providers/google/guides/provider_versions.html) for more details on beta resources.

To get more information about ServiceBinding, see:

* [API documentation](https://cloud.google.com/traffic-director/docs/reference/network-services/rest/v1beta1/projects.locations.serviceBindings)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=network_services_service_binding_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Network Services Service Binding Basic


```hcl
resource "google_service_directory_namespace" "default" {
  provider     = google-beta
  namespace_id = "my-namespace"
  location     = "us-central1"
}

resource "google_service_directory_service" "default" {
  provider   = google-beta
  service_id = "my-service"
  namespace  = google_service_directory_namespace.default.id

  metadata = {
    stage  = "prod"
    region = "us-central1"
  }
}

resource "google_network_services_service_binding" "default" {
  provider    = google-beta
  name        = "my-service-binding"
  labels      = {
    foo = "bar"
  }
  description = "my description"
  service = google_service_directory_service.default.id
}
```

## Argument Reference

The following arguments are supported:


* `service` -
  (Required)
  The full Service Directory Service name of the format 
  projects/*/locations/*/namespaces/*/services/*

* `name` -
  (Required)
  Name of the ServiceBinding resource.


- - -


* `labels` -
  (Optional)
  Set of label tags associated with the ServiceBinding resource.

* `description` -
  (Optional)
  A free-text description of the resource. Max length 1024 characters.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/global/serviceBindings/{{name}}`

* `create_time` -
  Time the ServiceBinding was created in UTC.

* `update_time` -
  Time the ServiceBinding was updated in UTC.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 10 minutes.
- `delete` - Default is 10 minutes.

## Import


ServiceBinding can be imported using any of these accepted formats:

```
$ terraform import google_network_services_service_binding.default projects/{{project}}/locations/global/serviceBindings/{{name}}
$ terraform import google_network_services_service_binding.default {{project}}/{{name}}
$ terraform import google_network_services_service_binding.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
