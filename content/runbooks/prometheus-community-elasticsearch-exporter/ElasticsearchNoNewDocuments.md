---
title: ElasticsearchNoNewDocuments
description: Troubleshooting for alert ElasticsearchNoNewDocuments
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchNoNewDocuments

No new documents for 10 min!

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchNoNewDocuments" %}}

{{% comment %}}

```yaml
alert: ElasticsearchNoNewDocuments
expr: increase(elasticsearch_indices_indexing_index_total{es_data_node="true"}[10m]) < 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Elasticsearch no new documents (instance {{ $labels.instance }})
    description: |-
        No new documents for 10 min!
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchNoNewDocuments.md

```

{{% /comment %}}

</details>


Here is the runbook for the ElasticsearchNoNewDocuments alert:

## Meaning

The ElasticsearchNoNewDocuments alert indicates that no new documents have been indexed in Elasticsearch for the past 10 minutes. This alert is triggered when the rate of new documents being indexed is less than 1 per 10 minutes.

## Impact

The impact of this alert can be significant, as it may indicate that data is not being properly ingested into Elasticsearch. This can lead to a lack of visibility into system metrics, logs, and other data, making it difficult to monitor and troubleshoot issues. Additionally, if data is not being indexed, it may not be searchable, which can impact the performance of dependent applications and services.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Elasticsearch cluster health to ensure it is operational and not experiencing any issues.
2. Verify that the Elasticsearch data node is properly configured and running.
3. Check the Elasticsearch indexing pipeline to ensure it is functioning correctly.
4. Review the application logs to determine if there are any issues with data ingestion.
5. Check the network connectivity between the data sources and Elasticsearch to ensure it is not blocked or restricted.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve any issues with the Elasticsearch cluster, data node, or indexing pipeline.
2. Verify that data sources are properly configured and sending data to Elasticsearch.
3. Check for any issues with network connectivity between the data sources and Elasticsearch.
4. Consider increasing the timeout period for data ingestion or indexing to ensure that data is not lost during periods of high load or network congestion.
5. If the issue persists, consider escalating to an Elasticsearch administrator or expert for further assistance.