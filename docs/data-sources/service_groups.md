---
page_title: "aeza_service_groups Data Source - terraform-provider-aeza"
subcategory: ""
description: |-
  Get service groups.
---

# aeza_service_groups (Data Source)

Use this data source to retrieve service groups.

## Example Usage

```terraform
data "aeza_service_groups" "all" {}

output "groups" {
  value = data.aeza_service_groups.all.groups
}
```

## Schema

### Read-Only

- `id` (String) The ID of this data source.
- `groups` (List of Object) List of service groups.

