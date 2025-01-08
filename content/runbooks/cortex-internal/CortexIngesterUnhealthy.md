---
title: CortexIngesterUnhealthy
description: Troubleshooting for alert CortexIngesterUnhealthy
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CortexIngesterUnhealthy

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Cortex has an unhealthy ingester

<details>
  <summary>Alert Rule</summary>

{{% rule "cortex/cortex-internal.yml" "CortexIngesterUnhealthy" %}}

<!-- Rule when generated

```yaml
alert: CortexIngesterUnhealthy
expr: cortex_ring_members{state="Unhealthy", name="ingester"} > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Cortex ingester unhealthy (instance {{ $labels.instance }})
    description: |-
        Cortex has an unhealthy ingester
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/cortex-internal/CortexIngesterUnhealthy.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
