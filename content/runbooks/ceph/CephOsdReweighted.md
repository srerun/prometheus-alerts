---
title: CephOsdReweighted
description: Troubleshooting for alert CephOsdReweighted
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephOsdReweighted

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Ceph Object Storage Daemon takes too much time to resize.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CephOsdReweighted
expr: ceph_osd_weight < 1
for: 2m
labels:
    severity: warning
annotations:
    summary: Ceph OSD reweighted (instance {{ $labels.instance }})
    description: |-
        Ceph Object Storage Daemon takes too much time to resize.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/CephOsdReweighted

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
