---
title: PrometheusTimeseriesCardinality
description: Troubleshooting for alert PrometheusTimeseriesCardinality
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTimeseriesCardinality

The "{{ $labels.name }}" timeseries cardinality is getting very high: {{ $value }}

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTimeseriesCardinality" %}}

{{% comment %}}

```yaml
alert: PrometheusTimeseriesCardinality
expr: label_replace(count by(__name__) ({__name__=~".+"}), "name", "$1", "__name__", "(.+)") > 10000
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus timeseries cardinality (instance {{ $labels.instance }})
    description: |-
        The "{{ $labels.name }}" timeseries cardinality is getting very high: {{ $value }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTimeseriesCardinality.md

```

{{% /comment %}}

</details>


Here is a runbook for the PrometheusTimeseriesCardinality alert rule:

## Meaning

The PrometheusTimeseriesCardinality alert is triggered when the number of unique time series in Prometheus exceeds 10,000. This can indicate a high cardinality of metrics, which can lead to performance issues and increased memory usage in Prometheus.

## Impact

If left unchecked, high time series cardinality can cause:

* Performance degradation in Prometheus, leading to slow query times and increased latency
* Increased memory usage, which can lead to OOM (Out of Memory) errors and crashes
* Decreased data retention and accuracy due to the high volume of metrics

## Diagnosis

To diagnose the root cause of the high time series cardinality, follow these steps:

1. Identify the specific metric with high cardinality using the `label_replace` function in the alert expression
2. Check the metric's label values to determine which dimension is causing the high cardinality (e.g., instance, namespace, pod, etc.)
3. Review the metric's configuration and ensure it is properly configured to avoid unnecessary label creation
4. Consult with the team responsible for the metric to determine if there are any changes that can be made to reduce cardinality

## Mitigation

To mitigate the high time series cardinality, follow these steps:

1. Optimize metric configurations to reduce label creation and cardinality
2. Implement data aggregation and rollup to reduce the number of time series
3. Consider implementing a metric filtering or dropping solution to remove unnecessary metrics
4. Increase the resources (CPU, memory, disk) allocated to Prometheus to handle the increased load, if necessary
5. Monitor the metric's cardinality and adjust the alert threshold as needed to ensure timely detection of high cardinality issues.