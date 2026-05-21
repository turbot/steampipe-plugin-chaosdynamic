package main

import (
	chaos_dynamic "github.com/turbot/steampipe-plugin-chaosdynamic/chaosdynamic"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: chaos_dynamic.Plugin})
}
