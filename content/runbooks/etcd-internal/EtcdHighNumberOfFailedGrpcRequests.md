---
title: EtcdHighNumberOfFailedGrpcRequests
description: Troubleshooting for alert EtcdHighNumberOfFailedGrpcRequests
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHighNumberOfFailedGrpcRequests

More than 1% GRPC request failure detected in Etcd

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdHighNumberOfFailedGrpcRequests" %}}

{{% comment %}}

```yaml
alert: EtcdHighNumberOfFailedGrpcRequests
expr: sum(rate(grpc_server_handled_total{grpc_code!="OK"}[1m])) BY (grpc_service, grpc_method) / sum(rate(grpc_server_handled_total[1m])) BY (grpc_service, grpc_method) > 0.01
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd high number of failed GRPC requests (instance {{ $labels.instance }})
    description: |-
        More than 1% GRPC request failure detected in Etcd
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdHighNumberOfFailedGrpcRequests.md

```

{{% /comment %}}

</details>


Here is a runbook for the EtcdHighNumberOfFailedGrpcRequests alert:

## Meaning

The EtcdHighNumberOfFailedGrpcRequests alert is triggered when the rate of failed GRPC requests in Etcd exceeds 1% of the total requests over a 1-minute period. This alert indicates that there is an issue with Etcd's GRPC requests, which can impact the overall reliability and performance of the system.

## Impact

If this alert is not addressed, it can lead to:

* Increased latency and errors in Etcd operations
* Decreased system reliability and availability
* Potential data inconsistencies and loss
* Increased load on the system, leading to further performance degradation

## Diagnosis

To diagnose the root cause of the issue, follow these steps:

1. Check the Etcd server logs for errors or warnings related to GRPC requests
2. Verify that the Etcd cluster is properly configured and that all nodes are synchronized
3. Check for any network connectivity issues between Etcd nodes
4. Verify that the GRPC requests are not being throttled or rate-limited
5. Check for any known bugs or issues in Etcd or related components

## Mitigation

To mitigate the issue, follow these steps:

1. Check the Etcd documentation for troubleshooting GRPC request failures
2. Verify that the Etcd cluster is properly scaled to handle the current load
3. Implement measures to reduce the load on the Etcd cluster, such as load balancing or caching
4. Consider increasing the resources (CPU, memory, etc.) allocated to the Etcd nodes
5. If the issue persists, consider rolling back to a previous version of Etcd or related components

Note: The runbook URL provided in the annotation points to a more comprehensive runbook that may include additional steps and details.