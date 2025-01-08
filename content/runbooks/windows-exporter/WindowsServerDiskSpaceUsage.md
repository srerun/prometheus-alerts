---
title: WindowsServerDiskSpaceUsage
description: Troubleshooting for alert WindowsServerDiskSpaceUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# WindowsServerDiskSpaceUsage

Disk usage is more than 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "windows-server/windows-exporter.yml" "WindowsServerDiskSpaceUsage" %}}

{{% comment %}}

```yaml
alert: WindowsServerDiskSpaceUsage
expr: 100.0 - 100 * ((windows_logical_disk_free_bytes / 1024 / 1024 ) / (windows_logical_disk_size_bytes / 1024 / 1024)) > 80
for: 2m
labels:
    severity: critical
annotations:
    summary: Windows Server disk Space Usage (instance {{ $labels.instance }})
    description: |-
        Disk usage is more than 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/windows-exporter/WindowsServerDiskSpaceUsage.md

```

{{% /comment %}}

</details>


Here is the runbook for the WindowsServerDiskSpaceUsage alert:

## Meaning

The WindowsServerDiskSpaceUsage alert indicates that the disk space usage on a Windows Server has exceeded 80%. This alert is critical, as it can lead to server crashes, slow performance, and data loss. It is essential to address this issue promptly to prevent downtime and ensure business continuity.

## Impact

High disk space usage can cause a range of issues, including:

* Server crashes or freezes
* Slow performance and responsiveness
* Inability to write data to disk, leading to data loss
* Failure of applications and services that rely on disk storage
* Increased risk of data corruption and security breaches

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the disk usage metrics to determine which disk(s) are experiencing high usage.
2. Identify the top disk space consumers using tools like Windows Explorer or disk usage analysis software.
3. Verify if there are any large files or directories that can be safely deleted or archived.
4. Check for any disk-related errors or warnings in the Windows Event Log.
5. Review disk usage trends to determine if the issue is sudden or gradual.

## Mitigation

To mitigate the issue, follow these steps:

1. **Delete unnecessary files and directories**: Remove any unnecessary files, logs, or data that are consuming disk space.
2. **Archive or move data**: Move infrequently used data to a secondary storage location, such as an external drive or cloud storage.
3. **Optimize disk usage**: Implement disk cleanup and optimization tools, such as Windows Disk Cleanup, to free up disk space.
4. **Expand disk capacity**: Consider expanding the disk capacity by adding more storage or replacing existing disks with larger ones.
5. **Implement disk monitoring**: Set up regular disk usage monitoring to detect potential issues before they become critical.

Remember to also investigate and address any underlying causes of high disk space usage, such as application or service misconfigurations, to prevent similar issues from occurring in the future.