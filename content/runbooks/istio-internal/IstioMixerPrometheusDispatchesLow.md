---
title: IstioMixerPrometheusDispatchesLow
description: Troubleshooting for alert IstioMixerPrometheusDispatchesLow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioMixerPrometheusDispatchesLow

Number of Mixer dispatches to Prometheus is too low. Istio metrics might not be being exported properly.

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioMixerPrometheusDispatchesLow" %}}

{{% comment %}}

```yaml
alert: IstioMixerPrometheusDispatchesLow
expr: sum(rate(mixer_runtime_dispatches_total{adapter=~"prometheus"}[1m])) < 180
for: 1m
labels:
    severity: warning
annotations:
    summary: Istio Mixer Prometheus dispatches low (instance {{ $labels.instance }})
    description: |-
        Number of Mixer dispatches to Prometheus is too low. Istio metrics might not be being exported properly.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioMixerPrometheusDispatchesLow.md

```

{{% /comment %}}

</details>


Here is a runbook for the IstioMixerPrometheusDispatchesLow alert rule:

## Meaning

The IstioMixerPrometheusDispatchesLow alert is triggered when the number of Mixer dispatches to Prometheus falls below a certain threshold (180 dispatches per minute). This could indicate that Istio metrics are not being exported properly, which can lead to incomplete or inaccurate monitoring data.

## Impact

The impact of this alert is that critical metrics about the Istio service mesh may not be available, making it difficult to monitor and troubleshoot issues in the system. This could lead to delayed or incomplete problem identification and resolution, potentially affecting the overall reliability and performance of the system.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Mixer logs for any errors or issues related to dispatching metrics to Prometheus.
2. Verify that the Prometheus instance is properly configured and running.
3. Check the Istio configuration to ensure that metrics are being sent to the correct Prometheus instance.
4. Review the Mixer runtime metrics to identify any trends or patterns that may indicate the root cause of the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Mixer component to ensure that it is properly configured and running.
2. Verify that the Prometheus instance is properly configured and running.
3. Check the Istio configuration to ensure that metrics are being sent to the correct Prometheus instance.
4. If the issue persists, consider scaling up the Mixer component or adjusting the thresholds in the alert rule.
5. Review the Mixer runtime metrics to identify any trends or patterns that may indicate the root cause of the issue and take corrective action.

Note: This runbook is just a starting point, and you may need to tailor it to your specific environment and requirements.