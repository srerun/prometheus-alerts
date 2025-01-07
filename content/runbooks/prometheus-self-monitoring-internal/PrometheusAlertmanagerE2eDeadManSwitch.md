---
title: PrometheusAlertmanagerE2eDeadManSwitch
description: Troubleshooting for alert PrometheusAlertmanagerE2eDeadManSwitch
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusAlertmanagerE2eDeadManSwitch

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Prometheus DeadManSwitch is an always-firing alert. It's used as an end-to-end test of Prometheus through the Alertmanager.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PrometheusAlertmanagerE2eDeadManSwitch
expr: vector(1)
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus AlertManager E2E dead man switch (instance {{ $labels.instance }})
    description: |-
        Prometheus DeadManSwitch is an always-firing alert. It's used as an end-to-end test of Prometheus through the Alertmanager.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PrometheusAlertmanagerE2eDeadManSwitch

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
