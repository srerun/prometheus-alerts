---
title: ThanosRuleHighRuleEvaluationFailures
description: Troubleshooting for alert ThanosRuleHighRuleEvaluationFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleHighRuleEvaluationFailures

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Rule {{$labels.instance}} is failing to evaluate rules.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ThanosRuleHighRuleEvaluationFailures
expr: (sum by (job, instance) (rate(prometheus_rule_evaluation_failures_total{job=~".*thanos-rule.*"}[5m])) / sum by (job, instance) (rate(prometheus_rule_evaluations_total{job=~".*thanos-rule.*"}[5m])) * 100 > 5)
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Rule High Rule Evaluation Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Rule {{$labels.instance}} is failing to evaluate rules.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ThanosRuleHighRuleEvaluationFailures

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
