---
title: HadoopResourceManagerDown
description: Troubleshooting for alert HadoopResourceManagerDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopResourceManagerDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The Hadoop ResourceManager service is unavailable.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HadoopResourceManagerDown
expr: up{job="hadoop-resourcemanager"} == 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Hadoop Resource Manager Down (instance {{ $labels.instance }})
    description: |-
        The Hadoop ResourceManager service is unavailable.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HadoopResourceManagerDown

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
