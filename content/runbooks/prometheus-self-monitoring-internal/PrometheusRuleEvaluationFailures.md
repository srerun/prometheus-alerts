---
title: PrometheusRuleEvaluationFailures
description: Troubleshooting for alert PrometheusRuleEvaluationFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusRuleEvaluationFailures

Prometheus encountered {{ $value }} rule evaluation failures, leading to potentially ignored alerts.

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusRuleEvaluationFailures" %}}

{{% comment %}}

```yaml
alert: PrometheusRuleEvaluationFailures
expr: increase(prometheus_rule_evaluation_failures_total[3m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus rule evaluation failures (instance {{ $labels.instance }})
    description: |-
        Prometheus encountered {{ $value }} rule evaluation failures, leading to potentially ignored alerts.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusRuleEvaluationFailures.md

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
