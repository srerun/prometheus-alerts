---
title: NetdataLowDiskSpace
description: Troubleshooting for alert NetdataLowDiskSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataLowDiskSpace

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Netdata low disk space (> 80%)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NetdataLowDiskSpace
expr: 100 / netdata_disk_space_GB_average * netdata_disk_space_GB_average{dimension=~"avail|cached"} < 20
for: 5m
labels:
    severity: warning
annotations:
    summary: Netdata low disk space (instance {{ $labels.instance }})
    description: |-
        Netdata low disk space (> 80%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/NetdataLowDiskSpace

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
