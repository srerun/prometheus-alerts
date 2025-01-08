---
title: EtcdGrpcRequestsSlow
description: Troubleshooting for alert EtcdGrpcRequestsSlow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdGrpcRequestsSlow

GRPC requests slowing down, 99th percentile is over 0.15s

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdGrpcRequestsSlow" %}}

{{% comment %}}

```yaml
alert: EtcdGrpcRequestsSlow
expr: histogram_quantile(0.99, sum(rate(grpc_server_handling_seconds_bucket{grpc_type="unary"}[1m])) by (grpc_service, grpc_method, le)) > 0.15
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd GRPC requests slow (instance {{ $labels.instance }})
    description: |-
        GRPC requests slowing down, 99th percentile is over 0.15s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdGrpcRequestsSlow.md

```

{{% /comment %}}

</details>


Here is a runbook for the EtcdGrpcRequestsSlow alert:

## Meaning

The EtcdGrpcRequestsSlow alert is triggered when the 99th percentile of etcd GRPC request handling time exceeds 0.15 seconds over a 1-minute window. This indicates that etcd GRPC requests are taking longer than expected to process, which can impact the overall performance and reliability of the etcd cluster.

## Impact

The impact of slow etcd GRPC requests can be significant, leading to:

* Increased latency for etcd operations
* Reduced throughput for etcd requests
* Potential timeouts and errors for clients relying on etcd
* Cascade failures in dependent systems

## Diagnosis

To diagnose the root cause of slow etcd GRPC requests, follow these steps:

1. Check etcd node logs for any signs of high CPU usage, memory pressure, or disk I/O issues.
2. Verify that etcd is properly configured and optimized for performance.
3. Check for any network issues or congestion that may be contributing to slow request processing.
4. Use tools like `etcdctl` or `grpcurl` to inspect etcd cluster health and verify that no nodes are experiencing high latency.
5. Review etcd request metrics (e.g., `grpc_server_handling_seconds_bucket`) to identify specific methods or services experiencing slow request processing.

## Mitigation

To mitigate the impact of slow etcd GRPC requests, follow these steps:

1. Apply performance optimizations to etcd nodes, such as adjusting resource allocations or tuning configuration settings.
2. Implement load balancing or sharding to distribute etcd request traffic and reduce load on individual nodes.
3. Investigate and resolve any underlying network issues or congestion contributing to slow request processing.
4. Consider increasing etcd node resources (e.g., CPU, memory) to alleviate performance bottlenecks.
5. If the issue persists, consider rolling back recent changes or seeking assistance from etcd experts or the etcd community.