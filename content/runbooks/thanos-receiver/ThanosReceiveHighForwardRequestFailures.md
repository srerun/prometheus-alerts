---
title: ThanosReceiveHighForwardRequestFailures
description: Troubleshooting for alert ThanosReceiveHighForwardRequestFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveHighForwardRequestFailures

Thanos Receive {{$labels.job}} is failing to forward {{$value | humanize}}% of requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-receiver.yml" "ThanosReceiveHighForwardRequestFailures" %}}

{{% comment %}}

```yaml
alert: ThanosReceiveHighForwardRequestFailures
expr: (sum by (job) (rate(thanos_receive_forward_requests_total{result="error", job=~".*thanos-receive.*"}[5m]))/  sum by (job) (rate(thanos_receive_forward_requests_total{job=~".*thanos-receive.*"}[5m]))) * 100 > 20
for: 5m
labels:
    severity: info
annotations:
    summary: Thanos Receive High Forward Request Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Receive {{$labels.job}} is failing to forward {{$value | humanize}}% of requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-receiver/ThanosReceiveHighForwardRequestFailures.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `ThanosReceiveHighForwardRequestFailures`:

## Meaning

The `ThanosReceiveHighForwardRequestFailures` alert is triggered when the percentage of failed forward requests in Thanos Receive exceeds 20% in a 5-minute window. This indicates that Thanos Receive is experiencing issues while forwarding requests, which can lead to data loss or inconsistencies.

## Impact

The impact of this alert is moderate to high, as failed forward requests can result in:

* Data loss or inconsistencies
* Delayed or incomplete data availability
* Increased latency or errors in dependent systems

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Receive logs for errors or exceptions related to forwarding requests.
2. Verify the configuration of Thanos Receive and ensure that it is correctly set up to forward requests.
3. Check the network connectivity and latency between Thanos Receive and the dependent systems.
4. Investigate any recent changes or updates to the Thanos Receive configuration or dependent systems.
5. Review the Prometheus metrics to identify any patterns or trends in the failed forward requests.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the Thanos Receive configuration and ensure that it is correctly set up to forward requests.
2. Restart the Thanos Receive service to ensure that it is running with the correct configuration.
3. Verify the network connectivity and latency between Thanos Receive and the dependent systems.
4. Implement temporary workarounds, such as increasing the retry count or timeout, to reduce the impact of failed forward requests.
5. Investigate and address any underlying issues or bugs in Thanos Receive or dependent systems.

Additional resources:

* [Thanos Receive documentation](https://github.com/thanos-io/thanos/blob/main/docs/components/receive.md)
* [Prometheus metrics for Thanos Receive](https://github.com/thanos-io/thanos/blob/main/docs/metrics.md)