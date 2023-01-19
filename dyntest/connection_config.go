package dyntest

import (
	"context"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type DyntestConfig struct {
	Tables []configTable `hcl:"tables"`
}

func (d DyntestConfig) String() string {
	return fmt.Sprintf("TABLES: %v", d.Tables)
}

type configColumn struct {
	Name string `hcl:"name" cty:"name"`
	Type string `hcl:"type" cty:"type"`
}

func (c configColumn) getType() proto.ColumnType {
	switch c.Type {
	case "bool":
		return proto.ColumnType_BOOL
	case "int":
		return proto.ColumnType_INT
	case "double":
		return proto.ColumnType_DOUBLE
	case "string":
		return proto.ColumnType_STRING
	case "ipaddr":
		return proto.ColumnType_IPADDR
	case "cidr":
		return proto.ColumnType_CIDR
	case "timestamp":
		return proto.ColumnType_TIMESTAMP
	case "ltree":
		return proto.ColumnType_LTREE
	default:
		panic(fmt.Sprintf("invalid type %s", c.Type))
	}
}

type configTable struct {
	Name    string         `hcl:"name" cty:"name"`
	Columns []configColumn `hcl:"columns" cty:"columns"`
}

func (t configTable) AsPluginTable() *plugin.Table {
	res := &plugin.Table{Name: t.Name,
		List: &plugin.ListConfig{
			Hydrate: nullList,
		}}
	res.Columns = make([]*plugin.Column, len(t.Columns))
	for i, c := range t.Columns {
		res.Columns[i] = &plugin.Column{
			Name: c.Name,
			Type: c.getType(),
		}
	}
	return res
}

func nullList(ctx context.Context, data *plugin.QueryData, data2 *plugin.HydrateData) (interface{}, error) {
	return nil, nil
}

var ConfigSchema = map[string]*schema.Attribute{
	"tables": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{
			Type: schema.TypeMap,
			AttrTypes: map[string]*schema.Attribute{
				"name": {
					Type: schema.TypeString,
				},
				"columns": {
					Type: schema.TypeList,
					Elem: &schema.Attribute{

						Type: schema.TypeMap,
						AttrTypes: map[string]*schema.Attribute{
							"name": {
								Type: schema.TypeString,
							},
							"type": {
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	},
}

func ConfigInstance() interface{} {
	return &DyntestConfig{}
}

func getConfig(connection *plugin.Connection) DyntestConfig {
	if connection == nil || connection.Config == nil {
		return DyntestConfig{}
	}
	config, ok := connection.Config.(DyntestConfig)
	if !ok {
		return DyntestConfig{}
	}
	return config
}
