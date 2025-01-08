---
title: ElasticsearchPendingTasks
description: Troubleshooting for alert ElasticsearchPendingTasks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchPendingTasks

Elasticsearch has pending tasks. Cluster works slowly.

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchPendingTasks" %}}

{{% comment %}}

```yaml
alert: ElasticsearchPendingTasks
expr: elasticsearch_cluster_health_number_of_pending_tasks > 0
for: 15m
labels:
    severity: warning
annotations:
    summary: Elasticsearch pending tasks (instance {{ $labels.instance }})
    description: |-
        Elasticsearch has pending tasks. Cluster works slowly.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchPendingTasks.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the ElasticsearchPendingTasks alert:

## Meaning

The ElasticsearchPendingTasks alert is triggered when the number of pending tasks in an Elasticsearch cluster exceeds 0. This alert indicates that the cluster is not processing tasks efficiently, which can lead to performance issues and slow query responses.

## Impact

The impact of this alert can be significant, as it can cause:

* Slow query responses
* Inefficient use of cluster resources
* Potential timeouts or errors in applications relying on the Elasticsearch cluster
* Decreased overall performance and availability of the cluster

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the Elasticsearch cluster health using the Elasticsearch API or a monitoring tool like Kibana.
2. Verify that the cluster is not overloaded or experiencing high latency.
3. Investigate the type of pending tasks (e.g., indexing, query, or administrative tasks).
4. Check the Elasticsearch logs for any errors or warnings related to task execution.
5. Review the cluster's resource utilization (e.g., CPU, memory, disk usage) to identify potential bottlenecks.

## Mitigation

To mitigate this alert, follow these steps:

1. Identify and cancel any unnecessary or stalled tasks.
2. Check for and resolve any cluster-level issues, such as node failures or network connectivity problems.
3. Optimize the cluster's configuration, such as adjusting the number of shards or replicas.
4. Consider upgrading the Elasticsearch cluster to a more performant version or adding more nodes to the cluster.
5. Implement task monitoring and alerting to detect pending tasks earlier and prevent them from accumulating.

Remember to investigate and address the root cause of the pending tasks to prevent future occurrences of this alert.