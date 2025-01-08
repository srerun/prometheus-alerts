---
title: ZfsPoolOutOfSpace
description: Troubleshooting for alert ZfsPoolOutOfSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZfsPoolOutOfSpace

Disk is almost full (< 10% left)

<details>
  <summary>Alert Rule</summary>

{{% rule "zfs/zfs_exporter.yml" "ZfsPoolOutOfSpace" %}}

{{% comment %}}

```yaml
alert: ZfsPoolOutOfSpace
expr: zfs_pool_free_bytes * 100 / zfs_pool_size_bytes < 10 and ON (instance, device, mountpoint) zfs_pool_readonly == 0
for: 0m
labels:
    severity: warning
annotations:
    summary: ZFS pool out of space (instance {{ $labels.instance }})
    description: |-
        Disk is almost full (< 10% left)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/zfs_exporter/ZfsPoolOutOfSpace.md

```

{{% /comment %}}

</details>


Here is the runbook for the ZfsPoolOutOfSpace alert rule:

## Meaning

The ZfsPoolOutOfSpace alert rule is triggered when a ZFS pool's free space falls below 10% and the pool is not read-only. This alert indicates that the disk is almost full and immediate action is required to prevent data loss or corruption.

## Impact

If this alert is not addressed promptly, the affected ZFS pool may become completely full, leading to:

* Data loss or corruption
* System crashes or instability
* Inability to write data to the affected pool
* Potential security risks due to inability to update the system or write logs

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the ZFS pool usage using the `zpool list` command or a similar tool.
2. Identify the pool that is running low on space.
3. Check the pool's current usage and available space using the `zfs list` command or a similar tool.
4. Verify that the pool is not read-only using the `zfs get readonly` command or a similar tool.
5. Check the system logs for any errors or warnings related to the affected pool.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately stop any unnecessary writes to the affected pool to prevent further space depletion.
2. Identify and delete unnecessary files or data on the affected pool to free up space.
3. Consider adding more storage capacity to the affected pool or migrating data to a different pool with more available space.
4. Verify that the pool is no longer read-only using the `zfs set readonly=off` command or a similar tool.
5. Monitor the pool's usage and available space to ensure the issue is resolved and does not recur.

Remember to update the runbook as needed based on your specific use case and environment.