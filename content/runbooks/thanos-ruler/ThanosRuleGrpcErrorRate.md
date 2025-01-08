---
title: ThanosRuleGrpcErrorRate
description: Troubleshooting for alert ThanosRuleGrpcErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleGrpcErrorRate

Thanos Rule {{$labels.job}} is failing to handle {{$value | humanize}}% of requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-ruler.yml" "ThanosRuleGrpcErrorRate" %}}

{{% comment %}}

```yaml
alert: ThanosRuleGrpcErrorRate
expr: (sum by (job, instance) (rate(grpc_server_handled_total{grpc_code=~"Unknown|ResourceExhausted|Internal|Unavailable|DataLoss|DeadlineExceeded", job=~".*thanos-rule.*"}[5m]))/  sum by (job, instance) (rate(grpc_server_started_total{job=~".*thanos-rule.*"}[5m])) * 100 > 5)
for: 5m
labels:
    severity: warning
annotations:
    summary: Thanos Rule Grpc Error Rate (instance {{ $labels.instance }})
    description: |-
        Thanos Rule {{$labels.job}} is failing to handle {{$value | humanize}}% of requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-ruler/ThanosRuleGrpcErrorRate.md

```

{{% /comment %}}

</details>


Here is the runbook for the ThanosRuleGrpcErrorRate alert:

## Meaning

The ThanosRuleGrpcErrorRate alert is triggered when the rate of gRPC errors for a Thanos Rule instance exceeds 5% of the total requests handled by the instance. This alert indicates that the Thanos Rule instance is experiencing issues with handling requests, which can lead to data loss, rule evaluation failures, or other problems.

## Impact

The impact of this alert can be significant, as it may lead to:

* Data loss or inconsistencies in the Thanos database
* Rule evaluation failures, causing alerts and notifications to fail
* Increased latency or timeouts in the system
* Increased load on the system, leading to performance degradation

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Rule instance logs for errors related to gRPC handling
2. Verify that the instance is not experiencing high CPU or memory usage
3. Check the network connectivity and configuration for issues
4. Review the rule evaluation metrics to identify any patterns or trends
5. Check for any recent changes or updates to the Thanos Rule configuration or code

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Thanos Rule instance to clear any temporary issues
2. Check and update the gRPC configuration and settings
3. Verify that the instance has sufficient resources (CPU, memory, etc.)
4. Implement circuit breakers or retries to handle temporary gRPC errors
5. Review and optimize the rule evaluation configuration to reduce load and errors
6. Consider enabling gRPC tracing to gain more insight into the issue
7. If the issue persists, consider rolling back recent changes or updates to the Thanos Rule configuration or code.