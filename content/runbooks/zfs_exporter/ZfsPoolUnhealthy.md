---
title: ZfsPoolUnhealthy
description: Troubleshooting for alert ZfsPoolUnhealthy
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZfsPoolUnhealthy

ZFS pool state is {{ $value }}. See comments for more information.

<details>
  <summary>Alert Rule</summary>

{{% rule "zfs/zfs_exporter.yml" "ZfsPoolUnhealthy" %}}

{{% comment %}}

```yaml
alert: ZfsPoolUnhealthy
expr: zfs_pool_health > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: ZFS pool unhealthy (instance {{ $labels.instance }})
    description: |-
        ZFS pool state is {{ $value }}. See comments for more information.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/zfs_exporter/ZfsPoolUnhealthy.md

```

{{% /comment %}}

</details>


## Meaning

The ZfsPoolUnhealthy alert is triggered when the `zfs_pool_health` metric reports a non-zero value, indicating that the ZFS pool is not healthy. This alert is critical, suggesting that immediate attention is required to prevent data loss or system downtime.

## Impact

A unhealthy ZFS pool can lead to:

* Data corruption or loss
* System instability or crashes
* Inaccessible data
* Performance degradation

If left unattended, an unhealthy ZFS pool can have a significant impact on system reliability, data integrity, and overall system availability.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the ZFS pool status using the `zpool status` command.
2. Identify the affected pool and its current health status.
3. Review the pool's configuration and layout to determine the cause of the issue.
4. Check the system logs for error messages related to the ZFS pool.
5. Verify that the ZFS pool is properly configured and that the underlying storage devices are functioning correctly.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately stop any writes to the affected pool to prevent further data corruption.
2. Attempt to repair the pool using the `zpool repair` command.
3. If the repair fails, try to resilver the pool using the `zpool resilver` command.
4. If the issue persists, consider replacing faulty storage devices or seeking assistance from a qualified system administrator.
5. Once the pool is healthy, ensure that regular backups are in place to prevent data loss in the future.

Note: This runbook is meant to serve as a general guideline and may need to be adapted to specific system configurations and requirements.