---
title: CephOsdLowSpace
description: Troubleshooting for alert CephOsdLowSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephOsdLowSpace

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Ceph Object Storage Daemon is going out of space. Please add more disks.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CephOsdLowSpace
expr: ceph_osd_utilization > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: Ceph OSD low space (instance {{ $labels.instance }})
    description: |-
        Ceph Object Storage Daemon is going out of space. Please add more disks.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/CephOsdLowSpace

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
