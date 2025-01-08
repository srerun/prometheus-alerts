---
title: ElasticsearchHealthyNodes
description: Troubleshooting for alert ElasticsearchHealthyNodes
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchHealthyNodes

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Missing node in Elasticsearch cluster

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchHealthyNodes" %}}

<!-- Rule when generated

```yaml
alert: ElasticsearchHealthyNodes
expr: elasticsearch_cluster_health_number_of_nodes < 3
for: 0m
labels:
    severity: critical
annotations:
    summary: Elasticsearch Healthy Nodes (instance {{ $labels.instance }})
    description: |-
        Missing node in Elasticsearch cluster
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchHealthyNodes.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
