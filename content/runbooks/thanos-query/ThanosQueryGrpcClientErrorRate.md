---
title: ThanosQueryGrpcClientErrorRate
description: Troubleshooting for alert ThanosQueryGrpcClientErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosQueryGrpcClientErrorRate

Thanos Query {{$labels.job}} is failing to send {{$value | humanize}}% of requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-query.yml" "ThanosQueryGrpcClientErrorRate" %}}

{{% comment %}}

```yaml
alert: ThanosQueryGrpcClientErrorRate
expr: (sum by (job) (rate(grpc_client_handled_total{grpc_code!="OK", job=~".*thanos-query.*"}[5m])) / sum by (job) (rate(grpc_client_started_total{job=~".*thanos-query.*"}[5m]))) * 100 > 5
for: 5m
labels:
    severity: warning
annotations:
    summary: Thanos Query Grpc Client Error Rate (instance {{ $labels.instance }})
    description: |-
        Thanos Query {{$labels.job}} is failing to send {{$value | humanize}}% of requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-query/ThanosQueryGrpcClientErrorRate.md

```

{{% /comment %}}

</details>


## Meaning

The ThanosQueryGrpcClientErrorRate alert rule is triggered when the error rate of the Thanos Query gRPC client exceeds 5% over a 5-minute period. This indicates that the Thanos Query instance is experiencing issues when sending requests to the gRPC server.

## Impact

* Thanos Query instance is not functioning correctly, leading to incomplete or inaccurate query results.
* Potential impact on downstream systems that rely on Thanos Query, such as Grafana dashboards or alerting systems.
* Downtime or errors in the monitoring and alerting system.

## Diagnosis

1. Check the Thanos Query instance logs for errors related to gRPC client connections.
2. Verify the gRPC server is operational and responding to requests.
3. Investigate recent changes to the Thanos Query configuration or environment that may be causing the issue.
4. Check the network connectivity between the Thanos Query instance and the gRPC server.

## Mitigation

1. Restart the Thanos Query instance to reset the gRPC client connections.
2. Verify the gRPC server is configured correctly and reachable by the Thanos Query instance.
3. Check for and resolve any network connectivity issues between the Thanos Query instance and the gRPC server.
4. Implement additional logging and monitoring to detect future issues with the gRPC client connections.
5. Consider increasing the gRPC client timeout or retries to improve resilience.

Note: For more detailed steps and troubleshooting guidance, refer to the runbook link provided in the alert annotation: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-query/ThanosQueryGrpcClientErrorRate.md