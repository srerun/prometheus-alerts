---
title: ThanosRuleHighRuleEvaluationWarnings
description: Troubleshooting for alert ThanosRuleHighRuleEvaluationWarnings
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleHighRuleEvaluationWarnings

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Rule {{$labels.instance}} has high number of evaluation warnings.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ThanosRuleHighRuleEvaluationWarnings
expr: sum by (job, instance) (rate(thanos_rule_evaluation_with_warnings_total{job=~".*thanos-rule.*"}[5m])) > 0
for: 15m
labels:
    severity: info
annotations:
    summary: Thanos Rule High Rule Evaluation Warnings (instance {{ $labels.instance }})
    description: |-
        Thanos Rule {{$labels.instance}} has high number of evaluation warnings.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ThanosRuleHighRuleEvaluationWarnings

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
