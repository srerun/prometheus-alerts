---
title: CassandraViewwriteLatency
description: Troubleshooting for alert CassandraViewwriteLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraViewwriteLatency

High viewwrite latency on {{ $labels.instance }} cassandra node

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraViewwriteLatency" %}}

{{% comment %}}

```yaml
alert: CassandraViewwriteLatency
expr: cassandra_stats{name="org:apache:cassandra:metrics:clientrequest:viewwrite:viewwritelatency:99thpercentile",service="cas"} > 100000
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra viewwrite latency (instance {{ $labels.instance }})
    description: |-
        High viewwrite latency on {{ $labels.instance }} cassandra node
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraViewwriteLatency.md

```

{{% /comment %}}

</details>


Here is the runbook for the CassandraViewwriteLatency alert rule:

## Meaning

The CassandraViewwriteLatency alert is triggered when the 99th percentile of view write latency exceeds 100,000 microseconds (100ms) for a Cassandra node. This indicates that the Cassandra node is experiencing high latency when writing data to views, which can impact the performance and availability of the Cassandra cluster.

## Impact

High view write latency can have several impacts on the Cassandra cluster and dependent applications:

* Increased latency for client requests: High view write latency can cause delays in processing client requests, leading to poor user experience and increased error rates.
* Decreased throughput: High latency can reduce the throughput of the Cassandra cluster, leading to increased latency and decreased performance.
* Increased resource utilization: High latency can lead to increased resource utilization, such as CPU and memory, which can exacerbate the issue and lead to node failures.

## Diagnosis

To diagnose the root cause of high view write latency, follow these steps:

1. Check the Cassandra node's resource utilization: Verify that the node has sufficient resources (CPU, memory, disk space) to handle the workload.
2. Investigate disk I/O performance: Disk I/O issues can contribute to high latency. Check disk usage, IOPS, and throughput to identify any bottlenecks.
3. Review Cassandra configuration: Verify that the Cassandra configuration is optimized for performance, including settings such as compaction, compression, and caching.
4. Analyze Cassandra logs: Review Cassandra logs to identify any errors or warnings that may indicate the root cause of the issue.
5. Check for network issues: Verify that the network connection between the Cassandra node and dependent applications is stable and performant.

## Mitigation

To mitigate high view write latency, follow these steps:

1. Tune Cassandra configuration: Optimize Cassandra configuration settings, such as compaction, compression, and caching, to improve performance.
2. Scale up Cassandra resources: Increase resources (CPU, memory, disk space) to improve the node's ability to handle the workload.
3. Implement caching: Implement caching mechanisms, such as Redis or Memcached, to reduce the load on the Cassandra node.
4. Improve disk I/O performance: Upgrade disk storage to improve IOPS and throughput, or consider using SSDs.
5. Consider load balancing: Consider load balancing across multiple Cassandra nodes to distribute the workload and improve performance.