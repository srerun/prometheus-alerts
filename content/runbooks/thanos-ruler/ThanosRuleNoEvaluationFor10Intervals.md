---
title: ThanosRuleNoEvaluationFor10Intervals
description: Troubleshooting for alert ThanosRuleNoEvaluationFor10Intervals
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleNoEvaluationFor10Intervals

Thanos Rule {{$labels.job}} has rule groups that did not evaluate for at least 10x of their expected interval.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-ruler.yml" "ThanosRuleNoEvaluationFor10Intervals" %}}

{{% comment %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-ruler/ThanosRuleNoEvaluationFor10Intervals.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule `ThanosRuleNoEvaluationFor10Intervals`:

## Meaning

The `ThanosRuleNoEvaluationFor10Intervals` alert is triggered when a Thanos rule group has not been evaluated for at least 10 times its expected interval. This means that the rule group has not been processed or executed for an extended period, which can lead to inconsistent monitoring data and potential issues with alerting and notification.

## Impact

The impact of this alert can be significant, as it may lead to:

* Inconsistent monitoring data, which can affect decision-making and incident response.
* Delayed or missing alerts and notifications, which can impact the timely detection and response to incidents.
* Inability to troubleshoot issues effectively, leading to prolonged downtime or service degradation.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos ruler logs for any errors or warnings related to rule group evaluation.
2. Verify that the rule group is correctly configured and enabled.
3. Check the Prometheus instance for any issues or overload conditions that may be preventing rule group evaluation.
4. Verify that the `prometheus_rule_group_last_evaluation_timestamp_seconds` and `prometheus_rule_group_interval_seconds` metrics are being correctly scraped and updated.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and address any underlying issues or errors in the Thanos ruler logs.
2. Verify that the rule group configuration is correct and up-to-date.
3. Check for any overload conditions or performance issues in the Prometheus instance and consider scaling or optimizing as needed.
4. Restart the Thanos ruler service or instance to trigger a re-evaluation of the rule groups.
5. Monitor the alert for resolution and verify that the rule groups are being evaluated correctly.

Additional resources:

* Thanos ruler documentation: [https://thanos.io/components/rule.md](https://thanos.io/components/rule.md)
* Prometheus alerting documentation: [https://prometheus.io/docs/alerting/latest/alerts](https://prometheus.io/docs/alerting/latest/alerts)