---
title: PrometheusTargetMissing
description: Troubleshooting for alert PrometheusTargetMissing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTargetMissing

A Prometheus target has disappeared. An exporter might be crashed.

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTargetMissing" %}}

{{% comment %}}

```yaml
alert: PrometheusTargetMissing
expr: up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus target missing (instance {{ $labels.instance }})
    description: |-
        A Prometheus target has disappeared. An exporter might be crashed.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTargetMissing.md

```

{{% /comment %}}

</details>


Here is a runbook for the PrometheusTargetMissing alert:

## Meaning

The PrometheusTargetMissing alert is triggered when a Prometheus target (an instance that exposes metrics to Prometheus) becomes unresponsive or disappears. This can indicate a problem with the exporter (the component that exposes the metrics) or the underlying system.

## Impact

The impact of this alert is high, as it means that Prometheus is no longer receiving metrics from the affected target. This can lead to:

* Loss of visibility into the system's performance and health
* Delayed detection of potential issues
* Inability to trigger alerts and notifications based on metrics from the missing target

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus target's logs for errors or signs of a crash
2. Verify that the exporter is running and configured correctly
3. Check the network connectivity between Prometheus and the target
4. Investigate any recent changes to the system or exporter configuration
5. Review the Prometheus metrics to identify any trends or patterns that may indicate the cause of the issue

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the exporter service (if it's not already running)
2. Verify that the exporter is configured correctly and pointing to the correct metrics endpoint
3. Check for any firewall or network connectivity issues between Prometheus and the target
4. If the issue persists, consider redeploying the exporter or restarting the underlying system
5. Investigate and address any underlying causes of the issue, such as resource constraints or configuration errors.