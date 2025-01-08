---
title: ThanosQueryGrpcServerErrorRate
description: Troubleshooting for alert ThanosQueryGrpcServerErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosQueryGrpcServerErrorRate

Thanos Query {{$labels.job}} is failing to handle {{$value | humanize}}% of requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-query.yml" "ThanosQueryGrpcServerErrorRate" %}}

{{% comment %}}

```yaml
alert: ThanosQueryGrpcServerErrorRate
expr: (sum by (job) (rate(grpc_server_handled_total{grpc_code=~"Unknown|ResourceExhausted|Internal|Unavailable|DataLoss|DeadlineExceeded", job=~".*thanos-query.*"}[5m]))/  sum by (job) (rate(grpc_server_started_total{job=~".*thanos-query.*"}[5m])) * 100 > 5)
for: 5m
labels:
    severity: warning
annotations:
    summary: Thanos Query Grpc Server Error Rate (instance {{ $labels.instance }})
    description: |-
        Thanos Query {{$labels.job}} is failing to handle {{$value | humanize}}% of requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-query/ThanosQueryGrpcServerErrorRate.md

```

{{% /comment %}}

</details>


Here is a runbook for the ThanosQueryGrpcServerErrorRate alert:

## Meaning

The ThanosQueryGrpcServerErrorRate alert is triggered when the rate of gRPC server errors for a Thanos query instance exceeds 5% over a 5-minute period. This alert indicates that the Thanos query instance is experiencing issues handling requests, which can lead to query failures and impact the overall performance of the system.

## Impact

* Delayed or failed queries
* Increased latency and response times
* Potential data loss or inconsistencies
* Impact on dependent systems and services that rely on Thanos query data

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos query instance logs for error messages related to gRPC server errors.
2. Verify that the Thanos query instance is properly configured and connected to the required dependencies.
3. Check the system resources (e.g., CPU, memory, disk space) to ensure they are not exhausted.
4. Investigate recent changes or updates that may have caused the issue.
5. Review the gRPC server metrics to identify the specific error codes and frequencies.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Thanos query instance to clear any temporary errors.
2. Check and update the Thanos query configuration to ensure it is correct and up-to-date.
3. Scale up or optimize system resources (e.g., CPU, memory, disk space) to handle increased load.
4. Implement retry mechanisms or circuit breakers to handle temporary failures.
5. Investigate and address any underlying issues or bugs causing the gRPC server errors.

Remember to review the alert annotations and labels for specific details about the affected instance and job, and to consult the Thanos query documentation and logs for further troubleshooting and resolution guidance.