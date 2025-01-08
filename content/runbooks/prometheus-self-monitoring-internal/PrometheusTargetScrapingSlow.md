---
title: PrometheusTargetScrapingSlow
description: Troubleshooting for alert PrometheusTargetScrapingSlow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTargetScrapingSlow

Prometheus is scraping exporters slowly since it exceeded the requested interval time. Your Prometheus server is under-provisioned.

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTargetScrapingSlow" %}}

{{% comment %}}

```yaml
alert: PrometheusTargetScrapingSlow
expr: prometheus_target_interval_length_seconds{quantile="0.9"} / on (interval, instance, job) prometheus_target_interval_length_seconds{quantile="0.5"} > 1.05
for: 5m
labels:
    severity: warning
annotations:
    summary: Prometheus target scraping slow (instance {{ $labels.instance }})
    description: |-
        Prometheus is scraping exporters slowly since it exceeded the requested interval time. Your Prometheus server is under-provisioned.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTargetScrapingSlow.md

```

{{% /comment %}}

</details>


## Meaning

The `PrometheusTargetScrapingSlow` alert is triggered when the 90th percentile of Prometheus target scraping intervals exceeds 1.05 times the 50th percentile for more than 5 minutes. This indicates that Prometheus is taking longer than expected to scrape targets, which can lead to delayed or missed metric updates.

## Impact

The impact of this alert is:

* Delays in metric updates, which can lead to delayed alerting and decision-making.
* Increased latency in monitoring and debugging systems.
* Potential loss of data or incomplete metric sets.
* Increased load on Prometheus, leading to further performance issues.

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the Prometheus server's resource utilization (CPU, memory, and disk) to identify if it is under-provisioned.
2. Investigate the target exporters to see if they are experiencing high load or latency.
3. Review the Prometheus configuration to ensure that the scrape interval and timeout values are reasonable.
4. Verify that there are no network connectivity issues between Prometheus and the target exporters.

## Mitigation

To mitigate this alert, follow these steps:

1. **Increase Prometheus server resources**: Upgrade the Prometheus server's resources (e.g., CPU, memory, and disk) to handle the increased load.
2. **Optimize target exporters**: Identify and optimize the slowest target exporters to reduce their latency and load.
3. **Adjust scrape interval and timeout**: Adjust the Prometheus configuration to increase the scrape interval and timeout values to allow more time for scraping targets.
4. **Distribute scrape load**: Consider distributing the scrape load across multiple Prometheus servers or instances to reduce the load on individual servers.
5. **Monitor and alert on Prometheus performance**: Set up monitoring and alerts for Prometheus performance metrics (e.g., `prometheus_target_interval_length_seconds`) to catch performance issues earlier.