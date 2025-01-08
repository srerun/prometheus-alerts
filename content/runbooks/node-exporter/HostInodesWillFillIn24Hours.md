---
title: HostInodesWillFillIn24Hours
description: Troubleshooting for alert HostInodesWillFillIn24Hours
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostInodesWillFillIn24Hours

Filesystem is predicted to run out of inodes within the next 24 hours at current write rate

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostInodesWillFillIn24Hours" %}}

{{% comment %}}

```yaml
alert: HostInodesWillFillIn24Hours
expr: (node_filesystem_files_free{fstype!="msdosfs"} / node_filesystem_files{fstype!="msdosfs"} * 100 < 10 and predict_linear(node_filesystem_files_free{fstype!="msdosfs"}[1h], 24 * 3600) < 0 and ON (instance, device, mountpoint) node_filesystem_readonly{fstype!="msdosfs"} == 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host inodes will fill in 24 hours (instance {{ $labels.instance }})
    description: |-
        Filesystem is predicted to run out of inodes within the next 24 hours at current write rate
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostInodesWillFillIn24Hours.md

```

{{% /comment %}}

</details>


Here is a runbook for the `HostInodesWillFillIn24Hours` alert:

## Meaning

The `HostInodesWillFillIn24Hours` alert is triggered when a host's inodes are predicted to be depleted within the next 24 hours at the current write rate. This alert indicates that the host is at risk of running out of file system inodes, which can lead to severe consequences, including data loss and system crashes.

## Impact

If left unchecked, this issue can cause:

* Data loss due to the inability to write new files or modify existing ones
* System crashes or instability due to file system errors
* Disruption to critical services and applications that rely on the affected file system
* Increased maintenance and troubleshooting efforts to resolve the issue

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the file system usage and inode count on the affected host using commands like `df -i` and `df -h`.
2. Review system logs for file system-related errors or warnings.
3. Verify that the host is not experiencing high write rates or unusual file system activity.
4. Check for any recent changes to the file system, such as new applications or services that may be consuming inodes.
5. Verify that the host has sufficient disk space and that the file system is not full.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the source of the inode consumption and address the root cause (e.g., remove unnecessary files, adjust application settings).
2. Increase the file system's inode count by resizing the file system or adding more disk space.
3. Implement inode-related monitoring and alerting to catch similar issues earlier.
4. Consider implementing file system quotas or limits to prevent future inode exhaustion.
5. Document the root cause and solution in the incident management system to improve future incident response.

Remember to follow your organization's incident response and change management processes when implementing these mitigation steps. Additionally, ensure that the affected host is properly monitored and alerted to prevent similar issues in the future.