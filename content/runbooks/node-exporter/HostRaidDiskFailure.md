---
title: HostRaidDiskFailure
description: Troubleshooting for alert HostRaidDiskFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostRaidDiskFailure

## Meaning
[//]: # "Short paragraph that explains what the alert means"
At least one device in RAID array on {{ $labels.instance }} failed. Array {{ $labels.md_device }} needs attention and possibly a disk swap

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HostRaidDiskFailure
expr: (node_md_disks{state="failed"} > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host RAID disk failure (instance {{ $labels.instance }})
    description: |-
        At least one device in RAID array on {{ $labels.instance }} failed. Array {{ $labels.md_device }} needs attention and possibly a disk swap
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HostRaidDiskFailure

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
