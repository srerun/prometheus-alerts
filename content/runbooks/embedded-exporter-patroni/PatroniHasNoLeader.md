---
title: PatroniHasNoLeader
description: Troubleshooting for alert PatroniHasNoLeader
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PatroniHasNoLeader

A leader node (neither primary nor standby) cannot be found inside the cluster {{ $labels.scope }}

<details>
  <summary>Alert Rule</summary>

{{% rule "patroni/embedded-exporter-patroni.yml" "PatroniHasNoLeader" %}}

{{% comment %}}

```yaml
alert: PatroniHasNoLeader
expr: (max by (scope) (patroni_master) < 1) and (max by (scope) (patroni_standby_leader) < 1)
for: 0m
labels:
    severity: critical
annotations:
    summary: Patroni has no Leader (instance {{ $labels.instance }})
    description: |-
        A leader node (neither primary nor standby) cannot be found inside the cluster {{ $labels.scope }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-patroni/PatroniHasNoLeader.md

```

{{% /comment %}}

</details>


## Meaning
The `PatroniHasNoLeader` alert is triggered when a Patroni cluster does not have a leader node, neither primary nor standby. This means that there is no node in the cluster that is responsible for making decisions and ensuring the cluster's overall health. This is a critical situation that requires immediate attention to prevent data loss and ensure business continuity.

## Impact
The lack of a leader node in a Patroni cluster can have severe consequences, including:

* Data loss or inconsistencies due to the absence of a primary node
* Inability to make decisions and respond to changes in the cluster
* Increased risk of cluster instability and further failures
* Potential revenue loss and business impact due to service unavailability

## Diagnosis
To diagnose the root cause of the `PatroniHasNoLeader` alert, follow these steps:

1. Check the Patroni cluster status and inspect the current node roles using the `patroni_cluster` metric.
2. Verify that all nodes in the cluster are correctly configured and running.
3. Review the system logs for any errors or warnings related to node failures or network connectivity issues.
4. Check for any recent changes or maintenance activities that may have caused the leader node to fail or step down.

## Mitigation
To mitigate the `PatroniHasNoLeader` alert, follow these steps:

1. **Immediately investigate and resolve the root cause of the leader node failure**. This may involve restarting the failed node, replacing a faulty node, or resolving network connectivity issues.
2. **Promote a standby node to primary** using the `patroni_ctl` command-line tool to ensure the cluster has a functional leader node.
3. **Verify the cluster status and node roles** using the `patroni_cluster` metric to ensure the cluster is healthy and functional.
4. **Perform a thorough analysis of the incident** to identify the root cause and implement measures to prevent similar incidents in the future.