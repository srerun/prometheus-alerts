---
title: ThanosRuleNoEvaluationFor10Intervals
description: Troubleshooting for alert ThanosRuleNoEvaluationFor10Intervals
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleNoEvaluationFor10Intervals

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Rule {{$labels.job}} has rule groups that did not evaluate for at least 10x of their expected interval.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ThanosRuleNoEvaluationFor10Intervals
expr: time() -  max by (job, instance, group) (prometheus_rule_group_last_evaluation_timestamp_seconds{job=~".*thanos-rule.*"})>10 * max by (job, instance, group) (prometheus_rule_group_interval_seconds{job=~".*thanos-rule.*"})
for: 5m
labels:
    severity: info
annotations:
    summary: Thanos Rule No Evaluation For10 Intervals (instance {{ $labels.instance }})
    description: |-
        Thanos Rule {{$labels.job}} has rule groups that did not evaluate for at least 10x of their expected interval.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ThanosRuleNoEvaluationFor10Intervals

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
