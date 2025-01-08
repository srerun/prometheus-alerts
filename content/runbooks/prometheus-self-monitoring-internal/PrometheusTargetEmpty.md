---
title: PrometheusTargetEmpty
description: Troubleshooting for alert PrometheusTargetEmpty
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTargetEmpty

Prometheus has no target in service discovery

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTargetEmpty" %}}

{{% comment %}}

```yaml
alert: PrometheusTargetEmpty
expr: prometheus_sd_discovered_targets == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus target empty (instance {{ $labels.instance }})
    description: |-
        Prometheus has no target in service discovery
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTargetEmpty.md

```

{{% /comment %}}

</details>


## Meaning

The PrometheusTargetEmpty alert is triggered when Prometheus has no targets in its service discovery, indicating that Prometheus is not scraping any metrics from any targets. This alert is critical as it means that Prometheus is not functioning correctly, and no metrics are being collected.

## Impact

The impact of this alert is high as it means that Prometheus is not monitoring any targets, resulting in:

* No metrics being collected
* No alerts being triggered
* No visibility into the system's performance and health
* Potential loss of critical monitoring data

## Diagnosis

To diagnose the cause of this alert, follow these steps:

1. Check the Prometheus configuration to ensure that the service discovery is correctly configured.
2. Verify that the targets are correctly configured and available.
3. Check the Prometheus logs for any errors or warnings related to service discovery or target scraping.
4. Use the Prometheus web interface to manually query the `prometheus_sd_discovered_targets` metric to verify the value.

## Mitigation

To mitigate this alert, follow these steps:

1. Review and update the Prometheus configuration to ensure that the service discovery is correctly configured.
2. Verify that the targets are correctly configured and available.
3. Restart the Prometheus service to re-initialize the service discovery.
4. Monitor the Prometheus logs for any errors or warnings related to service discovery or target scraping.
5. Perform a manual query of the `prometheus_sd_discovered_targets` metric to verify the value.

Note: Refer to the provided runbook URL for more detailed instructions and troubleshooting steps.