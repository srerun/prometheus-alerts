---
title: CortexNotificationAreBeingDropped
description: Troubleshooting for alert CortexNotificationAreBeingDropped
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CortexNotificationAreBeingDropped

Cortex notification are being dropped due to errors (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "cortex/cortex-internal.yml" "CortexNotificationAreBeingDropped" %}}

{{% comment %}}

```yaml
alert: CortexNotificationAreBeingDropped
expr: rate(cortex_prometheus_notifications_dropped_total[5m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Cortex notification are being dropped (instance {{ $labels.instance }})
    description: |-
        Cortex notification are being dropped due to errors (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/cortex-internal/CortexNotificationAreBeingDropped.md

```

{{% /comment %}}

</details>


## Meaning

The `CortexNotificationAreBeingDropped` alert is triggered when the rate of dropped Cortex notifications exceeds 0 within a 5-minute window. This indicates that some Cortex notifications are not being processed successfully, potentially leading to missed alerting opportunities or incomplete data.

## Impact

The impact of this alert is critical, as it may result in:

* Missed alerting opportunities, leading to delayed or missed responses to critical issues
* Incomplete or inaccurate data, affecting the reliability of monitoring and analytics
* Increased risk of system downtime or performance degradation due to undetected issues

## Diagnosis

To diagnose the root cause of the issue, follow these steps:

1. Check the Cortex logs for errors related to notification processing
2. Investigate the instance specified in the alert label (`{{ $labels.instance }}`) for any configuration issues or errors
3. Verify that the notification pipeline is properly configured and functional
4. Check the network connectivity and infrastructure supporting the notification pipeline
5. Review the `cortex_prometheus_notifications_dropped_total` metric to identify any patterns or trends in the dropped notifications

## Mitigation

To mitigate the issue, take the following steps:

1. Fix any configuration issues or errors in the Cortex instance specified in the alert label
2. Restart the Cortex service to ensure a clean slate for notification processing
3. Verify that the notification pipeline is properly configured and functional
4. Implement additional logging or monitoring to detect and alert on notification processing errors
5. Consider increasing the notification queue capacity or retry mechanisms to reduce the likelihood of dropped notifications
6. Review and optimize the notification pipeline to minimize processing errors and improve overall reliability