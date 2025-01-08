---
title: ThanosRuleHighRuleEvaluationWarnings
description: Troubleshooting for alert ThanosRuleHighRuleEvaluationWarnings
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleHighRuleEvaluationWarnings

Thanos Rule {{$labels.instance}} has high number of evaluation warnings.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-ruler.yml" "ThanosRuleHighRuleEvaluationWarnings" %}}

{{% comment %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-ruler/ThanosRuleHighRuleEvaluationWarnings.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the ThanosRuleHighRuleEvaluationWarnings alert:

## Meaning

This alert is triggered when Thanos Rule Evaluation has a high number of warnings for a certain instance. Thanos Rule Evaluation is responsible for evaluating Prometheus rules, and warnings indicate that there are issues with the rule evaluation process. This alert is raised when the number of warnings exceeds a certain threshold, indicating a potential problem that needs to be addressed.

## Impact

If this alert is not addressed, it can lead to issues with rule evaluation, which can result in incomplete or incorrect alerts being fired. This can impact the reliability and accuracy of the monitoring system, leading to missed incidents or false positives.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Rule Evaluation logs for errors and warnings related to the instance indicated in the alert.
2. Verify that the rule evaluation is working correctly and that there are no configuration issues.
3. Check the system resources (CPU, memory, disk space) to ensure they are within acceptable limits.
4. Review the rule definitions to ensure they are correct and up-to-date.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve any configuration issues or errors found in the Thanos Rule Evaluation logs.
2. Optimize system resources to ensure they are within acceptable limits.
3. Review and update rule definitions to ensure they are correct and up-to-date.
4. Consider increasing the logging level for Thanos Rule Evaluation to gather more information about the warnings.
5. If the issue persists, consider restarting the Thanos Rule Evaluation service or seeking assistance from a subject matter expert.