---
subcategory: "Tencent Container Registry(TCR)"
layout: "tencentcloud"
page_title: "TencentCloud: tencentcloud_tcr_webhook_trigger"
sidebar_current: "docs-tencentcloud-resource-tcr_webhook_trigger"
description: |-
  Provides a resource to create a tcr webhook_trigger
---

# tencentcloud_tcr_webhook_trigger

Provides a resource to create a tcr webhook_trigger

## Example Usage

```hcl
resource "tencentcloud_tcr_instance" "mytcr_webhooktrigger" {
  name          = "tf-test-tcr-%s"
  instance_type = "basic"
  delete_bucket = true

  tags = {
    test = "test"
  }
}

resource "tencentcloud_tcr_namespace" "my_ns" {
  instance_id    = tencentcloud_tcr_instance.mytcr_webhooktrigger.id
  name           = "tf_test_ns_%s"
  is_public      = true
  is_auto_scan   = true
  is_prevent_vul = true
  severity       = "medium"
  cve_whitelist_items {
    cve_id = "cve-xxxxx"
  }
}

data "tencentcloud_tcr_namespaces" "id_test" {
  instance_id = tencentcloud_tcr_namespace.my_ns.instance_id
}

locals {
  ns_id = data.tencentcloud_tcr_namespaces.id_test.namespace_list.0.id
}

resource "tencentcloud_tcr_webhook_trigger" "my_trigger" {
  registry_id = tencentcloud_tcr_instance.mytcr_webhooktrigger.id
  namespace   = tencentcloud_tcr_namespace.my_ns.name
  trigger {
    name = "trigger-%s"
    targets {
      address = "http://example.org/post"
      headers {
        key    = "X-Custom-Header"
        values = ["a"]
      }
    }
    event_types  = ["pushImage"]
    condition    = ".*"
    enabled      = true
    description  = "this is trigger description"
    namespace_id = local.ns_id

  }
  tags = {
    "createdBy" = "terraform"
  }
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required, String) namespace name.
* `registry_id` - (Required, String) instance Id.
* `trigger` - (Required, List) trigger parameters.
* `tags` - (Optional, Map) Tag description list.

The `headers` object supports the following:

* `key` - (Required, String) Header Key.
* `values` - (Required, Set) Header Values.

The `targets` object supports the following:

* `address` - (Required, String) target address.
* `headers` - (Optional, List) custom Headers.

The `trigger` object supports the following:

* `condition` - (Required, String) trigger rule.
* `enabled` - (Required, Bool) enable trigger.
* `event_types` - (Required, Set) trigger action.
* `name` - (Required, String) trigger name.
* `targets` - (Required, List) trigger target.
* `description` - (Optional, String) trigger description.
* `namespace_id` - (Optional, Int) the namespace Id to which the trigger belongs.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

tcr webhook_trigger can be imported using the id, e.g.

```
terraform import tencentcloud_tcr_webhook_trigger.webhook_trigger webhook_trigger_id
```

