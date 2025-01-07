---
title: CephHighOsdLatency
description: Troubleshooting for alert CephHighOsdLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephHighOsdLatency

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Ceph Object Storage Daemon latency is high. Please check if it doesn't stuck in weird state.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CephHighOsdLatency
expr: ceph_osd_perf_apply_latency_seconds > 5
for: 1m
labels:
    severity: warning
annotations:
    summary: Ceph high OSD latency (instance {{ $labels.instance }})
    description: |-
        Ceph Object Storage Daemon latency is high. Please check if it doesn't stuck in weird state.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CephHighOsdLatency

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
