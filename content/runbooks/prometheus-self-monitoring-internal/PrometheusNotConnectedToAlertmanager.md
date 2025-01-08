---
title: PrometheusNotConnectedToAlertmanager
description: Troubleshooting for alert PrometheusNotConnectedToAlertmanager
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusNotConnectedToAlertmanager

Prometheus cannot connect the alertmanager

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusNotConnectedToAlertmanager" %}}

{{% comment %}}

```yaml
alert: PrometheusNotConnectedToAlertmanager
expr: prometheus_notifications_alertmanagers_discovered < 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus not connected to alertmanager (instance {{ $labels.instance }})
    description: |-
        Prometheus cannot connect the alertmanager
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusNotConnectedToAlertmanager.md

```

{{% /comment %}}

</details>


Here is a runbook for the PrometheusNotConnectedToAlertmanager alert rule:

## Meaning
The PrometheusNotConnectedToAlertmanager alert is triggered when Prometheus is unable to connect to the Alertmanager instance. This is a critical alert as it means that Prometheus is not able to send alerts to the Alertmanager, which may cause important notifications to be lost.

## Impact
The impact of this alert is that Prometheus will not be able to send alerts to the Alertmanager, which may lead to:

* Missed notifications for critical events
* Delayed response to incidents
* Incomplete incident response due to lack of notifications
* Potential service disruptions or outages

## Diagnosis
To diagnose this issue, follow these steps:

1. Check the Prometheus logs for errors related to connecting to the Alertmanager
2. Verify that the Alertmanager is running and reachable from the Prometheus instance
3. Check the Alertmanager configuration to ensure it is correctly set up to receive alerts from Prometheus
4. Verify that the network connectivity between Prometheus and Alertmanager is not blocked by firewalls or other network issues

## Mitigation
To mitigate this issue, follow these steps:

1. Restart the Prometheus instance to attempt to re-establish the connection to the Alertmanager
2. Verify that the Alertmanager is correctly configured and running
3. Check the network connectivity between Prometheus and Alertmanager and resolve any issues found
4. If the issue persists, investigate further to determine the root cause of the connection issue and take corrective action

Additional resources:

* For more information on configuring Prometheus and Alertmanager, refer to the [Prometheus documentation](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#alertmanager)
* For troubleshooting tips, refer to the [Prometheus troubleshooting guide](https://prometheus.io/docs/prometheus/latest/troubleshooting/)