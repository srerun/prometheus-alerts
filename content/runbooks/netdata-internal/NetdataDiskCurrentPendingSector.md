---
title: NetdataDiskCurrentPendingSector
description: Troubleshooting for alert NetdataDiskCurrentPendingSector
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataDiskCurrentPendingSector

Disk current pending sector

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "NetdataDiskCurrentPendingSector" %}}

{{% comment %}}

```yaml
alert: NetdataDiskCurrentPendingSector
expr: netdata_smartd_log_current_pending_sector_count_sectors_average > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Netdata disk current pending sector (instance {{ $labels.instance }})
    description: |-
        Disk current pending sector
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/NetdataDiskCurrentPendingSector.md

```

{{% /comment %}}

</details>


## Meaning
This alert is triggered when the average count of current pending sectors on a disk is greater than 0, as reported by Netdata. A pending sector is a sector on a hard drive that is marked for replacement because it is likely to fail soon. This alert indicates that there is at least one disk with pending sectors that need attention.

## Impact
If left unaddressed, pending sectors can lead to data loss or corruption, and potentially even cause the disk to fail altogether. This can result in downtime, data loss, and potentially even affect the performance and reliability of the entire system.

## Diagnosis
To diagnose the issue, follow these steps:

1. Identify the disk with pending sectors by checking the `instance` label in the alert.
2. Check the disk's SMART (Self-Monitoring, Analysis and Reporting Technology) logs to confirm the presence of pending sectors.
3. Run a disk check and scan for bad blocks to identify the extent of the issue.
4. Verify that the disk is not already failing or showing signs of physical damage.

## Mitigation
To mitigate the issue, follow these steps:

1. Replace the disk with pending sectors as soon as possible to prevent data loss and corruption.
2. Run a disk check and scan for bad blocks to identify and repair any affected sectors.
3. Verify that the disk is properly configured and monitored for SMART errors.
4. Consider implementing disk redundancy and backup procedures to minimize the impact of disk failure.

Note: The runbook provides further guidance and instructions for resolving this issue. Please refer to the link provided in the alert annotations for more information.