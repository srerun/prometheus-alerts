---
title: ElasticsearchHighQueryLatency
description: Troubleshooting for alert ElasticsearchHighQueryLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchHighQueryLatency

The query latency on Elasticsearch cluster is higher than the threshold.

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchHighQueryLatency" %}}

{{% comment %}}

```yaml
alert: ElasticsearchHighQueryLatency
expr: elasticsearch_indices_search_fetch_time_seconds / elasticsearch_indices_search_fetch_total > 1
for: 5m
labels:
    severity: warning
annotations:
    summary: Elasticsearch High Query Latency (instance {{ $labels.instance }})
    description: |-
        The query latency on Elasticsearch cluster is higher than the threshold.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchHighQueryLatency.md

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
