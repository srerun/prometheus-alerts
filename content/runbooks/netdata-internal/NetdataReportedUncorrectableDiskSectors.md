---
title: NetdataReportedUncorrectableDiskSectors
description: Troubleshooting for alert NetdataReportedUncorrectableDiskSectors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataReportedUncorrectableDiskSectors

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Reported uncorrectable disk sectors

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NetdataReportedUncorrectableDiskSectors
expr: increase(netdata_smartd_log_offline_uncorrectable_sector_count_sectors_average[2m]) > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Netdata reported uncorrectable disk sectors (instance {{ $labels.instance }})
    description: |-
        Reported uncorrectable disk sectors
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/NetdataReportedUncorrectableDiskSectors

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
