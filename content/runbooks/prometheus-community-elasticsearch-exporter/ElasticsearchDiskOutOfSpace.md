---
title: ElasticsearchDiskOutOfSpace
description: Troubleshooting for alert ElasticsearchDiskOutOfSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchDiskOutOfSpace

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The disk usage is over 90%

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchDiskOutOfSpace" %}}

<!-- Rule when generated

```yaml
alert: ElasticsearchDiskOutOfSpace
expr: elasticsearch_filesystem_data_available_bytes / elasticsearch_filesystem_data_size_bytes * 100 < 10
for: 0m
labels:
    severity: critical
annotations:
    summary: Elasticsearch disk out of space (instance {{ $labels.instance }})
    description: |-
        The disk usage is over 90%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchDiskOutOfSpace.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
