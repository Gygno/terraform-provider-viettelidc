// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package voks

import (
	"terraform-provider-viettelidc/internal/provider"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var (
	providerConfig = `
provider "viettelidc" {
  domain_id = "9e9480cc-96aa-446e-b08b-5cd7b2f438ab"
  username = "iac"
  password = "Vtdc@12345"
  mfa_code = 123456
}
`
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"viettelidc": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}
