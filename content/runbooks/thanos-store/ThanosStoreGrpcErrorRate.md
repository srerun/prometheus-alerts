---
title: ThanosStoreGrpcErrorRate
description: Troubleshooting for alert ThanosStoreGrpcErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosStoreGrpcErrorRate

Thanos Store {{$labels.job}} is failing to handle {{$value | humanize}}% of requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-store.yml" "ThanosStoreGrpcErrorRate" %}}

{{% comment %}}

```yaml
alert: ThanosStoreGrpcErrorRate
expr: (sum by (job) (rate(grpc_server_handled_total{grpc_code=~"Unknown|ResourceExhausted|Internal|Unavailable|DataLoss|DeadlineExceeded", job=~".*thanos-store.*"}[5m]))/  sum by (job) (rate(grpc_server_started_total{job=~".*thanos-store.*"}[5m])) * 100 > 5)
for: 5m
labels:
    severity: warning
annotations:
    summary: Thanos Store Grpc Error Rate (instance {{ $labels.instance }})
    description: |-
        Thanos Store {{$labels.job}} is failing to handle {{$value | humanize}}% of requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-store/ThanosStoreGrpcErrorRate.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the ThanosStoreGrpcErrorRate alert:

## Meaning

The ThanosStoreGrpcErrorRate alert is triggered when the error rate of a Thanos Store instance exceeds 5% of all requests over a 5-minute period. This indicates that the Thanos Store instance is experiencing issues handling gRPC requests, which can lead to data loss or inconsistencies.

## Impact

The impact of this alert can be significant, as it affects the reliability and availability of the Thanos Store instance. If left unchecked, it can lead to:

* Data loss or inconsistencies
* Increased latency or timeouts for requests
* Reduced overall system performance
* Potential cascading failures in dependent systems

## Diagnosis

To diagnose the root cause of this alert, perform the following steps:

1. **Check the Thanos Store instance logs**: Review the logs of the affected Thanos Store instance to identify the specific errors causing the grpc errors.
2. **Verify gRPC request traffic**: Analyze the gRPC request traffic to identify any patterns or anomalies that may be contributing to the error rate.
3. **Check system resource utilization**: Verify that the Thanos Store instance has sufficient system resources (e.g., CPU, memory, disk space) to handle the request load.
4. **Review Thanos Store configuration**: Check the Thanos Store configuration to ensure it is correctly set up and tuned for the current workload.

## Mitigation

To mitigate this alert, perform the following steps:

1. **Restart the Thanos Store instance**: If the error rate is caused by a transient issue, restarting the Thanos Store instance may resolve the issue.
2. **Adjust gRPC request timeouts**: Consider increasing the gRPC request timeouts to reduce the error rate.
3. **Scale the Thanos Store instance**: If the error rate is due to high request load, consider scaling the Thanos Store instance to increase its capacity.
4. **Implement retries and fallbacks**: Implement retries and fallbacks in the gRPC clients to handle temporary failures and reduce the error rate.

Remember to monitor the system closely after taking these mitigation steps to ensure the error rate returns to a normal level.