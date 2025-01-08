---
title: MongodbVirtualMemoryUsage
description: Troubleshooting for alert MongodbVirtualMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbVirtualMemoryUsage

## Meaning
[//]: # "Short paragraph that explains what the alert means"
High memory usage

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbVirtualMemoryUsage" %}}

{{% comment %}}

```yaml
alert: MongodbVirtualMemoryUsage
expr: (sum(mongodb_memory{type="virtual"}) BY (instance) / sum(mongodb_memory{type="mapped"}) BY (instance)) > 3
for: 2m
labels:
    severity: warning
annotations:
    summary: MongoDB virtual memory usage (instance {{ $labels.instance }})
    description: |-
        High memory usage
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbVirtualMemoryUsage.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
