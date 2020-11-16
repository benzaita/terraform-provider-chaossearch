terraform {
  required_providers {
    chaossearch = {
      versions = ["0.2"]
      source = "github.com/benzaita/chaossearch"
    }
  }
}

provider "chaossearch" {}

# output "itamar_test" {
  # value = chaossearch_object_group.itamar_test
# }

# data "chaossearch_object_groups" "all" {}

# output "all_object_groups" {
  # value = data.chaossearch_object_groups.all.object_groups
# }
