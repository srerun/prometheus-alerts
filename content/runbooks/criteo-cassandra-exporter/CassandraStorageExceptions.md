---
title: CassandraStorageExceptions
description: Troubleshooting for alert CassandraStorageExceptions
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraStorageExceptions

Something is going wrong with cassandra storage

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraStorageExceptions" %}}

{{% comment %}}

```yaml
alert: CassandraStorageExceptions
expr: changes(cassandra_stats{name="org:apache:cassandra:metrics:storage:exceptions:count"}[1m]) > 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra storage exceptions (instance {{ $labels.instance }})
    description: |-
        Something is going wrong with cassandra storage
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraStorageExceptions.md

```

{{% /comment %}}

</details>


Here is a runbook for the CassandraStorageExceptions alert rule:

## Meaning

The CassandraStorageExceptions alert is triggered when there are more than one storage exceptions in Cassandra within a 1-minute interval. This indicates that there are issues with Cassandra's storage layer, which can cause data inconsistencies, latency, and other performance problems.

## Impact

The impact of this alert is critical, as storage exceptions can lead to:

* Data loss or corruption
* Increased latency and decreased performance
* Unavailability of Cassandra services
* Cascading failures in dependent systems

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cassandra logs for errors related to storage exceptions.
2. Verify the disk usage and available disk space on the Cassandra nodes.
3. Check for any recent configuration changes or software updates that may have caused the issue.
4. Use Cassandra's built-in monitoring tools, such as `nodetool` or `cqlsh`, to check the node's health and performance.
5. Review the Cassandra metrics, such as storage load, read/write latency, and exception rates, to identify the root cause of the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately investigate and address any underlying hardware or software issues causing the storage exceptions.
2. Check for and apply any pending Cassandra software updates or patches.
3. Consider increasing the disk space or adjusting the storage configuration to prevent storage exceptions.
4. Implement or adjust Cassandra's built-in mechanisms, such as hinted handoff or read repair, to minimize the impact of storage exceptions.
5. Perform a rolling restart of the Cassandra nodes to ensure that all nodes are running with the latest configuration and software updates.

Remember to update this runbook according to your team's specific needs and procedures.