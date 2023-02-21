package chaos_dynamic

import (
	"context"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const pluginName = "steampipe-plugin-chaosdynamic"

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:         pluginName,
		TableMapFunc: buildTables,
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},

		SchemaMode: plugin.SchemaModeDynamic,
	}

	return p
}

func buildTables(ctx context.Context, d *plugin.TableMapData) (_ map[string]*plugin.Table, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = helpers.ToError(r)
		}
	}()

	tables := map[string]*plugin.Table{}
	config := getConfig(d.Connection)
	for _, t := range config.Tables {
		tables[t.Name] = t.AsPluginTable()
	}
	return tables, nil

}
