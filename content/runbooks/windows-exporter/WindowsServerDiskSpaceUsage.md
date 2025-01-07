---
title: WindowsServerDiskSpaceUsage
description: Troubleshooting for alert WindowsServerDiskSpaceUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# WindowsServerDiskSpaceUsage

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Disk usage is more than 80%

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: WindowsServerDiskSpaceUsage
expr: 100.0 - 100 * ((windows_logical_disk_free_bytes / 1024 / 1024 ) / (windows_logical_disk_size_bytes / 1024 / 1024)) > 80
for: 2m
labels:
    severity: critical
annotations:
    summary: Windows Server disk Space Usage (instance {{ $labels.instance }})
    description: |-
        Disk usage is more than 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/WindowsServerDiskSpaceUsage

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
