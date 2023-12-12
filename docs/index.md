---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/steampipe.svg"
brand_color: "#a42a2d"
display_name: "Chaos Dynamic"
name: "chaosdynamic"
description: "Steampipe plugin to test aggregation of dynamic plugin connections."
---

# Chaos Dynamic + Steampipe

Chaos Dynamic Plugin for testing aggregation of dynamic plugin connections with the craziest edge cases we can think of..

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

For example:

```sql
select
  *
from
  chaosdynamic.test1;
```

```
+------+------+------------------------------------+
| c1   | c2   | _ctx                               |
+------+------+------------------------------------+
| c1-1 | c2-1 | {"connection_name":"chaosdynamic"} |
| c1-0 | c2-0 | {"connection_name":"chaosdynamic"} |
+------+------+------------------------------------+
```

## Get started

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install chaosdynamic
```

Update the connection config ($INSTALL_DIR/.steampipe/config/chaosdynamic.spc) as per your requirement, by default the connection config file will look as follows:

```hcl
connection "chaosdynamic"{
  plugin = "chaosdynamic"
  tables = [
    {
      name    = "test1"
      description = "Test table 1"
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

)