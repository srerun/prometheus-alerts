---
title: CortexNotificationError
description: Troubleshooting for alert CortexNotificationError
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CortexNotificationError

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Cortex is failing when sending alert notifications (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "cortex/cortex-internal.yml" "CortexNotificationError" %}}

{{% comment %}}

```yaml
alert: CortexNotificationError
expr: rate(cortex_prometheus_notifications_errors_total[5m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Cortex notification error (instance {{ $labels.instance }})
    description: |-
        Cortex is failing when sending alert notifications (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/cortex-internal/CortexNotificationError.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
