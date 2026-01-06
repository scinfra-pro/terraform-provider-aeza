---
page_title: "aeza_service_actions Resource - terraform-provider-aeza"
subcategory: ""
description: |-
  Perform actions on Aeza services.
---

# aeza_service_actions (Resource)

Perform actions on Aeza services such as start, stop, reboot, reinstall, etc.

## Example Usage

```terraform
resource "aeza_service_actions" "reboot_server" {
  service_id = aeza_service.web_server.id
  action     = "reboot"
}
```

## Schema

### Required

- `service_id` (String) The ID of the service to perform action on.
- `action` (String) Action to perform (start, stop, reboot, reinstall, etc.).

### Read-Only

- `id` (String) The ID of this resource.

