connection "dyntest"{
  plugin = "dyntest"
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