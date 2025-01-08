---
title: ElasticsearchHeapUsageWarning
description: Troubleshooting for alert ElasticsearchHeapUsageWarning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchHeapUsageWarning

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The heap usage is over 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchHeapUsageWarning" %}}

{{% comment %}}

```yaml
alert: ElasticsearchHeapUsageWarning
expr: (elasticsearch_jvm_memory_used_bytes{area="heap"} / elasticsearch_jvm_memory_max_bytes{area="heap"}) * 100 > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: Elasticsearch Heap Usage warning (instance {{ $labels.instance }})
    description: |-
        The heap usage is over 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchHeapUsageWarning.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
