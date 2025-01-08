---
title: ThanosRuleHighRuleEvaluationFailures
description: Troubleshooting for alert ThanosRuleHighRuleEvaluationFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleHighRuleEvaluationFailures

Thanos Rule {{$labels.instance}} is failing to evaluate rules.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-ruler.yml" "ThanosRuleHighRuleEvaluationFailures" %}}

{{% comment %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-ruler/ThanosRuleHighRuleEvaluationFailures.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule `ThanosRuleHighRuleEvaluationFailures`:

## Meaning

The `ThanosRuleHighRuleEvaluationFailures` alert is triggered when the rate of rule evaluation failures for a Thanos Rule instance exceeds 5% of the total rule evaluations over a 5-minute window. This indicates that the Thanos Rule instance is experiencing issues evaluating rules, which can lead to incomplete or inaccurate metrics.

## Impact

The impact of this alert can be significant, as incomplete or inaccurate metrics can lead to:

* Inaccurate alerting and monitoring
* Delayed or missed incident detection
* Incomplete or inaccurate reporting and analytics

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Rule instance logs for errors related to rule evaluation.
2. Verify that the Thanos Rule instance is properly configured and has access to all required resources (e.g., CPU, memory, and disk space).
3. Check the rule configuration for errors or invalid syntax.
4. Verify that the underlying storage backend (e.g., object storage or disk storage) is functioning correctly.
5. Check for any network connectivity issues between the Thanos Rule instance and dependent services.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Thanos Rule instance to clear any temporary errors.
2. Review and fix any errors in the rule configuration.
3. Verify that the Thanos Rule instance has sufficient resources (e.g., CPU, memory, and disk space) to handle the rule evaluation load.
4. Check and fix any issues with the underlying storage backend.
5. If the issue persists, consider scaling up the Thanos Rule instance or distributing the rule evaluation load across multiple instances.

Remember to update the rule configuration and Thanos Rule instance as needed to prevent similar issues from occurring in the future.