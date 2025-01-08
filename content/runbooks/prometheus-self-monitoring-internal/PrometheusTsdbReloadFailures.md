---
title: PrometheusTsdbReloadFailures
description: Troubleshooting for alert PrometheusTsdbReloadFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTsdbReloadFailures

Prometheus encountered {{ $value }} TSDB reload failures

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTsdbReloadFailures" %}}

{{% comment %}}

```yaml
alert: PrometheusTsdbReloadFailures
expr: increase(prometheus_tsdb_reloads_failures_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus TSDB reload failures (instance {{ $labels.instance }})
    description: |-
        Prometheus encountered {{ $value }} TSDB reload failures
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTsdbReloadFailures.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the PrometheusTsdbReloadFailures alert:

## Meaning

The PrometheusTsdbReloadFailures alert indicates that Prometheus has encountered failures while reloading its Time Series Database (TSDB). This is a critical alert as it may lead to data loss, inconsistent query results, or even render Prometheus unusable.

## Impact

* Data loss: TSDB reload failures can result in loss of metric data, making it impossible to query or alert on historical data.
* Inconsistent query results: Failures during TSDB reload can lead to inconsistent query results, affecting the accuracy of monitoring and alerting.
* Prometheus unavailability: In extreme cases, TSDB reload failures can render Prometheus unusable, causing a complete loss of monitoring and alerting capabilities.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus logs for errors related to TSDB reloads.
2. Verify that the disk usage is within acceptable limits, and there is enough disk space available for TSDB to function properly.
3. Review the TSDB configuration to ensure it is correctly set up and not causing issues.
4. Check for any system-level issues, such as high CPU usage, memory pressure, or network connectivity problems, that might be affecting TSDB reloads.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Prometheus instance to attempt a TSDB reload.
2. Check the disk usage and free up disk space if necessary.
3. Verify the TSDB configuration and make adjustments as needed.
4. Consider increasing the resources (e.g., CPU, memory) allocated to the Prometheus instance if system-level issues are suspected.
5. If the issue persists, consider seeking assistance from a Prometheus expert or the community.

Remember to monitor the situation closely and adjust the mitigation steps as needed to prevent further TSDB reload failures.