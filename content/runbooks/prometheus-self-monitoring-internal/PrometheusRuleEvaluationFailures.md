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


Here is the runbook for the PrometheusRuleEvaluationFailures alert:

## Meaning

The PrometheusRuleEvaluationFailures alert is triggered when there is an increase in the total number of Prometheus rule evaluation failures within a 3-minute window. This means that Prometheus is experiencing issues evaluating rules, which can lead to missed or delayed alerts.

## Impact

The impact of this alert is critical, as it can result in:

* Missed alerts: Rules that are not evaluated may not trigger alerts, leading to potential issues going undetected.
* Delayed alerts: Even if rules are eventually evaluated, the delay can cause alerts to be triggered late, reducing the effectiveness of monitoring and incident response.
* Lack of visibility: Unevaluated rules can lead to a lack of visibility into the system's state, making it harder to identify and respond to issues.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus logs for errors related to rule evaluation.
2. Verify that the Prometheus instance is running correctly and not experiencing high CPU usage or memory issues.
3. Review the rules that are failing to evaluate and check for any syntax errors or conflicts with other rules.
4. Check the Prometheus configuration file for any issues or typos that could be causing rule evaluation failures.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Prometheus instance to clear any temporary issues.
2. Verify that the Prometheus configuration file is correct and up-to-date.
3. Review and correct any syntax errors or conflicts with other rules that are causing evaluation failures.
4. Consider increasing the resources (e.g., CPU, memory) allocated to the Prometheus instance to ensure it can handle the load.
5. Implement additional monitoring and alerting to detect rule evaluation failures earlier and reduce the impact.