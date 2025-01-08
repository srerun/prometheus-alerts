---
title: HadoopDataNodeOutOfService
description: Troubleshooting for alert HadoopDataNodeOutOfService
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopDataNodeOutOfService

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The Hadoop DataNode is not sending heartbeats.

<details>
  <summary>Alert Rule</summary>

{{% rule "hadoop/jmx_exporter.yml" "HadoopDataNodeOutOfService" %}}

{{% comment %}}

```yaml
alert: HadoopDataNodeOutOfService
expr: hadoop_datanode_last_heartbeat == 0
for: 10m
labels:
    severity: warning
annotations:
    summary: Hadoop Data Node Out Of Service (instance {{ $labels.instance }})
    description: |-
        The Hadoop DataNode is not sending heartbeats.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/jmx_exporter/HadoopDataNodeOutOfService.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
