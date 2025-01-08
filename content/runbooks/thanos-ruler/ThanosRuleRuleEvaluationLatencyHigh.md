---
title: ThanosRuleRuleEvaluationLatencyHigh
description: Troubleshooting for alert ThanosRuleRuleEvaluationLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleRuleEvaluationLatencyHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Rule {{$labels.instance}} has higher evaluation latency than interval for {{$labels.rule_group}}.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-ruler.yml" "ThanosRuleRuleEvaluationLatencyHigh" %}}

{{% comment %}}

```yaml
alert: ThanosRuleRuleEvaluationLatencyHigh
expr: (sum by (job, instance, rule_group) (prometheus_rule_group_last_duration_seconds{job=~".*thanos-rule.*"}) > sum by (job, instance, rule_group) (prometheus_rule_group_interval_seconds{job=~".*thanos-rule.*"}))
for: 5m
labels:
    severity: warning
annotations:
    summary: Thanos Rule Rule Evaluation Latency High (instance {{ $labels.instance }})
    description: |-
        Thanos Rule {{$labels.instance}} has higher evaluation latency than interval for {{$labels.rule_group}}.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-ruler/ThanosRuleRuleEvaluationLatencyHigh.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
