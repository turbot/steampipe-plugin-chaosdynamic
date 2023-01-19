package main

import (
	"github.com/turbot/steampipe-plugin-dyntest/dyntest"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: dyntest.Plugin})
}
