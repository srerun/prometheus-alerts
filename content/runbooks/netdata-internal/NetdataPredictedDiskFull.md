---
title: NetdataPredictedDiskFull
description: Troubleshooting for alert NetdataPredictedDiskFull
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataPredictedDiskFull

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Netdata predicted disk full in 24 hours

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NetdataPredictedDiskFull
expr: predict_linear(netdata_disk_space_GB_average{dimension=~"avail|cached"}[3h], 24 * 3600) < 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Netdata predicted disk full (instance {{ $labels.instance }})
    description: |-
        Netdata predicted disk full in 24 hours
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/NetdataPredictedDiskFull

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
