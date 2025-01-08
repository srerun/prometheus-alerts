---
title: ThanosNoRuleEvaluations
description: Troubleshooting for alert ThanosNoRuleEvaluations
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosNoRuleEvaluations

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Rule {{$labels.instance}} did not perform any rule evaluations in the past 10 minutes.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-ruler.yml" "ThanosNoRuleEvaluations" %}}

{{% comment %}}

```yaml
alert: ThanosNoRuleEvaluations
expr: sum by (job, instance) (rate(prometheus_rule_evaluations_total{job=~".*thanos-rule.*"}[5m])) <= 0  and sum by (job, instance) (thanos_rule_loaded_rules{job=~".*thanos-rule.*"}) > 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos No Rule Evaluations (instance {{ $labels.instance }})
    description: |-
        Thanos Rule {{$labels.instance}} did not perform any rule evaluations in the past 10 minutes.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-ruler/ThanosNoRuleEvaluations.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
