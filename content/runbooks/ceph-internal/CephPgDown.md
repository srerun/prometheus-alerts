---
title: CephPgDown
description: Troubleshooting for alert CephPgDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephPgDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some Ceph placement groups are down. Please ensure that all the data are available.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CephPgDown
expr: ceph_pg_down > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Ceph PG down (instance {{ $labels.instance }})
    description: |-
        Some Ceph placement groups are down. Please ensure that all the data are available.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CephPgDown

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
