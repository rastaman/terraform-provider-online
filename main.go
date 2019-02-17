package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/rastaman/terraform-provider-online/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
