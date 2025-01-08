---
title: NetdataReportedUncorrectableDiskSectors
description: Troubleshooting for alert NetdataReportedUncorrectableDiskSectors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataReportedUncorrectableDiskSectors

Reported uncorrectable disk sectors

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "NetdataReportedUncorrectableDiskSectors" %}}

{{% comment %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/NetdataReportedUncorrectableDiskSectors.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule `NetdataReportedUncorrectableDiskSectors`:

## Meaning

This alert is triggered when Netdata reports an increase in uncorrectable disk sectors on a disk, indicating that the disk is experiencing errors that cannot be corrected by the disk's built-in error correction mechanisms. Uncorrectable disk sectors can lead to data loss and corruption, and if left unaddressed, can cause system instability and crashes.

## Impact

The impact of uncorrectable disk sectors can be significant, including:

* Data loss and corruption
* System instability and crashes
* Downtime and reduced availability
* Potential loss of business-critical data

## Diagnosis

To diagnose the root cause of the uncorrectable disk sectors, follow these steps:

1. Identify the affected disk and server
2. Check the disk's SMART (Self-Monitoring, Analysis and Reporting Technology) logs for errors and failures
3. Review system logs for disk-related errors and warnings
4. Run disk utility commands (e.g. `smartctl`, `badblocks`) to scan the disk for bad sectors and identify the extent of the problem
5. Check the disk's firmware and driver versions to ensure they are up-to-date

## Mitigation

To mitigate the effects of uncorrectable disk sectors, follow these steps:

1. Back up critical data immediately to prevent data loss
2. Replace the affected disk with a new one to prevent further data loss and corruption
3. Run disk utility commands (e.g. `fsck`, `e2fsck`) to repair file system errors and inconsistencies
4. Update the disk's firmware and driver versions to the latest available
5. Monitor the disk's health and performance closely to detect any further issues

Remember to consult the Netdata documentation and your organization's internal documentation for specific guidance on troubleshooting and resolving disk-related issues.