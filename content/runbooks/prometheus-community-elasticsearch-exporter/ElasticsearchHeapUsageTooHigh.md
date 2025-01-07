---
title: ElasticsearchHeapUsageTooHigh
description: Troubleshooting for alert ElasticsearchHeapUsageTooHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchHeapUsageTooHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The heap usage is over 90%

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ElasticsearchHeapUsageTooHigh
expr: (elasticsearch_jvm_memory_used_bytes{area="heap"} / elasticsearch_jvm_memory_max_bytes{area="heap"}) * 100 > 90
for: 2m
labels:
    severity: critical
annotations:
    summary: Elasticsearch Heap Usage Too High (instance {{ $labels.instance }})
    description: |-
        The heap usage is over 90%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ElasticsearchHeapUsageTooHigh

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
