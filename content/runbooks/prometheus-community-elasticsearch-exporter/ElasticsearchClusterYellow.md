---
title: ElasticsearchClusterYellow
description: Troubleshooting for alert ElasticsearchClusterYellow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchClusterYellow

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Elastic Cluster Yellow status

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchClusterYellow" %}}

{{% comment %}}

```yaml
alert: ElasticsearchClusterYellow
expr: elasticsearch_cluster_health_status{color="yellow"} == 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Elasticsearch Cluster Yellow (instance {{ $labels.instance }})
    description: |-
        Elastic Cluster Yellow status
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchClusterYellow.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
