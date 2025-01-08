---
title: ElasticsearchInitializingShards
description: Troubleshooting for alert ElasticsearchInitializingShards
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchInitializingShards

Elasticsearch is initializing shards

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchInitializingShards" %}}

{{% comment %}}

```yaml
alert: ElasticsearchInitializingShards
expr: elasticsearch_cluster_health_initializing_shards > 0
for: 0m
labels:
    severity: info
annotations:
    summary: Elasticsearch initializing shards (instance {{ $labels.instance }})
    description: |-
        Elasticsearch is initializing shards
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchInitializingShards.md

```

{{% /comment %}}

</details>


Here is a runbook for the ElasticsearchInitializingShards alert:

## Meaning

The ElasticsearchInitializingShards alert is triggered when the Prometheus metrics indicate that Elasticsearch is initializing shards. This means that the cluster is in the process of allocating shards to nodes, which can be a normal part of cluster operations, but can also indicate issues with the cluster's stability or performance.

## Impact

If left unchecked, initializing shards can lead to increased latency, reduced query performance, and potentially even cluster instability. This can have a cascading effect on dependent applications and services, leading to downtime and revenue loss.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Elasticsearch cluster health using the `_cluster/health` API endpoint to get an overview of the cluster's state.
2. Verify that the cluster is not experiencing any node failures or issues that could be preventing shards from being allocated correctly.
3. Check the Elasticsearch logs for any error messages related to shard initialization or allocation.
4. Use the Elasticsearch cat APIs (e.g. `_cat/shards`) to get a list of shards and their current state.
5. Verify that the cluster has sufficient resources (e.g. CPU, memory, disk space) to handle the current workload.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the Elasticsearch cluster configuration to ensure that it is properly sized for the current workload.
2. Verify that the cluster is properly configured for shard allocation, including setting the `cluster.routing.allocation.enable` property to `all` or `primaries`.
3. Consider increasing the `cluster.indices.recovery.max_bytes_per_sec` setting to allow for faster shard recovery.
4. If the issue persists, consider temporarily reducing the write load on the cluster or adding additional nodes to increase capacity.
5. Monitor the cluster's health and performance closely to ensure that the issue is fully resolved.

Remember to investigate the root cause of the issue to prevent it from happening again in the future.