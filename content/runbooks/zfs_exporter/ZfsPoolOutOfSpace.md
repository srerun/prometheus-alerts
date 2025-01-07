---
title: ZfsPoolOutOfSpace
description: Troubleshooting for alert ZfsPoolOutOfSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZfsPoolOutOfSpace

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Disk is almost full (< 10% left)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ZfsPoolOutOfSpace
expr: zfs_pool_free_bytes * 100 / zfs_pool_size_bytes < 10 and ON (instance, device, mountpoint) zfs_pool_readonly == 0
for: 0m
labels:
    severity: warning
annotations:
    summary: ZFS pool out of space (instance {{ $labels.instance }})
    description: |-
        Disk is almost full (< 10% left)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ZfsPoolOutOfSpace

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
