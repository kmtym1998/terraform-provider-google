// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"log"
	"testing"

	containerazure "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func init() {
	resource.AddTestSweepers("ContainerAzureClient", &resource.Sweeper{
		Name: "ContainerAzureClient",
		F:    testSweepContainerAzureClient,
	})
}

func testSweepContainerAzureClient(region string) error {
	log.Print("[INFO][SWEEPER_LOG] Starting sweeper for ContainerAzureClient")

	config, err := SharedConfigForRegion(region)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error getting shared config for region: %s", err)
		return err
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error loading: %s", err)
		return err
	}

	t := &testing.T{}
	billingId := GetTestBillingAccountFromEnv(t)

	// Setup variables to be used for Delete arguments.
	d := map[string]string{
		"project":         config.Project,
		"region":          region,
		"location":        region,
		"zone":            "-",
		"billing_account": billingId,
	}

	client := transport_tpg.NewDCLContainerAzureClient(config, config.UserAgent, "", 0)
	err = client.DeleteAllClient(context.Background(), d["project"], d["location"], isDeletableContainerAzureClient)
	if err != nil {
		return err
	}
	return nil
}

func isDeletableContainerAzureClient(r *containerazure.AzureClient) bool {
	return IsSweepableTestResource(*r.Name)
}
