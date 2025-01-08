---
title: CassandraCommitlogPendingTasks
description: Troubleshooting for alert CassandraCommitlogPendingTasks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraCommitlogPendingTasks

Unexpected number of Cassandra commitlog pending tasks

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraCommitlogPendingTasks" %}}

{{% comment %}}

```yaml
alert: CassandraCommitlogPendingTasks
expr: cassandra_stats{name="org:apache:cassandra:metrics:commitlog:pendingtasks:value"} > 15
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra commitlog pending tasks (instance {{ $labels.instance }})
    description: |-
        Unexpected number of Cassandra commitlog pending tasks
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraCommitlogPendingTasks.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `CassandraCommitlogPendingTasks`:

## Meaning

The `CassandraCommitlogPendingTasks` alert is triggered when the number of pending tasks in the Cassandra commit log exceeds 15 for a period of 2 minutes. This indicates that Cassandra is not able to process commit log tasks in a timely manner, which can lead to performance issues and data consistency problems.

## Impact

If left unchecked, a high number of pending commit log tasks can cause:

* Slow write performance
* Increased latency
* Risk of data loss or inconsistencies
* Reduced overall system reliability

## Diagnosis

To diagnose the root cause of this issue, follow these steps:

1. Check the Cassandra node's system logs for any errors or warnings related to the commit log.
2. Verify that the Cassandra node has sufficient resources (e.g., CPU, memory, disk space) to process the commit log.
3. Check the Cassandra cluster's overall health and topology to ensure that there are no issues with node connectivity or data replication.
4. Review the commit log's growth rate and adjust the commit log allocation accordingly.

## Mitigation

To mitigate the issue, take the following steps:

1. Investigate and address any underlying issues causing the commit log to grow excessively (e.g., high write workload, slow disk I/O).
2. Increase the commit log allocation to allow for more tasks to be processed concurrently.
3. Consider increasing the number of Cassandra nodes or upgrading node resources to improve overall system performance.
4. Monitor the commit log pending tasks metric to ensure the issue is resolved and does not reoccur.

Remember to refer to the Cassandra documentation and relevant tuning guides for more detailed information on commit log configuration and optimization.