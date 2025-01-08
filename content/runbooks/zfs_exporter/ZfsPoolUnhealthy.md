---
title: ZfsPoolUnhealthy
description: Troubleshooting for alert ZfsPoolUnhealthy
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZfsPoolUnhealthy

## Meaning
[//]: # "Short paragraph that explains what the alert means"
ZFS pool state is {{ $value }}. See comments for more information.

<details>
  <summary>Alert Rule</summary>

{{% rule "zfs/zfs_exporter.yml" "ZfsPoolUnhealthy" %}}

{{% comment %}}

```yaml
alert: ZfsPoolUnhealthy
expr: zfs_pool_health > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: ZFS pool unhealthy (instance {{ $labels.instance }})
    description: |-
        ZFS pool state is {{ $value }}. See comments for more information.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/zfs_exporter/ZfsPoolUnhealthy.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
