---
title: ThanosReceiveHighHashringFileRefreshFailures
description: Troubleshooting for alert ThanosReceiveHighHashringFileRefreshFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveHighHashringFileRefreshFailures

Thanos Receive {{$labels.job}} is failing to refresh hashring file, {{$value | humanize}} of attempts failed.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-receiver.yml" "ThanosReceiveHighHashringFileRefreshFailures" %}}

{{% comment %}}

```yaml
alert: ThanosReceiveHighHashringFileRefreshFailures
expr: (sum by (job) (rate(thanos_receive_hashrings_file_errors_total{job=~".*thanos-receive.*"}[5m])) / sum by (job) (rate(thanos_receive_hashrings_file_refreshes_total{job=~".*thanos-receive.*"}[5m])) > 0)
for: 15m
labels:
    severity: warning
annotations:
    summary: Thanos Receive High Hashring File Refresh Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Receive {{$labels.job}} is failing to refresh hashring file, {{$value | humanize}} of attempts failed.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-receiver/ThanosReceiveHighHashringFileRefreshFailures.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

The ThanosReceiveHighHashringFileRefreshFailures alert is triggered when the ratio of failed Thanos Receive hashring file refreshes to total refresh attempts exceeds 0 within a 5-minute window. This indicates that Thanos Receive is experiencing issues refreshing its hashring file, which can lead to data inconsistencies and affect the overall reliability of the system.

## Impact

* Data inconsistencies and potential loss due to failed hashring file refreshes
* Reduced reliability of the system, leading to potential outages or performance degradation
* Increased latency and errors in downstream systems that rely on Thanos Receive

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Receive logs for errors related to hashring file refreshes
2. Verify that the Thanos Receive instance has the necessary permissions and access to the hashring file
3. Check the network connectivity and latency between Thanos Receive and the hashring file storage
4. Investigate if there are any recent changes to the Thanos Receive configuration or hashring file format
5. Review the Thanos Receive metrics to identify any patterns or trends in the failure rates

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Thanos Receive instance to attempt to recover from the failed hashring file refreshes
2. Verify that the hashring file is up-to-date and accessible by Thanos Receive
3. Check and update the Thanos Receive configuration to ensure it is correct and valid
4. Implement temporary workarounds, such as increasing the hashring file refresh interval or reducing the load on Thanos Receive
5. Schedule a maintenance window to perform a thorough investigation and resolution of the underlying issue