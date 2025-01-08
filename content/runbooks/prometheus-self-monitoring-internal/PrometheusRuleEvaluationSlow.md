---
title: PrometheusRuleEvaluationSlow
description: Troubleshooting for alert PrometheusRuleEvaluationSlow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusRuleEvaluationSlow

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Prometheus rule evaluation took more time than the scheduled interval. It indicates a slower storage backend access or too complex query.

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusRuleEvaluationSlow" %}}

{{% comment %}}

```yaml
alert: PrometheusRuleEvaluationSlow
expr: prometheus_rule_group_last_duration_seconds > prometheus_rule_group_interval_seconds
for: 5m
labels:
    severity: warning
annotations:
    summary: Prometheus rule evaluation slow (instance {{ $labels.instance }})
    description: |-
        Prometheus rule evaluation took more time than the scheduled interval. It indicates a slower storage backend access or too complex query.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusRuleEvaluationSlow.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
