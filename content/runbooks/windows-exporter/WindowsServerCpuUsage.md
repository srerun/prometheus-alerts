---
title: WindowsServerCpuUsage
description: Troubleshooting for alert WindowsServerCpuUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# WindowsServerCpuUsage

CPU Usage is more than 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "windows-server/windows-exporter.yml" "WindowsServerCpuUsage" %}}

{{% comment %}}

```yaml
alert: WindowsServerCpuUsage
expr: 100 - (avg by (instance) (rate(windows_cpu_time_total{mode="idle"}[2m])) * 100) > 80
for: 0m
labels:
    severity: warning
annotations:
    summary: Windows Server CPU Usage (instance {{ $labels.instance }})
    description: |-
        CPU Usage is more than 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/windows-exporter/WindowsServerCpuUsage.md

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
