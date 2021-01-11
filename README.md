# Terraform Provider for ChaosSearch

This is a Terraform provider for [ChaosSearch](https://www.chaossearch.io/).

It is available in the Terraform Registry: [benzaita/chaossearch](https://registry.terraform.io/providers/benzaita/chaossearch/latest).

This is a work-in-progress.

TODO:

 - [x] Test with real credentials
 - [x] Parse XML response
 - [x] Provide authentication in the "provider" block in the Terraform manifest ([how to](https://learn.hashicorp.com/tutorials/terraform/provider-auth?in=terraform/providers))
 - [x] Implement create for resource_object_group
     - [x] Implement basic attributes (no filter)
     - [x] Implement filter in create request
 - [x] Implement update for resource_object_group
 - [x] Implement delete for resource_object_group
 - [x] Implement import for resource_object_group
 - [ ] Acceptance tests
 - [x] CI/CD pipeline (publish to the public Terraform Registry)
 - [ ] Deleting an OG that has indexes (should this even be allowed via Terraform?)
 
## Contributing

### How to test?

The `example` directory contains an example Terraform manifest. Make sure the provider version there matches the provider version in the `Makefile`, and then run:

```
make run
```

This will install the provider locally and run Terraform in the `example` directory.
