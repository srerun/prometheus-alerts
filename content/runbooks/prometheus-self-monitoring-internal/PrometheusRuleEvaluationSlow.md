---
title: PrometheusRuleEvaluationSlow
description: Troubleshooting for alert PrometheusRuleEvaluationSlow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusRuleEvaluationSlow

Prometheus rule evaluation took more time than the scheduled interval. It indicates a slower storage backend access or too complex query.

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusRuleEvaluationSlow" %}}

{{% comment %}}

```yaml
alert: PrometheusRuleEvaluationSlow
expr: prometheus_rule_group_last_duration_seconds > prometheus_rule_group_interval_seconds
for: 5m
labels:
    severity: warning
annotations:
    summary: Prometheus rule evaluation slow (instance {{ $labels.instance }})
    description: |-
        Prometheus rule evaluation took more time than the scheduled interval. It indicates a slower storage backend access or too complex query.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusRuleEvaluationSlow.md

```

{{% /comment %}}

</details>


Here is a runbook for the PrometheusRuleEvaluationSlow alert:

## Meaning
The PrometheusRuleEvaluationSlow alert is triggered when the time it takes to evaluate Prometheus rules exceeds the scheduled interval. This can indicate a slower storage backend access or overly complex queries.

## Impact
If left unchecked, slow rule evaluation can lead to:

* Delays in alerting and notification delivery
* Increased load on the Prometheus server
* Potential loss of Prometheus data due to slow processing
* Inability to detect issues in a timely manner

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the Prometheus server logs for any errors or slowdowns related to rule evaluation.
2. Investigate the storage backend performance, such as disk I/O or database query times.
3. Review the complexity of the rules being evaluated, looking for any unnecessary or inefficient queries.
4. Check the `prometheus_rule_group_last_duration_seconds` and `prometheus_rule_group_interval_seconds` metrics to understand the scale of the issue.

## Mitigation
To mitigate the issue, follow these steps:

1. Optimize the storage backend configuration for better performance.
2. Simplify or optimize complex rules to reduce evaluation time.
3. Increase the `prometheus_rule_group_interval_seconds` to give the server more time to evaluate rules.
4. Consider scaling up the Prometheus server or distributing the load across multiple servers.
5. Implement a more efficient storage solution, such as(SSD) or a distributed database.