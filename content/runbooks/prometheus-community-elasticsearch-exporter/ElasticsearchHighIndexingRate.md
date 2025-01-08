---
title: ElasticsearchHighIndexingRate
description: Troubleshooting for alert ElasticsearchHighIndexingRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchHighIndexingRate

The indexing rate on Elasticsearch cluster is higher than the threshold.

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchHighIndexingRate" %}}

{{% comment %}}

```yaml
alert: ElasticsearchHighIndexingRate
expr: sum(rate(elasticsearch_indices_indexing_index_total[1m]))> 10000
for: 5m
labels:
    severity: warning
annotations:
    summary: Elasticsearch High Indexing Rate (instance {{ $labels.instance }})
    description: |-
        The indexing rate on Elasticsearch cluster is higher than the threshold.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchHighIndexingRate.md

```

{{% /comment %}}

</details>


## Meaning

The ElasticsearchHighIndexingRate alert is triggered when the indexing rate on an Elasticsearch cluster exceeds 10,000 documents per minute for a sustained period of 5 minutes. This alert indicates that the cluster is experiencing high indexing load, which can lead to performance issues, increased latency, and potential data loss.

## Impact

* Performance degradation: High indexing rates can cause Elasticsearch nodes to slow down, leading to increased latency and decreased query performance.
* Resource utilization: High indexing rates can consume excessive CPU, memory, and disk resources, potentially causing node instability or even crashes.
* Data inconsistencies: High indexing rates can lead to indexing timeouts, causing data inconsistencies and potential data loss.

## Diagnosis

To diagnose the cause of the high indexing rate, follow these steps:

1. Check the Elasticsearch cluster's current indexing rate using the `elasticsearch_indices_indexing_index_total` metric.
2. Identify the specific indices or types of data that are causing the high indexing rate.
3. Investigate recent changes to the indexing pipeline, such as new data sources or changed indexing patterns.
4. Review Elasticsearch node logs for any errors or warnings related to indexing.
5. Check the cluster's resource utilization, such as CPU, memory, and disk usage.

## Mitigation

To mitigate the impact of high indexing rates, follow these steps:

1. **Temporary mitigation**: Reduce the indexing rate by pausing or slowing down the data ingestion pipeline.
2. **Index optimization**: Optimize index settings, such as increasing the `index.buffer.size` and `index.refresh.interval` to reduce the indexing load.
3. **Add more resources**: Scale up the Elasticsearch cluster by adding more nodes or increasing the resources allocated to existing nodes.
4. **Data pipeline optimization**: Optimize the data pipeline to reduce the volume of data being indexed or to distribute the indexing load more evenly.
5. **Long-term solution**: Implement a more sustainable indexing strategy, such as using a message queue or batch processing to reduce the real-time indexing load.