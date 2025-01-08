---
title: HostRaidDiskFailure
description: Troubleshooting for alert HostRaidDiskFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostRaidDiskFailure

At least one device in RAID array on {{ $labels.instance }} failed. Array {{ $labels.md_device }} needs attention and possibly a disk swap

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostRaidDiskFailure" %}}

{{% comment %}}

```yaml
alert: HostRaidDiskFailure
expr: (node_md_disks{state="failed"} > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host RAID disk failure (instance {{ $labels.instance }})
    description: |-
        At least one device in RAID array on {{ $labels.instance }} failed. Array {{ $labels.md_device }} needs attention and possibly a disk swap
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostRaidDiskFailure.md

```

{{% /comment %}}

</details>


## Meaning

The HostRaidDiskFailure alert is triggered when at least one device in a RAID array on a host fails. This alert is critical because it can lead to data loss or unavailability if not addressed promptly.

## Impact

The impact of a failed RAID disk can be significant, including:

* Data loss or corruption
* Reduced system performance
* Increased risk of system crashes or failures
* Potential for cascading failures of other disks in the RAID array
* Downtime and loss of productivity for users or services dependent on the host

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus alert details for the specific host and RAID array affected.
2. Log in to the host and check the system logs for errors related to the RAID array.
3. Use the `mdadm` command to check the status of the RAID array and identify the failed disk.
4. Verify that the failed disk is not causing any other system issues.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the failed disk and replace it as soon as possible.
2. Use the `mdadm` command to remove the failed disk from the RAID array.
3. Add a new disk to the RAID array and allow it to resync.
4. Monitor the RAID array for any further issues or errors.
5. Consider running a filesystem check on the affected host to ensure data integrity.
6. Update the monitoring system to reflect the changes made to the RAID array.

Note: Always follow proper safety protocols when replacing hardware components, and ensure that you have the necessary expertise and resources to perform the replacement.