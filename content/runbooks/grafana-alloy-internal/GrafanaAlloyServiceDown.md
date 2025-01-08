---
title: GrafanaAlloyServiceDown
description: Troubleshooting for alert GrafanaAlloyServiceDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# GrafanaAlloyServiceDown

Alloy on (instance {{ $labels.instance }}) is not responding or has stopped running.

<details>
  <summary>Alert Rule</summary>

{{% rule "grafana-alloy/grafana-alloy-internal.yml" "GrafanaAlloyServiceDown" %}}

{{% comment %}}

```yaml
alert: GrafanaAlloyServiceDown
expr: 'count by (instance) (alloy_build_info) unless count by (instance) (alloy_build_info offset 2m)  '
for: 0m
labels:
    severity: critical
annotations:
    summary: Grafana Alloy service down (instance {{ $labels.instance }})
    description: |-
        Alloy on (instance {{ $labels.instance }}) is not responding or has stopped running.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/grafana-alloy-internal/GrafanaAlloyServiceDown.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the GrafanaAlloyServiceDown alert:

## Meaning

The GrafanaAlloyServiceDown alert is triggered when the Alloy service on a Grafana instance is not responding or has stopped running. This is detected by the absence of recent `alloy_build_info` metrics from the instance.

## Impact

The impact of this alert is that the Alloy service, which is responsible for processing and rendering graphs, is unavailable. This means that users may not be able to access or update their dashboards, and any dependent services may fail or behave erratically.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Grafana instance logs for errors or exceptions related to the Alloy service.
2. Verify that the Alloy service is running and listening on the expected port.
3. Check the `alloy_build_info` metrics in Prometheus to see if there are any data gaps or anomalies.
4. Check the system resources (CPU, memory, disk space) of the Grafana instance to ensure they are within acceptable limits.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Alloy service on the affected Grafana instance.
2. Check for any software updates or patches that may be available for the Alloy service or Grafana.
3. Verify that the system resources (CPU, memory, disk space) of the Grafana instance are within acceptable limits.
4. Consider increasing the monitoring frequency or setting up additional alerting rules to catch potential issues earlier.

Note: This runbook is just a sample and may need to be adapted to your specific environment and requirements.