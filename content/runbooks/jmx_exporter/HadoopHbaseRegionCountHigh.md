---
title: HadoopHbaseRegionCountHigh
description: Troubleshooting for alert HadoopHbaseRegionCountHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopHbaseRegionCountHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The HBase cluster has an unusually high number of regions.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HadoopHbaseRegionCountHigh
expr: hadoop_hbase_region_count > 5000
for: 15m
labels:
    severity: warning
annotations:
    summary: Hadoop HBase Region Count High (instance {{ $labels.instance }})
    description: |-
        The HBase cluster has an unusually high number of regions.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HadoopHbaseRegionCountHigh

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
