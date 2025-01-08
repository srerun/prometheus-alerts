---
title: PrometheusNotificationsBacklog
description: Troubleshooting for alert PrometheusNotificationsBacklog
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusNotificationsBacklog

The Prometheus notification queue has not been empty for 10 minutes

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusNotificationsBacklog" %}}

{{% comment %}}

```yaml
alert: PrometheusNotificationsBacklog
expr: min_over_time(prometheus_notifications_queue_length[10m]) > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus notifications backlog (instance {{ $labels.instance }})
    description: |-
        The Prometheus notification queue has not been empty for 10 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusNotificationsBacklog.md

```

{{% /comment %}}

</details>


Here is a runbook for the PrometheusNotificationsBacklog alert rule:

## Meaning

The PrometheusNotificationsBacklog alert is triggered when the Prometheus notification queue has not been empty for 10 minutes. This means that Prometheus is not able to process notifications (e.g. sending alerts to receivers) in a timely manner, which can lead to delayed or lost notifications.

## Impact

If the Prometheus notification queue is not processed in a timely manner, it can lead to:

* Delayed or lost notifications, which can result in delayed response to critical issues
* Increased latency in alerting and incident response
* Potential overload of the notification queue, leading to further delays or even crashes

## Diagnosis

To diagnose the issue, please follow these steps:

1. Check the Prometheus UI to see if there are any errors or warnings related to notification processing.
2. Check the Prometheus notification queue length to see how many notifications are pending.
3. Check the system resources (e.g. CPU, memory) to see if they are overwhelmed.
4. Check the notification receivers (e.g. Slack, PagerDuty) to see if they are experiencing any issues.

## Mitigation

To mitigate the issue, please follow these steps:

1. Restart the Prometheus server to clear the notification queue.
2. Check and adjust the notification queue settings to ensure they are adequate for the current load.
3. Investigate and resolve any issues with the notification receivers.
4. Consider increasing system resources (e.g. CPU, memory) to ensure they can handle the load.
5. Monitor the situation closely to ensure the queue is being processed correctly and no further delays occur.