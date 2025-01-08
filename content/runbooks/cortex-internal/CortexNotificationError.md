---
title: CortexNotificationError
description: Troubleshooting for alert CortexNotificationError
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CortexNotificationError

Cortex is failing when sending alert notifications (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "cortex/cortex-internal.yml" "CortexNotificationError" %}}

{{% comment %}}

```yaml
alert: CortexNotificationError
expr: rate(cortex_prometheus_notifications_errors_total[5m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Cortex notification error (instance {{ $labels.instance }})
    description: |-
        Cortex is failing when sending alert notifications (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/cortex-internal/CortexNotificationError.md

```

{{% /comment %}}

</details>


Here is a runbook for the CortexNotificationError alert rule:

## Meaning

This alert is triggered when Cortex, the Prometheus alert manager, encounters an error when sending alert notifications. This error prevents alerts from being delivered to the intended recipients, which can lead to delayed or missed responses to critical issues.

## Impact

The impact of this error is significant, as it can lead to:

* Delays in responding to critical issues, potentially causing further damage or downtime
* Missed opportunities to address issues before they escalate
* Reduced confidence in the overall monitoring and alerting system

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cortex logs for error messages related to notification sending
2. Verify that the notification recipients and channels are correctly configured
3. Check the network connectivity between Cortex and the notification servers (e.g. email, Slack, etc.)
4. Review the notification queue to see if there are any stuck or failed notifications

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Cortex service to attempt to resend any stuck notifications
2. Verify that the notification recipients and channels are correctly configured
3. Check the network connectivity between Cortex and the notification servers and resolve any issues
4. Consider implementing a backup notification system or redundant notification channels to minimize the impact of notification errors
5. Review the Cortex configuration and logs to identify any underlying issues that may be causing the notification errors and address them accordingly.