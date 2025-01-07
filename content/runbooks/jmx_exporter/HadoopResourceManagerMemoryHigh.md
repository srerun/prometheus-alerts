---
title: HadoopResourceManagerMemoryHigh
description: Troubleshooting for alert HadoopResourceManagerMemoryHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopResourceManagerMemoryHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The Hadoop ResourceManager is approaching its memory limit.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HadoopResourceManagerMemoryHigh
expr: hadoop_resourcemanager_memory_bytes / hadoop_resourcemanager_memory_max_bytes > 0.8
for: 15m
labels:
    severity: warning
annotations:
    summary: Hadoop Resource Manager Memory High (instance {{ $labels.instance }})
    description: |-
        The Hadoop ResourceManager is approaching its memory limit.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HadoopResourceManagerMemoryHigh

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
