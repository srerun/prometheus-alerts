---
title: ElasticsearchRelocatingShardsTooLong
description: Troubleshooting for alert ElasticsearchRelocatingShardsTooLong
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchRelocatingShardsTooLong

Elasticsearch has been relocating shards for 15min

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchRelocatingShardsTooLong" %}}

{{% comment %}}

```yaml
alert: ElasticsearchRelocatingShardsTooLong
expr: elasticsearch_cluster_health_relocating_shards > 0
for: 15m
labels:
    severity: warning
annotations:
    summary: Elasticsearch relocating shards too long (instance {{ $labels.instance }})
    description: |-
        Elasticsearch has been relocating shards for 15min
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchRelocatingShardsTooLong.md

```

{{% /comment %}}

</details>


## Meaning

The "ElasticsearchRelocatingShardsTooLong" alert is triggered when Elasticsearch has been relocating shards for an extended period of time (15 minutes). This alert indicates that there is an issue with the Elasticsearch cluster that is preventing it from completing the shard relocation process in a timely manner.

## Impact

The impact of this alert can be significant, as it can lead to:

* Slow query performance
* Increased latency
* Data inconsistencies
* Potential data loss
* Cluster instability

If the shard relocation process is stuck, it can cause a cascading effect on the entire Elasticsearch cluster, leading to a degradation of service quality and potential downtime.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Elasticsearch cluster health using the Elasticsearch API or a monitoring tool like Kibana.
2. Verify that the cluster is not undergoing maintenance or experiencing high load.
3. Check the Elasticsearch logs for any errors or warnings related to shard relocation.
4. Verify that the disk space is sufficient and the nodes have enough resources (CPU, RAM, and disk space) to perform the shard relocation.
5. Check if there are any network connectivity issues between the nodes.
6. Verify that the Elasticsearch configuration is correct and the shard allocation is properly set up.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the Elasticsearch cluster health and identify the nodes that are stuck in the relocating shards state.
2. Restart the affected nodes or the entire cluster, if necessary, to reset the shard relocation process.
3. Verify that the disk space is sufficient and the nodes have enough resources (CPU, RAM, and disk space) to perform the shard relocation.
4. Check and adjust the Elasticsearch configuration to optimize the shard allocation and relocation process.
5. Implement a monitoring solution to detect and alert on shard relocation issues earlier.
6. Consider setting up a retry mechanism to automatically retry the shard relocation process if it fails.

Remember to investigate the root cause of the issue and fix it to prevent similar issues from occurring in the future.