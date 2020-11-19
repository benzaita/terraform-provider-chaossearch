# Terraform Provider for ChaosSearch

This is a Terraform provider for [ChaosSearch](https://www.chaossearch.io/).

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
 - [ ] CI/CD pipeline (publish to the public Terraform Registry)
 - [ ] Deleting an OG that has indexes (should this even be allowed via Terraform?)