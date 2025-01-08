---
title: EtcdHttpRequestsSlow
description: Troubleshooting for alert EtcdHttpRequestsSlow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHttpRequestsSlow

HTTP requests slowing down, 99th percentile is over 0.15s

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdHttpRequestsSlow" %}}

{{% comment %}}

```yaml
alert: EtcdHttpRequestsSlow
expr: histogram_quantile(0.99, rate(etcd_http_successful_duration_seconds_bucket[1m])) > 0.15
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd HTTP requests slow (instance {{ $labels.instance }})
    description: |-
        HTTP requests slowing down, 99th percentile is over 0.15s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdHttpRequestsSlow.md

```

{{% /comment %}}

</details>


Here is a runbook for the EtcdHttpRequestsSlow alert:

## Meaning

The EtcdHttpRequestsSlow alert is triggered when the 99th percentile of etcd HTTP request durations exceeds 0.15 seconds over a 1-minute window. This indicates that etcd is experiencing slower-than-expected HTTP request processing times, which can impact the overall performance and reliability of the etcd cluster.

## Impact

* Slower etcd request processing can lead to increased latency and decreased responsiveness in dependent systems.
* Prolonged periods of slow request processing can cause etcd to become unavailable or even lead to cluster instability.
* This can have a cascading effect on the overall system, causing failures or errors in dependent applications and services.

## Diagnosis

To diagnose the root cause of the slow etcd HTTP requests, follow these steps:

1. Review etcd logs for any errors or warnings related to request processing.
2. Check the etcd cluster's resource utilization (CPU, memory, disk) to ensure it is within acceptable limits.
3. Verify that the etcd cluster is properly configured and that no network issues are affecting communication between nodes.
4. Investigate any recent changes to the etcd configuration, network topology, or dependent systems that may be contributing to the slow request processing.
5. Use tools like `etcdctl` or `curl` to verify etcd's response times and latency.

## Mitigation

To mitigate the impact of slow etcd HTTP requests, follow these steps:

1. Check for any stuck or slow requests in etcd and cancel them if necessary.
2. Implement request timeouts to prevent slow requests from blocking other requests.
3. Consider increasing the etcd cluster's resources (e.g., adding more nodes) to handle increased load.
4. Optimize etcd configuration for better performance, such as adjusting the `--listen-peer-urls` or `--listen-client-urls` settings.
5. Implement load balancing or proxying to reduce the load on individual etcd nodes.
6. Consider upgrading etcd to a newer version that includes performance enhancements.

Remember to investigate and address the root cause of the slow request processing to prevent future occurrences of this alert.