---
title: PrometheusAlertmanagerJobMissing
description: Troubleshooting for alert PrometheusAlertmanagerJobMissing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusAlertmanagerJobMissing

A Prometheus AlertManager job has disappeared

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusAlertmanagerJobMissing" %}}

{{% comment %}}

```yaml
alert: PrometheusAlertmanagerJobMissing
expr: absent(up{job="alertmanager"})
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus AlertManager job missing (instance {{ $labels.instance }})
    description: |-
        A Prometheus AlertManager job has disappeared
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusAlertmanagerJobMissing.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus Alertmanager Job Missing alert:

## Meaning

The Prometheus Alertmanager Job Missing alert is triggered when the Alertmanager job is not reporting its metrics to Prometheus. This indicates that the Alertmanager instance is not available or not properly configured.

## Impact

The impact of this alert is that Alertmanager may not be functioning correctly, which can lead to:

* Alerts not being sent to the notification channels (e.g., email, Slack)
* Delays in detecting and responding to critical incidents
* Incomplete or missing alert history
* Potential security risks due to unmonitored systems

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Alertmanager instance's logs for errors or connection issues
2. Verify that the Alertmanager configuration is correct and matches the expected settings
3. Check the Prometheus configuration to ensure that the Alertmanager job is correctly scrape-configured
4. Verify that the network connection between Alertmanager and Prometheus is stable and working

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Alertmanager instance to attempt to reconnect to Prometheus
2. Check and update the Alertmanager configuration to ensure it is correct and matches the expected settings
3. Verify that the Prometheus scrape configuration is correct and includes the Alertmanager job
4. Investigate and resolve any network connectivity issues between Alertmanager and Prometheus
5. If the issue persists, consider reaching out to the Alertmanager maintainers or Prometheus administrators for further assistance.