---
title: PrometheusAlertmanagerNotificationFailing
description: Troubleshooting for alert PrometheusAlertmanagerNotificationFailing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusAlertmanagerNotificationFailing

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Alertmanager is failing sending notifications

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusAlertmanagerNotificationFailing" %}}

{{% comment %}}

```yaml
alert: PrometheusAlertmanagerNotificationFailing
expr: rate(alertmanager_notifications_failed_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus AlertManager notification failing (instance {{ $labels.instance }})
    description: |-
        Alertmanager is failing sending notifications
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusAlertmanagerNotificationFailing.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
