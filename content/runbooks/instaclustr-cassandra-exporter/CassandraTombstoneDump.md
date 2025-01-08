---
title: CassandraTombstoneDump
description: Troubleshooting for alert CassandraTombstoneDump
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraTombstoneDump

Cassandra tombstone dump - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/instaclustr-cassandra-exporter.yml" "CassandraTombstoneDump" %}}

{{% comment %}}

```yaml
alert: CassandraTombstoneDump
expr: avg(cassandra_table_tombstones_scanned{quantile="0.99"}) by (instance,cassandra_cluster,keyspace) > 100
for: 2m
labels:
    severity: critical
annotations:
    summary: Cassandra tombstone dump (instance {{ $labels.instance }})
    description: |-
        Cassandra tombstone dump - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/instaclustr-cassandra-exporter/CassandraTombstoneDump.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the CassandraTombstoneDump alert:

## Meaning
The CassandraTombstoneDump alert is triggered when the 99th percentile of Cassandra table tombstones scanned exceeds 100 per instance, cluster, and keyspace. This indicates that Cassandra is experiencing high tombstone dumping, which can lead to performance issues and increased latency.

## Impact
High tombstone dumping can cause:

* Increased memory usage and garbage collection pauses
* Slower query performance and increased latency
* Higher disk usage due to bloom filter and index creation
* Potential for Cassandra node crashes or failures

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the Cassandra node logs for errors or warnings related to tombstone dumping.
2. Verify the Cassandra configuration to ensure that the tombstone threshold is set correctly.
3. Check the Cassandra metrics to identify the specific tables and keyspaces experiencing high tombstone dumping.
4. Run the `nodetool` command to inspect the Cassandra node's tombstone statistics.

## Mitigation
To mitigate the issue, follow these steps:

1. Increase the Cassandra node's heap size to accommodate increased memory usage.
2. Adjust the Cassandra configuration to optimize tombstone collection and removal.
3. Run the `nodetool compact` command to compact the Cassandra data files and remove tombstones.
4. Consider adding more Cassandra nodes to the cluster to distribute the load and reduce tombstone dumping.
5. Monitor the Cassandra metrics and adjust the configuration as needed to prevent future tombstone dumping issues.

Note: This is just a sample runbook, and the specific diagnosis and mitigation steps may vary depending on your Cassandra environment and configuration.