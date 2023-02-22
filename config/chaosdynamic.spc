connection "chaosdynamic"{
  plugin = "chaosdynamic"
  tables = [
    {
      name = "test1"
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