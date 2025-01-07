---
title: HadoopHbaseWriteRequestsLatencyHigh
description: Troubleshooting for alert HadoopHbaseWriteRequestsLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopHbaseWriteRequestsLatencyHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
HBase Write Requests are experiencing high latency.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HadoopHbaseWriteRequestsLatencyHigh
expr: hadoop_hbase_write_requests_latency_seconds > 0.5
for: 10m
labels:
    severity: warning
annotations:
    summary: Hadoop HBase Write Requests Latency High (instance {{ $labels.instance }})
    description: |-
        HBase Write Requests are experiencing high latency.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HadoopHbaseWriteRequestsLatencyHigh

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
