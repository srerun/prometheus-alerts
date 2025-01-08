---
title: PrometheusAlertmanagerConfigurationReloadFailure
description: Troubleshooting for alert PrometheusAlertmanagerConfigurationReloadFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusAlertmanagerConfigurationReloadFailure

AlertManager configuration reload error

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusAlertmanagerConfigurationReloadFailure" %}}

{{% comment %}}

```yaml
alert: PrometheusAlertmanagerConfigurationReloadFailure
expr: alertmanager_config_last_reload_successful != 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus AlertManager configuration reload failure (instance {{ $labels.instance }})
    description: |-
        AlertManager configuration reload error
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusAlertmanagerConfigurationReloadFailure.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
