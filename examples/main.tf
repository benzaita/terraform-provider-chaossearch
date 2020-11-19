terraform {
  required_providers {
    chaossearch = {
      versions = ["0.2"]
      source = "benzaita/chaossearch"
    }
  }
}

provider "chaossearch" {}

resource "chaossearch_object_group" "my-object-group" {
  name = "my-object-group"
  source_bucket = "<s3 bucket name>"
  live_events_sqs_arn ="<sqs arn>"

  filter_json = jsonencode({
    AND = [
      {
        field = "key"
        regex = ".*"
      }
    ]
  })

  compression = "gzip"
  format = "JSON"

  partition_by = "<regex>"
}
