---
title: MinioDiskSpaceUsage
description: Troubleshooting for alert MinioDiskSpaceUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MinioDiskSpaceUsage

Minio available free space is low (< 10%)

<details>
  <summary>Alert Rule</summary>

{{% rule "minio/minio-internal.yml" "MinioDiskSpaceUsage" %}}

{{% comment %}}

```yaml
alert: MinioDiskSpaceUsage
expr: disk_storage_available / disk_storage_total * 100 < 10
for: 0m
labels:
    severity: warning
annotations:
    summary: Minio disk space usage (instance {{ $labels.instance }})
    description: |-
        Minio available free space is low (< 10%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/minio-internal/MinioDiskSpaceUsage.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
