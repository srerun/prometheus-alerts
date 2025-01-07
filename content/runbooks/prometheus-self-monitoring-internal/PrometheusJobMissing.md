---
title: PrometheusJobMissing
description: Troubleshooting for alert PrometheusJobMissing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusJobMissing

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A Prometheus job has disappeared

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PrometheusJobMissing
expr: absent(up{job="prometheus"})
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus job missing (instance {{ $labels.instance }})
    description: |-
        A Prometheus job has disappeared
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PrometheusJobMissing

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
