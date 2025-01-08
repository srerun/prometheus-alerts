---
title: HostOutOfInodes
description: Troubleshooting for alert HostOutOfInodes
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostOutOfInodes

Disk is almost running out of available inodes (< 10% left)

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostOutOfInodes" %}}

{{% comment %}}

```yaml
alert: HostOutOfInodes
expr: (node_filesystem_files_free{fstype!="msdosfs"} / node_filesystem_files{fstype!="msdosfs"} * 100 < 10 and ON (instance, device, mountpoint) node_filesystem_readonly == 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host out of inodes (instance {{ $labels.instance }})
    description: |-
        Disk is almost running out of available inodes (< 10% left)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostOutOfInodes.md

```

{{% /comment %}}

</details>


Here is a runbook for the `HostOutOfInodes` alert rule:

## Meaning

The `HostOutOfInodes` alert is triggered when a host's file system is running low on available inodes. Inodes are data structures that represent files and directories on a file system. When a file system runs out of inodes, it can no longer create new files or directories, even if there is still available disk space.

## Impact

If left unaddressed, a host out of inodes can cause:

* Failure to write data to files or create new files and directories
* System crashes or hangs due to file system errors
* Disruption to applications and services that rely on writing to files or creating new files and directories

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected host and file system using the `instance` and `mountpoint` labels.
2. Use the `df -i` command to check the inode usage on the affected file system.
3. Verify that the file system is not read-only using the `mount` command.
4. Check for any disk space issues using the `df` command.
5. Review system logs for any errors or warnings related to file system errors.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and remove any unnecessary files or directories to free up inodes.
2. Consider increasing the inode count on the affected file system by resizing the file system or adding more disk space.
3. Verify that the file system is properly configured and that there are no issues with disk space allocation.
4. Consider implementing inode quotas or limits to prevent future inode exhaustion.
5. Monitor the host's file system inode usage regularly to prevent similar issues in the future.