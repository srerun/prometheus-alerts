---
title: ElasticsearchUnassignedShards
description: Troubleshooting for alert ElasticsearchUnassignedShards
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchUnassignedShards

Elasticsearch has unassigned shards

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchUnassignedShards" %}}

{{% comment %}}

```yaml
alert: ElasticsearchUnassignedShards
expr: elasticsearch_cluster_health_unassigned_shards > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Elasticsearch unassigned shards (instance {{ $labels.instance }})
    description: |-
        Elasticsearch has unassigned shards
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchUnassignedShards.md

```

{{% /comment %}}

</details>


## Meaning

The ElasticsearchUnassignedShards alert is triggered when the Prometheus expression `elasticsearch_cluster_health_unassigned_shards > 0` is true, indicating that there are unassigned shards in the Elasticsearch cluster. This means that one or more shards in the cluster are not assigned to a node, which can lead to data unavailability and indexing issues.

## Impact

The impact of unassigned shards in Elasticsearch can be severe, leading to:

* Data unavailability: Unassigned shards mean that data is not available for search or query.
* Indexing issues: Unassigned shards can cause indexing requests to fail, leading to data loss and inconsistencies.
* Performance degradation: Unassigned shards can lead to increased latency and decreased performance in the cluster.
* Data inconsistency: Unassigned shards can cause data inconsistencies across the cluster, leading to incorrect or incomplete data.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Elasticsearch cluster health: Verify the overall health of the cluster using the Elasticsearch API or a monitoring tool like Kibana.
2. Identify the unassigned shards: Use the Elasticsearch API or a tool like Kibana to identify which shards are unassigned and which indices are affected.
3. Check node availability: Verify that all nodes in the cluster are available and running.
4. Check disk space: Ensure that all nodes in the cluster have sufficient disk space to accommodate the shards.
5. Check for errors: Check the Elasticsearch logs for any errors or exceptions related to the unassigned shards.

## Mitigation

To mitigate the issue, follow these steps:

1. Assign the unassigned shards: Use the Elasticsearch API to assign the unassigned shards to available nodes in the cluster.
2. Add more nodes: If the cluster is under-provisioned, consider adding more nodes to the cluster to increase capacity.
3. Reduce shard size: Consider reducing the shard size to make it easier to assign and manage.
4. Check for configuration issues: Verify that the Elasticsearch configuration is correct and that there are no configuration issues preventing shards from being assigned.
5. Consider a rolling restart: If all else fails, consider performing a rolling restart of the Elasticsearch cluster to ensure that all nodes are healthy and shards are properly assigned.