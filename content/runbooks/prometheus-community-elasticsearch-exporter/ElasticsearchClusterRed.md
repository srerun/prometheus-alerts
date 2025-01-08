---
title: ElasticsearchClusterRed
description: Troubleshooting for alert ElasticsearchClusterRed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchClusterRed

Elastic Cluster Red status

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchClusterRed" %}}

{{% comment %}}

```yaml
alert: ElasticsearchClusterRed
expr: elasticsearch_cluster_health_status{color="red"} == 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Elasticsearch Cluster Red (instance {{ $labels.instance }})
    description: |-
        Elastic Cluster Red status
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchClusterRed.md

```

{{% /comment %}}

</details>


Here is a runbook for the ElasticsearchClusterRed alert rule:

## Meaning

The ElasticsearchClusterRed alert is triggered when the Elasticsearch cluster health status is reported as "red" by the Elasticsearch exporter. This indicates a critical problem with the Elasticsearch cluster, such as node failures, data loss, or shard allocation issues.

## Impact

A red Elasticsearch cluster health status can have severe consequences on the system, including:

* Data loss or inconsistency
* Search and indexing delays or failures
* Increased latency and errors
* Potential data corruption

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Elasticsearch cluster health API to confirm the red status.
2. Review the Elasticsearch logs for errors or warnings related to cluster health.
3. Verify that all nodes in the cluster are online and reachable.
4. Check for any ongoing indexing or search operations that may be causing the issue.
5. Investigate if there are any pending tasks, such as shard relocations or cluster updates, that may be contributing to the problem.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and resolve any node failures or connectivity issues.
2. Check for and address any data inconsistencies or corruption.
3. Run the Elasticsearch cluster health API to retrieve more detailed information about the cluster status.
4. Consider restarting the Elasticsearch cluster or individual nodes if necessary.
5. Implement any necessary configuration changes or patches to prevent similar issues in the future.
6. Monitor the Elasticsearch cluster health status closely to ensure the issue is resolved and does not recur.

Remember to refer to the Elasticsearch documentation and your organization's specific guidelines for additional troubleshooting and mitigation steps.