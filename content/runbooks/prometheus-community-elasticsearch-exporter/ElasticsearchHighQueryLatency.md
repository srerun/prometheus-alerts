---
title: ElasticsearchHighQueryLatency
description: Troubleshooting for alert ElasticsearchHighQueryLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchHighQueryLatency

The query latency on Elasticsearch cluster is higher than the threshold.

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchHighQueryLatency" %}}

{{% comment %}}

```yaml
alert: ElasticsearchHighQueryLatency
expr: elasticsearch_indices_search_fetch_time_seconds / elasticsearch_indices_search_fetch_total > 1
for: 5m
labels:
    severity: warning
annotations:
    summary: Elasticsearch High Query Latency (instance {{ $labels.instance }})
    description: |-
        The query latency on Elasticsearch cluster is higher than the threshold.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchHighQueryLatency.md

```

{{% /comment %}}

</details>


Here is a runbook for the ElasticsearchHighQueryLatency alert:

## Meaning

The ElasticsearchHighQueryLatency alert is triggered when the average query latency on an Elasticsearch cluster exceeds 1 second. This is calculated by dividing the total search fetch time (in seconds) by the total number of search fetch operations.

## Impact

High query latency on Elasticsearch can have a significant impact on the performance and responsiveness of dependent applications and services. This can lead to:

* Slow query response times, affecting user experience
* Increased load on the Elasticsearch cluster, potentially leading to additional latency or even node failures
* Potential impact on business-critical operations, such as search, logging, and analytics

## Diagnosis

To diagnose the root cause of high query latency, follow these steps:

1. Check the Elasticsearch cluster health: Verify that the cluster is healthy and all nodes are available.
2. Investigate query patterns: Analyze query logs to identify any unusual or heavy query patterns that may be contributing to the latency.
3. Check indexing and shard configuration: Verify that indexing and shard configurations are optimal for the workload.
4. Review cluster resource utilization: Monitor CPU, memory, and disk usage to identify any resource bottlenecks.
5. Check for any ongoing maintenance or upgrades: Verify that there are no ongoing maintenance or upgrades that may be affecting query performance.

## Mitigation

To mitigate high query latency, follow these steps:

1. Optimize indexing and shard configuration: Review and adjust indexing and shard configurations to improve query performance.
2. Adjust query patterns: Optimize query patterns to reduce the load on the Elasticsearch cluster.
3. Add more resources: Consider adding more resources (e.g., nodes, CPU, memory) to the Elasticsearch cluster to improve query performance.
4. Implement query caching: Implement query caching to reduce the load on the Elasticsearch cluster.
5. Consider upgrading Elasticsearch: If the cluster is running an older version of Elasticsearch, consider upgrading to a newer version that may have performance improvements.

Remember to investigate and address the root cause of the high query latency to prevent similar issues in the future.