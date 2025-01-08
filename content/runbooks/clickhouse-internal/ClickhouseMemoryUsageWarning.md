---
title: ClickhouseMemoryUsageWarning
description: Troubleshooting for alert ClickhouseMemoryUsageWarning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseMemoryUsageWarning

Memory usage is over 80%.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseMemoryUsageWarning" %}}

{{% comment %}}

```yaml
alert: ClickhouseMemoryUsageWarning
expr: ClickHouseAsyncMetrics_CGroupMemoryUsed / ClickHouseAsyncMetrics_CGroupMemoryTotal * 100 > 80
for: 5m
labels:
    severity: warning
annotations:
    summary: ClickHouse Memory Usage Warning (instance {{ $labels.instance }})
    description: |-
        Memory usage is over 80%.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseMemoryUsageWarning.md

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
