---
title: NetdataDiskReallocatedSectors
description: Troubleshooting for alert NetdataDiskReallocatedSectors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataDiskReallocatedSectors

Reallocated sectors on disk

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "NetdataDiskReallocatedSectors" %}}

{{% comment %}}

```yaml
alert: NetdataDiskReallocatedSectors
expr: increase(netdata_smartd_log_reallocated_sectors_count_sectors_average[1m]) > 0
for: 0m
labels:
    severity: info
annotations:
    summary: Netdata disk reallocated sectors (instance {{ $labels.instance }})
    description: |-
        Reallocated sectors on disk
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/NetdataDiskReallocatedSectors.md

```

{{% /comment %}}

</details>


Here is a runbook for the NetdataDiskReallocatedSectors alert:

## Meaning

The NetdataDiskReallocatedSectors alert is triggered when the rate of reallocated sectors on a disk increases over a 1-minute window. This alert indicates that the disk is experiencing issues and is reallocating sectors, which can be a sign of impending disk failure.

## Impact

If left unaddressed, this issue can lead to:

* Data corruption or loss
* Disk failure, resulting in system downtime
* Performance degradation, leading to slower system response times

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the disk's SMART attributes using tools such as `smartctl` or `netdata` to determine the extent of the disk's health issues.
2. Review system logs for any errors or warnings related to the disk.
3. Verify that the disk is properly configured and mounted correctly.
4. Run a disk self-test to identify any bad blocks or other issues.

## Mitigation

To mitigate this issue, follow these steps:

1. Back up critical data to a separate storage device to prevent data loss in case of disk failure.
2. Replace the faulty disk with a new one as soon as possible.
3. Consider implementing redundant storage configurations, such as RAID, to minimize the impact of disk failure.
4. Monitor the disk's health regularly to catch any potential issues before they become critical.

Remember to follow proper procedures for replacing the disk and updating the system configuration to ensure minimal disruption to system operation.