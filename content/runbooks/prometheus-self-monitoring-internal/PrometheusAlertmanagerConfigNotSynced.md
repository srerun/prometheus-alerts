---
title: PrometheusAlertmanagerConfigNotSynced
description: Troubleshooting for alert PrometheusAlertmanagerConfigNotSynced
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusAlertmanagerConfigNotSynced

Configurations of AlertManager cluster instances are out of sync

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusAlertmanagerConfigNotSynced" %}}

{{% comment %}}

```yaml
alert: PrometheusAlertmanagerConfigNotSynced
expr: count(count_values("config_hash", alertmanager_config_hash)) > 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus AlertManager config not synced (instance {{ $labels.instance }})
    description: |-
        Configurations of AlertManager cluster instances are out of sync
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusAlertmanagerConfigNotSynced.md

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
