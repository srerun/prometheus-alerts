---
title: CortexNotConnectedToAlertmanager
description: Troubleshooting for alert CortexNotConnectedToAlertmanager
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CortexNotConnectedToAlertmanager

Cortex not connected to Alertmanager (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "cortex/coretex-internal.yml" "CortexNotConnectedToAlertmanager" %}}

{{% comment %}}

```yaml
alert: CortexNotConnectedToAlertmanager
expr: cortex_prometheus_notifications_alertmanagers_discovered < 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Cortex not connected to Alertmanager (instance {{ $labels.instance }})
    description: |-
        Cortex not connected to Alertmanager (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/coretex-internal/CortexNotConnectedToAlertmanager.md

```

{{% /comment %}}

</details>


## Meaning

The CortexNotConnectedToAlertmanager alert is triggered when Cortex is not connected to Alertmanager. This means that Cortex, a crucial component of the monitoring system, is unable to send alerts to Alertmanager, which is responsible for handling and notifying teams of critical incidents. This disconnection can lead to delayed or missed alerts, potentially causing significant impact to the system and business.

## Impact

The impact of this alert is significant, as it can lead to:

* Delayed or missed alerts, causing critical incidents to go unnoticed
* Unmonitored system errors, leading to prolonged downtime or data loss
* Inefficient incident response, resulting in delayed resolution and increased mean time to detect (MTTD) and mean time to resolve (MTTR)
* Potential revenue loss and damage to reputation due to unplanned outages

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cortex logs for errors related to Alertmanager connectivity
2. Verify the Alertmanager configuration and ensure it is correctly set up
3. Check the network connectivity between Cortex and Alertmanager
4. Review the Cortex and Alertmanager pod/container statuses to ensure they are running correctly
5. Check the Prometheus notifications metrics to ensure they are being generated correctly

## Mitigation

To mitigate this issue, follow these steps:

1. Check the Cortex configuration and ensure it is correctly set up to connect to Alertmanager
2. Restart the Cortex service to re-establish the connection to Alertmanager
3. Verify that the Alertmanager API is accessible from Cortex
4. Check the network connectivity between Cortex and Alertmanager and resolve any issues
5. Consider implementing a fallback or redundant Alertmanager instance to ensure high availability

Additional resources:

* Refer to the Cortex documentation for configuration and troubleshooting guides
* Review the Alertmanager documentation for configuration and troubleshooting guides
* Consult with the monitoring team or system administrators for assistance with diagnosis and mitigation.