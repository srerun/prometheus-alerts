---
title: ElasticsearchDiskSpaceLow
description: Troubleshooting for alert ElasticsearchDiskSpaceLow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchDiskSpaceLow

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The disk usage is over 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchDiskSpaceLow" %}}

<!-- Rule when generated

```yaml
alert: ElasticsearchDiskSpaceLow
expr: elasticsearch_filesystem_data_available_bytes / elasticsearch_filesystem_data_size_bytes * 100 < 20
for: 2m
labels:
    severity: warning
annotations:
    summary: Elasticsearch disk space low (instance {{ $labels.instance }})
    description: |-
        The disk usage is over 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchDiskSpaceLow.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
