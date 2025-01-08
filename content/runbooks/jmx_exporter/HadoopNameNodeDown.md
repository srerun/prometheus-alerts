---
title: HadoopNameNodeDown
description: Troubleshooting for alert HadoopNameNodeDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopNameNodeDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The Hadoop NameNode service is unavailable.

<details>
  <summary>Alert Rule</summary>

{{% rule "hadoop/jmx_exporter.yml" "HadoopNameNodeDown" %}}

<!-- Rule when generated

```yaml
alert: HadoopNameNodeDown
expr: up{job="hadoop-namenode"} == 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Hadoop Name Node Down (instance {{ $labels.instance }})
    description: |-
        The Hadoop NameNode service is unavailable.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/jmx_exporter/HadoopNameNodeDown.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
