---
title: ZfsOfflinePool
description: Troubleshooting for alert ZfsOfflinePool
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZfsOfflinePool

A ZFS zpool is in a unexpected state: {{ $labels.state }}.

<details>
  <summary>Alert Rule</summary>

{{% rule "zfs/node-exporter.yml" "ZfsOfflinePool" %}}

{{% comment %}}

```yaml
alert: ZfsOfflinePool
expr: node_zfs_zpool_state{state!="online"} > 0
for: 1m
labels:
    severity: critical
annotations:
    summary: ZFS offline pool (instance {{ $labels.instance }})
    description: |-
        A ZFS zpool is in a unexpected state: {{ $labels.state }}.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/ZfsOfflinePool.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule "ZfsOfflinePool":

## Meaning

The ZfsOfflinePool alert is triggered when a ZFS zpool is in an unexpected state, meaning it's not online. This alert is critical because it can indicate a storage system failure, which can lead to data loss or unavailability.

## Impact

The impact of this alert is high, as it can cause:

* Data loss or corruption
* System downtime or unavailability
* Performance degradation
* Potential for cascading failures in dependent systems

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the ZFS zpool status using the `zpool status` command.
2. Verify the zpool configuration and ensure it's correct.
3. Check the system logs for any errors or warnings related to ZFS or the zpool.
4. Run `zpool scrub` to check for any data corruption or inconsistencies.
5. Review the node exporter metrics to identify any trends or patterns leading up to the alert.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately investigate and resolve any underlying system issues causing the zpool to be offline.
2. If the zpool is offline due to a faulty disk, replace the disk and resilver the zpool using `zpool replace` and `zpool resilver`.
3. If the zpool is offline due to a configuration issue, correct the configuration and bring the zpool online using `zpool online`.
4. Monitor the zpool status and node exporter metrics closely to ensure the issue is resolved and the system is stable.
5. Consider implementing additional monitoring and alerting for ZFS zpool health to detect potential issues before they become critical.