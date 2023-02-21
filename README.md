# Chaos Dynamic Plugin for Steampipe

Use SQL to query users, clients, connections, keys and more from Chaos Dynamic.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/chaos)**
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-chaos-dynamic/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install chaosdynamic
```

Setup connection config($INSTALL_DIR/.steampipe/config):

```hcl
connection "chaosdynamic"{
  plugin = "chaosdynamic"
  tables = [
    {
      name    = "test1"
      columns = [
        {
          name = "c1"
          type = "string"
        },
        {
          name = "c2"
          type = "string"
        },
      ]
    }
  ]
}
```

Run a query:

```sql
select
  *
from
  chaosdynamic.test1;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-chaos-dynamic.git
cd steampipe-plugin-chaos-dynamic
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/chaosdynamic.spc
```

Try it!

```
steampipe query
> .inspect chaosdynamic
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-chaos/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Chaos Dynamic Plugin](https://github.com/turbot/steampipe-plugin-chaos-dynamic/labels/help%20wanted)