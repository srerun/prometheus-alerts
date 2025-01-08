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


Here is a runbook for the ThanosNoRuleEvaluations alert:

## Meaning

The ThanosNoRuleEvaluations alert is triggered when a Thanos Rule instance has not performed any rule evaluations in the past 10 minutes, despite having loaded rules. This indicates that the rule evaluation pipeline is not functioning correctly, which can lead to missing alerts and notifications.

## Impact

The impact of this alert is high, as it can result in:

* Missed alerts and notifications, leading to potential service disruptions or security breaches
* Incomplete monitoring and observability, making it difficult to identify and troubleshoot issues
* Delayed or inadequate response to incidents, leading to increased downtime and revenue loss

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Rule instance logs for errors or warnings related to rule evaluation.
2. Verify that the Thanos Rule instance is correctly configured and has access to the necessary resources (e.g., memory, CPU).
3. Check the Prometheus metrics for the Thanos Rule instance to ensure that the `prometheus_rule_evaluations_total` metric is being reported correctly.
4. Verify that the loaded rules are correct and up-to-date.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Thanos Rule instance to ensure that any stuck or hung processes are terminated.
2. Check and update the Thanos Rule configuration to ensure that it is correct and up-to-date.
3. Verify that the necessary resources (e.g., memory, CPU) are available to the Thanos Rule instance.
4. Check for any network connectivity issues that may be preventing the Thanos Rule instance from communicating with Prometheus.
5. If the issue persists, consider increasing the logging level or enabling debug logs to gather more detailed information about the issue.

Note: This runbook provides general guidance for diagnosing and mitigating the ThanosNoRuleEvaluations alert. The specific steps may vary depending on your environment and setup.