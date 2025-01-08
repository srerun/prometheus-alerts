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

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The indexing rate on Elasticsearch cluster is higher than the threshold.

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchHighIndexingRate" %}}

<!-- Rule when generated

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

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
