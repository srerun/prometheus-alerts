---
title: ThanosRuleQueueIsDroppingAlerts
description: Troubleshooting for alert ThanosRuleQueueIsDroppingAlerts
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleQueueIsDroppingAlerts

Thanos Rule {{$labels.instance}} is failing to queue alerts.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-ruler.yml" "ThanosRuleQueueIsDroppingAlerts" %}}

{{% comment %}}

```yaml
alert: ThanosRuleQueueIsDroppingAlerts
expr: sum by (job, instance) (rate(thanos_alert_queue_alerts_dropped_total{job=~".*thanos-rule.*"}[5m])) > 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Rule Queue Is Dropping Alerts (instance {{ $labels.instance }})
    description: |-
        Thanos Rule {{$labels.instance}} is failing to queue alerts.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-ruler/ThanosRuleQueueIsDroppingAlerts.md

```

{{% /comment %}}

</details>


## Meaning

The `ThanosRuleQueueIsDroppingAlerts` alert is triggered when Thanos Rule Queue is dropping alerts. This means that the Thanos Ruler instance is unable to queue alerts, which can lead to missing or delayed notifications for critical issues.

## Impact

The impact of this alert is high, as it can result in:

* Missing or delayed notifications for critical issues, potentially leading to undetected problems or service outages.
* Incomplete or inaccurate alerting, which can lead to incorrect incident responses or root cause analyses.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Ruler instance logs for errors or warnings related to alert queuing.
2. Investigate the Thanos Ruler configuration to ensure that the alert queue is properly configured.
3. Verify that the Thanos Ruler instance has sufficient resources (e.g., memory, CPU) to process alerts.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Thanos Ruler instance to attempt to recover from any temporary issues.
2. Investigate and resolve any configuration issues or resource constraints that may be contributing to the alert queuing failure.
3. Consider increasing the resources (e.g., memory, CPU) available to the Thanos Ruler instance to improve its ability to process alerts.
4. If the issue persists, consider escalating to the Thanos Ruler development team or seeking additional support from a qualified engineer.

Remember to refer to the [runbook](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-ruler/ThanosRuleQueueIsDroppingAlerts.md) for more detailed steps and troubleshooting guidance.