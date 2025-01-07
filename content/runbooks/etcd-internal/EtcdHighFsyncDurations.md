---
title: EtcdHighFsyncDurations
description: Troubleshooting for alert EtcdHighFsyncDurations
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHighFsyncDurations

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Etcd WAL fsync duration increasing, 99th percentile is over 0.5s

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: EtcdHighFsyncDurations
expr: histogram_quantile(0.99, rate(etcd_disk_wal_fsync_duration_seconds_bucket[1m])) > 0.5
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd high fsync durations (instance {{ $labels.instance }})
    description: |-
        Etcd WAL fsync duration increasing, 99th percentile is over 0.5s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/EtcdHighFsyncDurations

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
