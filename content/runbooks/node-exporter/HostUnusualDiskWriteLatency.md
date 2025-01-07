---
title: HostUnusualDiskWriteLatency
description: Troubleshooting for alert HostUnusualDiskWriteLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostUnusualDiskWriteLatency

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Disk latency is growing (write operations > 100ms)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HostUnusualDiskWriteLatency
expr: (rate(node_disk_write_time_seconds_total[1m]) / rate(node_disk_writes_completed_total[1m]) > 0.1 and rate(node_disk_writes_completed_total[1m]) > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host unusual disk write latency (instance {{ $labels.instance }})
    description: |-
        Disk latency is growing (write operations > 100ms)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HostUnusualDiskWriteLatency

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
