---
title: HadoopMapReduceTaskFailures
description: Troubleshooting for alert HadoopMapReduceTaskFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopMapReduceTaskFailures

## Meaning
[//]: # "Short paragraph that explains what the alert means"
There is an unusually high number of MapReduce task failures.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HadoopMapReduceTaskFailures
expr: hadoop_mapreduce_task_failures_total > 100
for: 10m
labels:
    severity: critical
annotations:
    summary: Hadoop Map Reduce Task Failures (instance {{ $labels.instance }})
    description: |-
        There is an unusually high number of MapReduce task failures.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HadoopMapReduceTaskFailures

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
