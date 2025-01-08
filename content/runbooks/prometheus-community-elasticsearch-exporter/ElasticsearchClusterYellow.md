---
title: ElasticsearchClusterYellow
description: Troubleshooting for alert ElasticsearchClusterYellow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchClusterYellow

Elastic Cluster Yellow status

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchClusterYellow" %}}

{{% comment %}}

```yaml
alert: ElasticsearchClusterYellow
expr: elasticsearch_cluster_health_status{color="yellow"} == 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Elasticsearch Cluster Yellow (instance {{ $labels.instance }})
    description: |-
        Elastic Cluster Yellow status
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchClusterYellow.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the ElasticsearchClusterYellow alert:

## Meaning

The ElasticsearchClusterYellow alert is triggered when the Elasticsearch cluster health status turns yellow, indicating that one or more nodes in the cluster are not functioning correctly. This can have a significant impact on the cluster's performance and data integrity.

## Impact

* Data may not be available or up-to-date
* Searches and indexing may be slow or unresponsive
* Cluster performance may degrade
* Automatic failover may not be possible, leading to data loss or inconsistencies

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Elasticsearch cluster health status using the Elasticsearch API or a monitoring tool like Kibana.
2. Identify which node(s) are causing the cluster to turn yellow.
3. Check the node's logs for errors or warnings that may indicate the cause of the issue.
4. Verify that the node's resources (e.g., CPU, memory, disk space) are within acceptable limits.
5. Check for network connectivity issues between nodes.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and resolve the underlying cause of the issue (e.g., node failure, network issues, resource constraints).
2. Restart the affected node(s) if necessary.
3. Rebalance the cluster to redistribute data and indexing tasks.
4. Monitor the cluster's health status to ensure it returns to a green state.
5. Consider adding more nodes to the cluster to improve resilience and performance.

Additional resources:

* Elasticsearch documentation: [Cluster Health](https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html)
* Prometheus community runbook: [Elasticsearch Cluster Yellow](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchClusterYellow.md)