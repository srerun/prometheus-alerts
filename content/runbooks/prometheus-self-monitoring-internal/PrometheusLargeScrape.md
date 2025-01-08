---
title: PrometheusLargeScrape
description: Troubleshooting for alert PrometheusLargeScrape
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusLargeScrape

Prometheus has many scrapes that exceed the sample limit

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusLargeScrape" %}}

{{% comment %}}

```yaml
alert: PrometheusLargeScrape
expr: increase(prometheus_target_scrapes_exceeded_sample_limit_total[10m]) > 10
for: 5m
labels:
    severity: warning
annotations:
    summary: Prometheus large scrape (instance {{ $labels.instance }})
    description: |-
        Prometheus has many scrapes that exceed the sample limit
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusLargeScrape.md

```

{{% /comment %}}

</details>


Here is a runbook for the PrometheusLargeScrape alert:

## Meaning

The PrometheusLargeScrape alert is triggered when Prometheus has many scrapes that exceed the sample limit, indicating that Prometheus is collecting more data than it can handle. This can lead to performance issues, increased memory usage, and inaccurate metrics.

## Impact

The impact of this alert is significant, as it can cause:

* Performance degradation of Prometheus and other dependent systems
* Increased memory usage, potentially leading to OOM (Out of Memory) errors
* Inaccurate metrics and incomplete data

## Diagnosis

To diagnose the root cause of this alert, perform the following steps:

1. Check the Prometheus logs for errors related to sample limit exceedance
2. Review the `prometheus_target_scrapes_exceeded_sample_limit_total` metric to identify the specific targets and scrapes that are exceeding the sample limit
3. Verify that the scrape configuration is correct and not overloading the system
4. Check for any recent changes to the Prometheus configuration or target configuration that may be contributing to the issue

## Mitigation

To mitigate this alert, perform the following steps:

1. Increase the sample limit for the affected targets to reduce the number of scrapes that exceed the limit
2. Optimize the scrape configuration to reduce the number of scrapes and improve efficiency
3. Implement more efficient metrics collection and storage, such as using summarization or aggregation
4. Consider increasing the resources (e.g., CPU, memory) allocated to Prometheus to handle the increased load
5. Review and adjust the alert threshold to prevent false positives or unnecessary notifications