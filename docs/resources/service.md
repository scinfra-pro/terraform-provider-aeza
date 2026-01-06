---
page_title: "aeza_service Resource - terraform-provider-aeza"
subcategory: ""
description: |-
  Manages an Aeza service (VPS, dedicated server, etc.).
---

# aeza_service (Resource)

Manages an Aeza service including VPS, Hi-CPU servers, and dedicated servers.

## Example Usage

### Basic VPS

```terraform
resource "aeza_service" "web_server" {
  product_id = 181
  name       = "web-server"
  type       = "vps"
}
```

### VPS with Custom Configuration

```terraform
data "aeza_products" "vps" {}

resource "aeza_service" "app_server" {
  product_id = data.aeza_products.vps.products[0].id
  name       = "app-server"
  type       = "vps"
  
  # Optional: OS and recipe
  os_id     = 1
  recipe_id = 10
}
```

## Schema

### Required

- `product_id` (Number) The ID of the product to create. Use `aeza_products` data source to find available products.
- `name` (String) Name of the service.
- `type` (String) Type of service (vps, dedicated, hicpu, etc.). Use `aeza_service_types` data source for available types.

### Optional

- `os_id` (Number) Operating system ID. Use `aeza_os` data source for available operating systems.
- `recipe_id` (Number) Installation recipe ID. Use `aeza_recipes` data source for available recipes.

### Read-Only

- `id` (String) The ID of the service.
- `status` (String) Current status of the service.
- `ip_addresses` (List of String) IP addresses assigned to the service.
- `created_at` (String) Service creation timestamp.

## Import

Services can be imported using the service ID:

```bash
terraform import aeza_service.example 12345
```

