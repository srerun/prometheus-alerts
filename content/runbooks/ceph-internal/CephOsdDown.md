---
title: CephOsdDown
description: Troubleshooting for alert CephOsdDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephOsdDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Ceph Object Storage Daemon Down

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CephOsdDown
expr: ceph_osd_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Ceph OSD Down (instance {{ $labels.instance }})
    description: |-
        Ceph Object Storage Daemon Down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CephOsdDown

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
