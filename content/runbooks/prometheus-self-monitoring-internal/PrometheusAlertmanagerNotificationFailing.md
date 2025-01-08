---
title: PrometheusAlertmanagerNotificationFailing
description: Troubleshooting for alert PrometheusAlertmanagerNotificationFailing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusAlertmanagerNotificationFailing

Alertmanager is failing sending notifications

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusAlertmanagerNotificationFailing" %}}

{{% comment %}}

```yaml
alert: PrometheusAlertmanagerNotificationFailing
expr: rate(alertmanager_notifications_failed_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus AlertManager notification failing (instance {{ $labels.instance }})
    description: |-
        Alertmanager is failing sending notifications
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusAlertmanagerNotificationFailing.md

```

{{% /comment %}}

</details>


## Meaning

The `PrometheusAlertmanagerNotificationFailing` alert is triggered when Alertmanager is experiencing issues sending notifications. This means that critical alerts may not be reaching the intended recipients, potentially leading to delayed or missed responses to incidents.

## Impact

The impact of this alert is high, as it can lead to:

* Delays in responding to critical incidents, potentially causing further damage or downtime
* Missed opportunities to address issues promptly, resulting in longer resolution times
* Inefficient use of resources, as teams may not be notified in a timely manner to address issues

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Alertmanager logs for errors or warnings related to notification sending
2. Verify that the notification configuration is correct and up-to-date
3. Check the status of the notification channels (e.g., email, PagerDuty, etc.) to ensure they are functioning correctly
4. Review the Alertmanager metrics to identify any trends or patterns that may indicate the cause of the issue

## Mitigation

To mitigate this issue, follow these steps:

1. Restart the Alertmanager service to reset the notification sending mechanism
2. Verify that the notification configuration is correct and up-to-date
3. Check the status of the notification channels and resolve any issues found
4. Implement a temporary workaround, such as manual notification or alternative notification channels, to ensure critical alerts are still being sent
5. Investigate and address any underlying issues that may be causing the notification failures, such as network connectivity problems or authentication issues.
6. Consider implementing redundancy or failover mechanisms for Alertmanager to minimize the impact of notification failures in the future.

Remember to refer to the [runbook](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusAlertmanagerNotificationFailing.md) for more detailed instructions and guidance on resolving this issue.