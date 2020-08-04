package main

import (
	"github.com/dedunu/terraform-provider-awslambda/dedunu"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return dedunu.Provider()
		},
	})
}
