---
title: WindowsServerMemoryUsage
description: Troubleshooting for alert WindowsServerMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# WindowsServerMemoryUsage

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Memory usage is more than 90%

<details>
  <summary>Alert Rule</summary>

{{% rule "windows-server/windows-exporter.yml" "WindowsServerMemoryUsage" %}}

{{% comment %}}

```yaml
alert: WindowsServerMemoryUsage
expr: 100 - ((windows_os_physical_memory_free_bytes / windows_cs_physical_memory_bytes) * 100) > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: Windows Server memory Usage (instance {{ $labels.instance }})
    description: |-
        Memory usage is more than 90%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/windows-exporter/WindowsServerMemoryUsage.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
