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

Too much tombstones scanned in queries

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraTombstoneDump" %}}

{{% comment %}}

```yaml
alert: CassandraTombstoneDump
expr: cassandra_stats{name="org:apache:cassandra:metrics:table:tombstonescannedhistogram:99thpercentile"} > 1000
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra tombstone dump (instance {{ $labels.instance }})
    description: |-
        Too much tombstones scanned in queries
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraTombstoneDump.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule "CassandraTombstoneDump":

## Meaning

This alert is triggered when the 99th percentile of tombstones scanned in Cassandra queries exceeds 1000. Tombstones are a mechanism used by Cassandra to handle deletes, and excessive tombstone scanning can indicate inefficiencies in query planning or data modeling.

## Impact

* Increased latency and slower query performance due to excessive tombstone scanning
* Potential for query timeouts or failures
* Impact on overall Cassandra cluster performance and stability

## Diagnosis

1. Check the Cassandra server metrics to identify the specific table or tables affected by the high tombstone scanning.
2. Investigate recent changes to the data model or query patterns that may be contributing to the issue.
3. Review the Cassandra configuration to ensure that proper tuning and optimization have been applied.
4. Verify that there are no issues with data consistency or data corruption that could be contributing to the high tombstone count.

## Mitigation

1. Immediately investigate and address any data model or query pattern changes that may be contributing to the issue.
2. Optimize Cassandra configuration to improve query performance and reduce tombstone scanning (e.g., adjust the tombstone threshold, or enable parallel query execution).
3. Consider implementing data GC (garbage collection) or data compaction to remove tombstones and improve data efficiency.
4. Monitor Cassandra server metrics closely to ensure the issue is resolved and performance returns to normal.
5. If the issue persists, consider escalating to Cassandra experts or seeking additional support.