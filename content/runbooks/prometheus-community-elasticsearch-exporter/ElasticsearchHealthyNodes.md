---
title: ElasticsearchHealthyNodes
description: Troubleshooting for alert ElasticsearchHealthyNodes
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchHealthyNodes

Missing node in Elasticsearch cluster

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchHealthyNodes" %}}

{{% comment %}}

```yaml
alert: ElasticsearchHealthyNodes
expr: elasticsearch_cluster_health_number_of_nodes < 3
for: 0m
labels:
    severity: critical
annotations:
    summary: Elasticsearch Healthy Nodes (instance {{ $labels.instance }})
    description: |-
        Missing node in Elasticsearch cluster
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchHealthyNodes.md

```

{{% /comment %}}

</details>


## Meaning

The ElasticsearchHealthyNodes alert is triggered when the number of healthy nodes in an Elasticsearch cluster falls below 3. This indicates that the cluster is not in a healthy state and may be experiencing issues with data replication, indexing, or search functionality.

## Impact

The impact of this alert is critical, as it may lead to:

* Data loss or inconsistencies
* Search functionality degradation
* Increased latency or timeouts
* Potential cluster instability or even complete cluster failure

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Elasticsearch cluster health using the Elasticsearch API or a tool like Kibana.
2. Verify that the number of healthy nodes is indeed less than 3.
3. Investigate the nodes that are not healthy and check their logs for errors or issues.
4. Check the network connectivity between nodes and ensure that they can communicate with each other.
5. Review the Elasticsearch configuration and cluster settings to ensure they are correct and up-to-date.

## Mitigation

To mitigate this issue, follow these steps:

1. Identify the unhealthy nodes and investigate the root cause of their issues.
2. Bring the unhealthy nodes back online or replace them if necessary.
3. Ensure that the cluster is properly configured and that all nodes are properly connected.
4. Monitor the cluster health and node status to ensure that the issue does not reoccur.
5. Consider increasing the number of nodes in the cluster to improve resilience and redundancy.

For more detailed instructions and troubleshooting steps, refer to the runbook located at [https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchHealthyNodes.md](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchHealthyNodes.md).