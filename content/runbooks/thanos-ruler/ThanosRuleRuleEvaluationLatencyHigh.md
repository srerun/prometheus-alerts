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


Here is a runbook for the Prometheus alert rule `ThanosRuleRuleEvaluationLatencyHigh`:

## Meaning

The `ThanosRuleRuleEvaluationLatencyHigh` alert is triggered when the evaluation latency of a Thanos rule group exceeds its interval. This means that the rule group is taking longer to evaluate than its intended interval, which can lead to delayed alerting and increased latency in the system.

## Impact

The impact of this alert is that Thanos rule evaluations may be delayed, leading to:

* Delayed alerting and notification
* Increased latency in the system
* Potential for missed alerts or notifications

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos rule group's configuration and ensure that the interval is set correctly.
2. Verify that the rule group's evaluation latency is not excessively high by checking the `prometheus_rule_group_last_duration_seconds` metric.
3. Investigate any recent changes to the rule group's configuration or the underlying system that may have caused the evaluation latency to increase.
4. Check the system's resource utilization (e.g., CPU, memory) to ensure that it is not overloaded.

## Mitigation

To mitigate the issue, follow these steps:

1. Adjust the rule group's interval to a higher value to give the system more time to evaluate the rules.
2. Optimize the rule group's configuration to reduce the evaluation latency (e.g., simplify rules, reduce the number of rules).
3. Scale up the underlying system to increase its capacity and reduce the evaluation latency.
4. Consider implementing a caching mechanism to reduce the load on the system and improve evaluation latency.