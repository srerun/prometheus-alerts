---
title: PrometheusAlertmanagerJobMissing
description: Troubleshooting for alert PrometheusAlertmanagerJobMissing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusAlertmanagerJobMissing

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A Prometheus AlertManager job has disappeared

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PrometheusAlertmanagerJobMissing
expr: absent(up{job="alertmanager"})
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus AlertManager job missing (instance {{ $labels.instance }})
    description: |-
        A Prometheus AlertManager job has disappeared
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PrometheusAlertmanagerJobMissing

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
