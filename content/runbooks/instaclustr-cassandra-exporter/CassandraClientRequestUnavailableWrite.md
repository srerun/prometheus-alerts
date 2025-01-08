---
title: CassandraClientRequestUnavailableWrite
description: Troubleshooting for alert CassandraClientRequestUnavailableWrite
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraClientRequestUnavailableWrite

Some Cassandra client requests are unavailable to write - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/instaclustr-cassandra-exporter.yml" "CassandraClientRequestUnavailableWrite" %}}

{{% comment %}}

```yaml
alert: CassandraClientRequestUnavailableWrite
expr: changes(cassandra_client_request_unavailable_exceptions_total{operation="write"}[1m]) > 0
for: 2m
labels:
    severity: critical
annotations:
    summary: Cassandra client request unavailable write (instance {{ $labels.instance }})
    description: |-
        Some Cassandra client requests are unavailable to write - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/instaclustr-cassandra-exporter/CassandraClientRequestUnavailableWrite.md

```

{{% /comment %}}

</details>


## Meaning

The CassandraClientRequestUnavailableWrite alert is triggered when the number of Cassandra client request unavailable exceptions for write operations increases within a 1-minute window. This indicates that the Cassandra cluster is currently unable to handle write requests, which can lead to data loss, inconsistencies, or application errors.

## Impact

* Data may not be written to the Cassandra cluster, leading to data loss or inconsistencies.
* Applications that rely on Cassandra for writes may experience errors or failures, causing disruptions to business-critical processes.
* The inability to write data to Cassandra can have a cascading effect on dependent systems, leading to a broader impact on the overall system.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cassandra cluster status and node health using the Cassandra nodetool or the Instaclustr Cassandra console.
2. Investigate the Cassandra system logs for any errors or exceptions related to write operations.
3. Verify that the Cassandra cluster is properly configured and that there are no network connectivity issues.
4. Check the application logs for any errors or exceptions related to Cassandra write operations.
5. Review the Prometheus metrics for any anomalies or trends that may indicate the root cause of the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the Cassandra cluster status and node health, and restart any nodes that are not operational.
2. Investigate and resolve any network connectivity issues between the Cassandra cluster and the application.
3. Verify that the Cassandra cluster is properly configured, including checking the replication factor, consistency level, and write timeouts.
4. Implement retries or fallback mechanisms in the application to handle temporary Cassandra unavailability.
5. Consider increasing the Cassandra cluster capacity or adjusting the write workload to prevent similar issues in the future.

Note: For more detailed instructions and troubleshooting steps, refer to the [runbook](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/instaclustr-cassandra-exporter/CassandraClientRequestUnavailableWrite.md) provided in the alert annotation.