---
title: CephMonitorClockSkew
description: Troubleshooting for alert CephMonitorClockSkew
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephMonitorClockSkew

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Ceph monitor clock skew detected. Please check ntp and hardware clock settings

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CephMonitorClockSkew
expr: abs(ceph_monitor_clock_skew_seconds) > 0.2
for: 2m
labels:
    severity: warning
annotations:
    summary: Ceph monitor clock skew (instance {{ $labels.instance }})
    description: |-
        Ceph monitor clock skew detected. Please check ntp and hardware clock settings
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/CephMonitorClockSkew

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
