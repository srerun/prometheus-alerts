---
title: ClickhouseMemoryUsageCritical
description: Troubleshooting for alert ClickhouseMemoryUsageCritical
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseMemoryUsageCritical

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Memory usage is critically high, over 90%.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseMemoryUsageCritical" %}}

{{% comment %}}

```yaml
alert: ClickhouseMemoryUsageCritical
expr: ClickHouseAsyncMetrics_CGroupMemoryUsed / ClickHouseAsyncMetrics_CGroupMemoryTotal * 100 > 90
for: 5m
labels:
    severity: critical
annotations:
    summary: ClickHouse Memory Usage Critical (instance {{ $labels.instance }})
    description: |-
        Memory usage is critically high, over 90%.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseMemoryUsageCritical.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
