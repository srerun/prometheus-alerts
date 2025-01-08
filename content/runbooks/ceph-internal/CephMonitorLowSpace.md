---
title: CephMonitorLowSpace
description: Troubleshooting for alert CephMonitorLowSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephMonitorLowSpace

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Ceph monitor storage is low.

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephMonitorLowSpace" %}}

<!-- Rule when generated

```yaml
alert: CephMonitorLowSpace
expr: ceph_monitor_avail_percent < 10
for: 2m
labels:
    severity: warning
annotations:
    summary: Ceph monitor low space (instance {{ $labels.instance }})
    description: |-
        Ceph monitor storage is low.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephMonitorLowSpace.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
