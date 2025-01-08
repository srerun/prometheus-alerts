---
title: ElasticsearchHighIndexingLatency
description: Troubleshooting for alert ElasticsearchHighIndexingLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchHighIndexingLatency

The indexing latency on Elasticsearch cluster is higher than the threshold.

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchHighIndexingLatency" %}}

{{% comment %}}

```yaml
alert: ElasticsearchHighIndexingLatency
expr: elasticsearch_indices_indexing_index_time_seconds_total / elasticsearch_indices_indexing_index_total > 0.0005
for: 10m
labels:
    severity: warning
annotations:
    summary: Elasticsearch High Indexing Latency (instance {{ $labels.instance }})
    description: |-
        The indexing latency on Elasticsearch cluster is higher than the threshold.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchHighIndexingLatency.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
