---
title: ElasticsearchRelocatingShards
description: Troubleshooting for alert ElasticsearchRelocatingShards
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchRelocatingShards

Elasticsearch is relocating shards

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchRelocatingShards" %}}

{{% comment %}}

```yaml
alert: ElasticsearchRelocatingShards
expr: elasticsearch_cluster_health_relocating_shards > 0
for: 0m
labels:
    severity: info
annotations:
    summary: Elasticsearch relocating shards (instance {{ $labels.instance }})
    description: |-
        Elasticsearch is relocating shards
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchRelocatingShards.md

```

{{% /comment %}}

</details>


Here is a runbook for the ElasticsearchRelocatingShards alert rule:

## Meaning

The ElasticsearchRelocatingShards alert is triggered when Elasticsearch is relocating shards. This means that Elasticsearch is currently moving shards from one node to another, which can cause temporary data unavailability and increased load on the cluster.

## Impact

The impact of this alert can be significant, as it can lead to:

* Temporary data unavailability: While shards are being relocated, data may be unavailable or partially available, which can affect application performance and user experience.
* Increased load on the cluster: The relocation process can cause additional load on the cluster, leading to increased CPU usage, memory usage, and network traffic.
* Potential for data loss: In rare cases, shard relocation can result in data loss or corruption if the process is interrupted or fails.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Elasticsearch cluster health: Verify that the cluster is in a healthy state and that there are no other issues affecting the cluster.
2. Identify the affected nodes: Check which nodes are involved in the shard relocation process and verify that they have sufficient resources (CPU, memory, disk space) to complete the process.
3. Check the shard relocation progress: Monitor the shard relocation progress using the Elasticsearch API or a monitoring tool like Kibana.
4. Review the cluster configuration: Check the cluster configuration to ensure that it is properly configured and optimized for shard relocation.

## Mitigation

To mitigate the issue, follow these steps:

1. Monitor the shard relocation progress: Closely monitor the shard relocation progress and take action if the process is taking longer than expected.
2. Ensure adequate resources: Verify that the nodes involved in the shard relocation process have sufficient resources (CPU, memory, disk space) to complete the process.
3. Optimize the cluster configuration: Review the cluster configuration and optimize it to improve the shard relocation process.
4. Consider rolling restarts: If the shard relocation process is taking too long, consider performing rolling restarts of the affected nodes to reduce the load on the cluster.

Remember to always follow best practices for troubleshooting and mitigation, and consult the Elasticsearch documentation and community resources if needed.