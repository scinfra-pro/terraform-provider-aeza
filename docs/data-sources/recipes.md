---
page_title: "aeza_recipes Data Source - terraform-provider-aeza"
subcategory: ""
description: |-
  Get available installation recipes.
---

# aeza_recipes (Data Source)

Use this data source to retrieve available installation recipes (pre-configured software stacks).

## Example Usage

```terraform
data "aeza_recipes" "all" {}

output "recipes" {
  value = data.aeza_recipes.all.recipes
}
```

## Schema

### Read-Only

- `id` (String) The ID of this data source.
- `recipes` (List of Object) List of available installation recipes.

