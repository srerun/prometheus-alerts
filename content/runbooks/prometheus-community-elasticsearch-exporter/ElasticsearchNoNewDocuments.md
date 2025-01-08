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

## Meaning
[//]: # "Short paragraph that explains what the alert means"
No new documents for 10 min!

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchNoNewDocuments" %}}

<!-- Rule when generated

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

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
