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
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
