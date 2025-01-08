---
title: ClickhouseDiskSpaceLowOnBackups
description: Troubleshooting for alert ClickhouseDiskSpaceLowOnBackups
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseDiskSpaceLowOnBackups

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Disk space on backups is below 20%.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseDiskSpaceLowOnBackups" %}}

<!-- Rule when generated

```yaml
alert: ClickhouseDiskSpaceLowOnBackups
expr: ClickHouseAsyncMetrics_DiskAvailable_backups / (ClickHouseAsyncMetrics_DiskAvailable_backups + ClickHouseAsyncMetrics_DiskUsed_backups) * 100 < 20
for: 2m
labels:
    severity: warning
annotations:
    summary: ClickHouse Disk Space Low on Backups (instance {{ $labels.instance }})
    description: |-
        Disk space on backups is below 20%.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseDiskSpaceLowOnBackups.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
