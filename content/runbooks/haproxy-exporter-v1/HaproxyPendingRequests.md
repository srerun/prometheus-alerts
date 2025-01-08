---
title: HaproxyPendingRequests
description: Troubleshooting for alert HaproxyPendingRequests
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyPendingRequests

Some HAProxy requests are pending on {{ $labels.fqdn }}/{{ $labels.backend }} backend

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-exporter-v1.yml" "HaproxyPendingRequests" %}}

{{% comment %}}

```yaml
alert: HaproxyPendingRequests
expr: sum by (backend) (haproxy_backend_current_queue) > 0
for: 2m
labels:
    severity: warning
annotations:
    summary: HAProxy pending requests (instance {{ $labels.instance }})
    description: |-
        Some HAProxy requests are pending on {{ $labels.fqdn }}/{{ $labels.backend }} backend
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/haproxy-exporter-v1/HaproxyPendingRequests.md

```

{{% /comment %}}

</details>


Here is a runbook for the HaproxyPendingRequests alert rule:

## Meaning

The HaproxyPendingRequests alert is triggered when there are pending requests in one or more HAProxy backends. This alert indicates that HAProxy is not able to process incoming requests quickly enough, leading to a backlog of requests waiting to be processed.

## Impact

The impact of this alert is that users may experience delays or timeouts when accessing applications or services proxied by HAProxy. In severe cases, it can lead to a complete unavailability of the services. Furthermore, pending requests can also lead to increased memory usage and CPU load on the HAProxy nodes, potentially causing performance issues or even node failures.

## Diagnosis

To diagnose the root cause of the issue, follow these steps:

1. Check the HAProxy logs to identify the backend with pending requests.
2. Verify the current traffic volume and patterns to determine if there's a spike in traffic that's causing the backlog.
3. Check the HAProxy configuration to ensure that the backend is properly configured and not overwhelmed.
4. Verify that the backend servers are responding correctly and not causing the delay.
5. Check the system resources (CPU, memory, disk space) of the HAProxy nodes to ensure they are not overloaded.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the root cause of the issue (e.g. high traffic, misconfigured backend, overload on backend servers) and address it accordingly.
2. Consider increasing the capacity of the HAProxy nodes or adding more nodes to the cluster to handle the increased load.
3. Implement traffic management strategies such as load balancing, rate limiting, or queuing to manage incoming requests.
4. Optimize the HAProxy configuration to improve performance and reduce latency.
5. Monitor the situation closely and consider escalating the issue to the relevant teams (e.g. development, infrastructure) if the issue persists.