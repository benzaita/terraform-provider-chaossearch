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

### How to bump the version?

We follow the SemVer specification for our version string. Please read [here](https://semver.org) if you are not familiar with it.

To bump the version, simply bump the `VERSION` variable in the [Makefile](./Makefile). Note that this does not trigger a release -- see the "How to release" section for that.

### How to release?

Update the version in the Makefile.

Create a new tag of the following format: `vX.Y.Z` e.g. `v0.3.0`, and push it. Follow [this](https://github.com/benzaita/terraform-provider-chaossearch/actions?query=workflow%3Arelease) GitHub Workflow - it will push a new version to the Terraform Registry.
