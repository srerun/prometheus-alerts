---
title: ClickhouseDiskSpaceCriticalOnDefault
description: Troubleshooting for alert ClickhouseDiskSpaceCriticalOnDefault
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseDiskSpaceCriticalOnDefault

Disk space on default disk is critically low, below 10%.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseDiskSpaceCriticalOnDefault" %}}

{{% comment %}}

```yaml
alert: ClickhouseDiskSpaceCriticalOnDefault
expr: ClickHouseAsyncMetrics_DiskAvailable_default / (ClickHouseAsyncMetrics_DiskAvailable_default + ClickHouseAsyncMetrics_DiskUsed_default) * 100 < 10
for: 2m
labels:
    severity: critical
annotations:
    summary: ClickHouse Disk Space Critical on Default (instance {{ $labels.instance }})
    description: |-
        Disk space on default disk is critically low, below 10%.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseDiskSpaceCriticalOnDefault.md

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
