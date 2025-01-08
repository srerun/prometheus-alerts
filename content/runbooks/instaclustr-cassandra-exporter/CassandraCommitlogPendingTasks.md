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

Cassandra commitlog pending tasks - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/instaclustr-cassandra-exporter.yml" "CassandraCommitlogPendingTasks" %}}

{{% comment %}}

```yaml
alert: CassandraCommitlogPendingTasks
expr: cassandra_commit_log_pending_tasks > 15
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra commitlog pending tasks (instance {{ $labels.instance }})
    description: |-
        Cassandra commitlog pending tasks - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/instaclustr-cassandra-exporter/CassandraCommitlogPendingTasks.md

```

{{% /comment %}}

</details>


Here is a runbook for the CassandraCommitlogPendingTasks alert:

## Meaning

The CassandraCommitlogPendingTasks alert is triggered when the number of pending tasks in the Cassandra commit log exceeds 15 for more than 2 minutes. This indicates that the Cassandra node is experiencing high latency or is overwhelmed with write requests, leading to a backlog of uncommitted data in the commit log.

## Impact

A high number of pending tasks in the commit log can cause:

* Increased latency for writes and reads
* Increased memory usage on the Cassandra node
* Risk of data loss or corruption if the node fails or is restarted
* Performance degradation of the entire Cassandra cluster

## Diagnosis

To diagnose the cause of the alert, follow these steps:

1. Check the Cassandra node's system metrics (e.g., CPU, memory, disk usage) to identify any resource bottlenecks.
2. Review the Cassandra logs to identify any errors or exceptions related to write operations.
3. Check the Cassandra cluster's configuration and topology to ensure that the node is properly configured and not overloaded.
4. Verify that the Cassandra node is receiving an unusual amount of write traffic.

## Mitigation

To mitigate the alert, follow these steps:

1. Reduce the write load on the Cassandra node by:
	* Load balancing write traffic across multiple nodes
	* Implementing rate limiting or queuing mechanisms for writes
	* Optimizing application code to reduce write frequency or size
2. Increase the resources available to the Cassandra node, such as:
	* Adding more CPU or memory resources
	* Upgrading the node's hardware or infrastructure
3. Implement data compression or compaction to reduce the size of the commit log
4. Consider implementing a more robust backup and recovery strategy to minimize data loss in case of node failure.

Remember to investigate the root cause of the issue and implement a long-term solution to prevent the alert from recurring.