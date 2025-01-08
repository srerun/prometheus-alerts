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


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
