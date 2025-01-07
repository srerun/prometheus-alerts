---
title: HostUnusualDiskWriteRate
description: Troubleshooting for alert HostUnusualDiskWriteRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostUnusualDiskWriteRate

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Disk is probably writing too much data (> 50 MB/s)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HostUnusualDiskWriteRate
expr: (sum by (instance) (rate(node_disk_written_bytes_total[2m])) / 1024 / 1024 > 50) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host unusual disk write rate (instance {{ $labels.instance }})
    description: |-
        Disk is probably writing too much data (> 50 MB/s)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HostUnusualDiskWriteRate

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
