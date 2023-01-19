<p align="center">
    <h1 align="center">Chaos Plugin for Steampipe</h1>
</p>

<p align="center">
  <a aria-label="Steampipe logo" href="https://steampipe.io">
    <img src="https://steampipe.io/images/steampipe_logo_wordmark_padding.svg" height="28">
  </a>
  <a aria-label="Plugin version" href="https://hub.steampipe.io/plugins/turbot/dyntest">
    <img alt="" src="https://img.shields.io/static/v1?label=turbot/dyntest&message=v0.0.3&style=for-the-badge&labelColor=777777&color=F3F1F0">
  </a>
  &nbsp;
  <a aria-label="License" href="LICENSE">
    <img alt="" src="https://img.shields.io/static/v1?label=license&message=Apache-2.0&style=for-the-badge&labelColor=777777&color=F3F1F0">
  </a>
</p>

# Chaos Plugin for Steampipe

Use SQL to query all column types table, all numeric column type table and more to test your plugins.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/dyntest)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/
turbot/dyntest/tables)
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-dyntest/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install dyntest
```

Run a query:

```sql
select
  id,
  string_column,
  json_column
from
  dyntest_all_column_types
where
  id = '10';
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-dyntest.git
cd steampipe-plugin-dyntest
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/dyntest.spc
```

Try it!

```
steampipe query
> .inspect dyntest
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-dyntest/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Chaos Plugin](https://github.com/turbot/steampipe-plugin-dyntest/labels/help%20wanted)