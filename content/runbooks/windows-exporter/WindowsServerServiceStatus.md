---
title: WindowsServerServiceStatus
description: Troubleshooting for alert WindowsServerServiceStatus
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# WindowsServerServiceStatus

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Windows Service state is not OK

<details>
  <summary>Alert Rule</summary>

{{% rule "windows-server/windows-exporter.yml" "WindowsServerServiceStatus" %}}

{{% comment %}}

```yaml
alert: WindowsServerServiceStatus
expr: windows_service_status{status="ok"} != 1
for: 1m
labels:
    severity: critical
annotations:
    summary: Windows Server service Status (instance {{ $labels.instance }})
    description: |-
        Windows Service state is not OK
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/windows-exporter/WindowsServerServiceStatus.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
