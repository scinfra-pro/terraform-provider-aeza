---
page_title: "aeza_os Data Source - terraform-provider-aeza"
subcategory: ""
description: |-
  Get available operating systems.
---

# aeza_os (Data Source)

Use this data source to retrieve available operating systems for services.

## Example Usage

```terraform
data "aeza_os" "all" {}

output "operating_systems" {
  value = data.aeza_os.all.operating_systems
}
```

## Schema

### Read-Only

- `id` (String) The ID of this data source.
- `operating_systems` (List of Object) List of available operating systems.

