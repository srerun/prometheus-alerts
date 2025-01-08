---
title: ElasticsearchHeapUsageTooHigh
description: Troubleshooting for alert ElasticsearchHeapUsageTooHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchHeapUsageTooHigh

The heap usage is over 90%

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchHeapUsageTooHigh" %}}

{{% comment %}}

```yaml
alert: ElasticsearchHeapUsageTooHigh
expr: (elasticsearch_jvm_memory_used_bytes{area="heap"} / elasticsearch_jvm_memory_max_bytes{area="heap"}) * 100 > 90
for: 2m
labels:
    severity: critical
annotations:
    summary: Elasticsearch Heap Usage Too High (instance {{ $labels.instance }})
    description: |-
        The heap usage is over 90%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchHeapUsageTooHigh.md

```

{{% /comment %}}

</details>


Here is the runbook for the ElasticsearchHeapUsageTooHigh alert:

## Meaning

The ElasticsearchHeapUsageTooHigh alert is triggered when the heap usage of an Elasticsearch instance exceeds 90%. This is a critical alert, as high heap usage can lead to performance issues, slow query responses, and even node failures.

## Impact

The impact of high heap usage on Elasticsearch can be severe, leading to:

* Slow query responses and timeouts
* Node failures and crashes
* Increased latency and reduced throughput
* Potential data loss or corruption

## Diagnosis

To diagnose the root cause of high heap usage, follow these steps:

1. Check the Elasticsearch node's memory usage and garbage collection patterns.
2. Review the instance's configuration, including heap size, buffer sizes, and cache settings.
3. Investigate recent changes to the cluster, such as new indices, increased data ingestion, or changes to query patterns.
4. Check for any JVM-related issues, such as OutOfMemory errors or GC pauses.
5. Verify that the Elasticsearch instance is running with sufficient resources (e.g., CPU, memory, and disk space).

## Mitigation

To mitigate high heap usage, follow these steps:

1. **Increase the heap size**: Temporarily increase the heap size to give the node more breathing room.
2. **Optimize instance configuration**: Review and optimize the instance's configuration, including heap size, buffer sizes, and cache settings.
3. **Reduce data ingestion**: Reduce the rate of data ingestion or throttle heavy queries to alleviate pressure on the node.
4. **Investigate JVM-related issues**: Address any JVM-related issues, such as OutOfMemory errors or GC pauses, by adjusting JVM settings or upgrading the JVM version.
5. **Scale out the cluster**: Consider scaling out the cluster by adding more nodes or increasing the capacity of existing nodes.

Remember to monitor the Elasticsearch instance's heap usage and performance closely after implementing these mitigation steps to ensure the issue is fully resolved.