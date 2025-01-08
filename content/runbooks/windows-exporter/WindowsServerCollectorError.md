---
title: WindowsServerCollectorError
description: Troubleshooting for alert WindowsServerCollectorError
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# WindowsServerCollectorError

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Collector {{ $labels.collector }} was not successful

<details>
  <summary>Alert Rule</summary>

{{% rule "windows-server/windows-exporter.yml" "WindowsServerCollectorError" %}}

<!-- Rule when generated

```yaml
alert: WindowsServerCollectorError
expr: windows_exporter_collector_success == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Windows Server collector Error (instance {{ $labels.instance }})
    description: |-
        Collector {{ $labels.collector }} was not successful
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/windows-exporter/WindowsServerCollectorError.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
