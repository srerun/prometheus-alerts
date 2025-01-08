---
title: NetdataPredictedDiskFull
description: Troubleshooting for alert NetdataPredictedDiskFull
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataPredictedDiskFull

Netdata predicted disk full in 24 hours

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "NetdataPredictedDiskFull" %}}

{{% comment %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/NetdataPredictedDiskFull.md

```

{{% /comment %}}

</details>


## Meaning

The NetdataPredictedDiskFull alert is triggered when the available disk space is predicted to reach 0 GB within the next 24 hours. This alert is based on the `netdata_disk_space_GB_average` metric, which measures the average available and cached disk space over a 3-hour period. The `predict_linear` function is used to forecast the disk space usage over the next 24 hours, and if the prediction indicates that the disk will be full, the alert is triggered.

## Impact

If this alert is triggered, it means that the disk is expected to run out of free space within the next 24 hours, which can cause significant disruptions to the system. This can lead to:

* Data loss or corruption
* System crashes or freezes
* Inability to write new data or perform critical tasks
* Downtime and lost productivity

## Diagnosis

To diagnose the root cause of the issue, follow these steps:

1. Check the disk usage graphs in Netdata to identify the rate of disk space consumption.
2. Investigate which applications or services are consuming the most disk space.
3. Check for any unnecessary files or data that can be safely removed.
4. Verify that the disk is not experiencing any hardware issues or errors.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and remove any unnecessary files or data that are consuming disk space.
2. Optimize applications or services to reduce their disk space usage.
3. Consider expanding the disk capacity or adding additional storage.
4. Implement disk space monitoring and alerting to detect potential issues earlier.
5. Develop a plan to regularly clean up and maintain the disk space to prevent future issues.

Note: This runbook is a general guide and may need to be adapted to the specific environment and requirements.