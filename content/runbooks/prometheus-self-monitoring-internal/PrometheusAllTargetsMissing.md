---
title: PrometheusAllTargetsMissing
description: Troubleshooting for alert PrometheusAllTargetsMissing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusAllTargetsMissing

A Prometheus job does not have living target anymore.

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusAllTargetsMissing" %}}

{{% comment %}}

```yaml
alert: PrometheusAllTargetsMissing
expr: sum by (job) (up) == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus all targets missing (instance {{ $labels.instance }})
    description: |-
        A Prometheus job does not have living target anymore.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusAllTargetsMissing.md

```

{{% /comment %}}

</details>


## Meaning

The PrometheusAllTargetsMissing alert is triggered when all targets for a Prometheus job are missing. This means that Prometheus is not scraping any metrics from any instances for a particular job. This is a critical alert as it indicates that Prometheus is not functioning correctly and metrics are not being collected.

## Impact

The impact of this alert is that there will be a gap in metric collection, leading to incomplete or missing data. This can cause issues with monitoring, alerting, and troubleshooting. Without metrics, it is difficult to detect issues or identify trends, which can lead to prolonged downtime or degraded performance.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus targets page to verify that all targets are indeed missing.
2. Verify that the Prometheus configuration is correct and the job is correctly defined.
3. Check the Prometheus logs for any errors or issues related to target discovery or scraping.
4. Verify that the instances being scraped are running and responding to requests.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the instances being scraped to ensure they are running and responsive.
2. Verify that the Prometheus configuration is correct and the job is correctly defined.
3. Check the network connectivity between Prometheus and the instances being scraped.
4. Restart the Prometheus server to reload the configuration and targets.
5. If the issue persists, check the Prometheus logs for any errors or issues related to target discovery or scraping.

Note: For more detailed steps and additional troubleshooting, refer to the [runbook](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusAllTargetsMissing.md)