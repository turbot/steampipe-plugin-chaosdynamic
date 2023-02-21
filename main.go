package main

import (
	chaos_dynamic "github.com/turbot/steampipe-plugin-chaos-dynamic/chaosdynamic"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: chaos_dynamic.Plugin})
}
