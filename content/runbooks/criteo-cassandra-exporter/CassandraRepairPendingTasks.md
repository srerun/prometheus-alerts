---
title: CassandraRepairPendingTasks
description: Troubleshooting for alert CassandraRepairPendingTasks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraRepairPendingTasks

Some Cassandra repair tasks are pending

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraRepairPendingTasks" %}}

{{% comment %}}

```yaml
alert: CassandraRepairPendingTasks
expr: cassandra_stats{name="org:apache:cassandra:metrics:threadpools:internal:antientropystage:pendingtasks:value"} > 2
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra repair pending tasks (instance {{ $labels.instance }})
    description: |-
        Some Cassandra repair tasks are pending
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraRepairPendingTasks.md

```

{{% /comment %}}

</details>


Here is a runbook for the CassandraRepairPendingTasks alert:

## Meaning

The CassandraRepairPendingTasks alert is triggered when the number of pending repair tasks in Cassandra exceeds 2 for more than 2 minutes. This indicates that Cassandra is not able to keep up with the rate of repairs, which can lead to data inconsistencies and availability issues.

## Impact

* Data inconsistencies: Unrepaired data can lead to inconsistencies across the cluster, which can cause issues with data accuracy and integrity.
* Availability issues: If the pending tasks continue to accumulate, it can lead to node failures, causing unavailability of the cluster.
* Performance degradation: Excessive pending repair tasks can cause Cassandra to slow down, leading to performance degradation and increased latency.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cassandra logs for any errors or exceptions related to repair tasks.
2. Verify that the Cassandra nodes are properly configured and have sufficient resources (CPU, memory, disk space) to handle the repair tasks.
3. Check the Cassandra metrics for any signs of high latency, high CPU usage, or disk usage.
4. Verify that the cassandra-repair tool is correctly configured and running as expected.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Cassandra node to clear out any stuck repair tasks.
2. Verify that the cassandra-repair tool is correctly configured and running as expected.
3. Increase the resources (CPU, memory, disk space) allocated to the Cassandra nodes to handle the repair tasks.
4. Consider scaling out the Cassandra cluster to distribute the repair tasks across more nodes.
5. If the issue persists, consider running a manual repair using the cassandra-repair tool.

Remember to monitor the Cassandra metrics closely to ensure that the issue is resolved and does not recur.