
variable "url" {
  type = string
  default = "mysql://root:pass@localhost:3306/"
}
variable "dev_url" {
  type = string
  default = "docker://mysql/8/dev"
}

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./models",
    "--dialect", "mysql", // | postgres | sqlite | sqlserver
  ]
}

env "local" {
  src = data.external_schema.gorm.url
  url = var.url
  dev = var.dev_url

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "dev" {
  src = data.external_schema.gorm.url
  url = var.url
  dev = var.dev_url

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}