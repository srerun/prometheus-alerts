---
title: PrometheusAlertmanagerE2eDeadManSwitch
description: Troubleshooting for alert PrometheusAlertmanagerE2eDeadManSwitch
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusAlertmanagerE2eDeadManSwitch

Prometheus DeadManSwitch is an always-firing alert. It's used as an end-to-end test of Prometheus through the Alertmanager.

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusAlertmanagerE2eDeadManSwitch" %}}

{{% comment %}}

```yaml
alert: PrometheusAlertmanagerE2eDeadManSwitch
expr: vector(1)
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus AlertManager E2E dead man switch (instance {{ $labels.instance }})
    description: |-
        Prometheus DeadManSwitch is an always-firing alert. It's used as an end-to-end test of Prometheus through the Alertmanager.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusAlertmanagerE2eDeadManSwitch.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus Alertmanager E2E Dead Man Switch alert:

## Meaning

The Prometheus Alertmanager E2E Dead Man Switch alert is an always-firing alert that serves as an end-to-end test of Prometheus through the Alertmanager. This alert is intentionally configured to always be firing, and its purpose is to ensure that the entire monitoring pipeline, from Prometheus to Alertmanager, is functioning correctly.

## Impact

This alert has no direct impact on the system or service being monitored, as it is a self-test of the monitoring infrastructure. However, if this alert is not firing, it may indicate a problem with the monitoring setup, which could lead to undetected issues or errors in the system.

## Diagnosis

To diagnose the issue, check the following:

* Verify that Prometheus is scraping metrics correctly and that the Alertmanager is receiving and processing alerts as expected.
* Check the Prometheus and Alertmanager logs for any errors or issues that may be preventing the Dead Man Switch alert from firing.
* Ensure that the Alertmanager is configured correctly and that the Dead Man Switch alert is properly defined and enabled.

## Mitigation

If the Dead Man Switch alert is not firing, perform the following steps to mitigate the issue:

* Restart the Prometheus and Alertmanager services to ensure that they are running correctly.
* Review and update the configuration of Prometheus and Alertmanager to ensure that they are properly set up and that the Dead Man Switch alert is correctly defined.
* Verify that the network connection between Prometheus and Alertmanager is stable and that there are no firewall rules blocking communication between the two services.
* If the issue persists, seek assistance from a monitoring administrator or developer to troubleshoot and resolve the underlying issue.