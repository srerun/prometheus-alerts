---
title: EtcdHighNumberOfFailedHttpRequests
description: Troubleshooting for alert EtcdHighNumberOfFailedHttpRequests
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHighNumberOfFailedHttpRequests

More than 1% HTTP failure detected in Etcd

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdHighNumberOfFailedHttpRequests" %}}

{{% comment %}}

```yaml
alert: EtcdHighNumberOfFailedHttpRequests
expr: sum(rate(etcd_http_failed_total[1m])) BY (method) / sum(rate(etcd_http_received_total[1m])) BY (method) > 0.01
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd high number of failed HTTP requests (instance {{ $labels.instance }})
    description: |-
        More than 1% HTTP failure detected in Etcd
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdHighNumberOfFailedHttpRequests.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `EtcdHighNumberOfFailedHttpRequests`:

## Meaning

The `EtcdHighNumberOfFailedHttpRequests` alert is triggered when the rate of failed HTTP requests to Etcd exceeds 1% of the total HTTP requests received within a 1-minute window. This indicates a potential issue with Etcd's ability to handle requests successfully.

## Impact

A high number of failed HTTP requests to Etcd can have several consequences:

* Increased latency and errors for applications relying on Etcd for data storage and retrieval
* Potential data inconsistencies or losses due to failed writes or reads
* Decreased overall system reliability and availability

## Diagnosis

To diagnose the root cause of the issue, follow these steps:

1. Check the Etcd server logs for errors or warnings related to HTTP requests
2. Verify the network connectivity and configuration between the Etcd server and the clients making HTTP requests
3. Check the Etcd server's resource utilization (CPU, memory, disk space) to ensure it is not overloaded
4. Review the Etcd cluster's health and ensures that all members are in a healthy state
5. Check for any recent changes or deployments that may have introduced a regression

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Etcd server to refresh its connection to the clients and clear any temporary issues
2. Investigate and resolve any underlying issues with network connectivity or configuration
3. Implement retries and exponential backoff strategies in clients to handle temporary failures
4. Consider increasing the resources (e.g., CPU, memory) allocated to the Etcd server to handle increased load
5. Verify that the Etcd cluster is properly configured and that all members are in a healthy state

Remember to also address the root cause of the issue to prevent it from happening again in the future.