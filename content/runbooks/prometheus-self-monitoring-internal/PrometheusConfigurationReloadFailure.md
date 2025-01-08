---
title: PrometheusConfigurationReloadFailure
description: Troubleshooting for alert PrometheusConfigurationReloadFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusConfigurationReloadFailure

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Prometheus configuration reload error

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusConfigurationReloadFailure" %}}

{{% comment %}}

```yaml
alert: PrometheusConfigurationReloadFailure
expr: prometheus_config_last_reload_successful != 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus configuration reload failure (instance {{ $labels.instance }})
    description: |-
        Prometheus configuration reload error
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusConfigurationReloadFailure.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
