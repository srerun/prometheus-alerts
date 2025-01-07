---
title: HostUnusualDiskIo
description: Troubleshooting for alert HostUnusualDiskIo
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostUnusualDiskIo

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Time spent in IO is too high on {{ $labels.instance }}. Check storage for issues.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HostUnusualDiskIo
expr: (rate(node_disk_io_time_seconds_total[1m]) > 0.5) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 5m
labels:
    severity: warning
annotations:
    summary: Host unusual disk IO (instance {{ $labels.instance }})
    description: |-
        Time spent in IO is too high on {{ $labels.instance }}. Check storage for issues.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HostUnusualDiskIo

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
