---
title: CortexNotificationAreBeingDropped
description: Troubleshooting for alert CortexNotificationAreBeingDropped
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CortexNotificationAreBeingDropped

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Cortex notification are being dropped due to errors (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "cortex/coretex-internal.yml" "CortexNotificationAreBeingDropped" %}}

{{% comment %}}

```yaml
alert: CortexNotificationAreBeingDropped
expr: rate(cortex_prometheus_notifications_dropped_total[5m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Cortex notification are being dropped (instance {{ $labels.instance }})
    description: |-
        Cortex notification are being dropped due to errors (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/coretex-internal/CortexNotificationAreBeingDropped.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
