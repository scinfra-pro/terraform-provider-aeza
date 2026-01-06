---
page_title: "aeza_service_prolong Resource - terraform-provider-aeza"
subcategory: ""
description: |-
  Extend Aeza service subscription period.
---

# aeza_service_prolong (Resource)

Extend the subscription period for an Aeza service.

## Example Usage

```terraform
resource "aeza_service_prolong" "extend_server" {
  service_id = aeza_service.web_server.id
  period     = "month"
}
```

## Schema

### Required

- `service_id` (String) The ID of the service to extend.
- `period` (String) Subscription period (hour, month, quarter_year, half_year, year).

### Read-Only

- `id` (String) The ID of this resource.
- `expires_at` (String) New expiration date after prolongation.

