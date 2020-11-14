terraform {
  required_providers {
    chaossearch = {
      versions = ["0.2"]
      source = "github.com/benzaita/chaossearch"
    }
  }
}

provider "chaossearch" {}

resource "chaossearch_object_group" "itamar_test" {
  name = "itamar-test"
  source_bucket = "foo"
  daily_interval = true
  compression = "GZIP"
}

output "itamar_test" {
  value = chaossearch_object_group.itamar_test
}

data "chaossearch_object_groups" "all" {}

output "all_object_groups" {
  value = data.chaossearch_object_groups.all.object_groups
}
