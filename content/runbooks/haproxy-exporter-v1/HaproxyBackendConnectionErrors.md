---
title: HaproxyBackendConnectionErrors
description: Troubleshooting for alert HaproxyBackendConnectionErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyBackendConnectionErrors

Too many connection errors to {{ $labels.fqdn }}/{{ $labels.backend }} backend (> 100 req/s). Request throughput may be too high.

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-exporter-v1.yml" "HaproxyBackendConnectionErrors" %}}

{{% comment %}}

```yaml
alert: HaproxyBackendConnectionErrors
expr: sum by (backend) (rate(haproxy_backend_connection_errors_total[1m])) > 100
for: 1m
labels:
    severity: critical
annotations:
    summary: HAProxy backend connection errors (instance {{ $labels.instance }})
    description: |-
        Too many connection errors to {{ $labels.fqdn }}/{{ $labels.backend }} backend (> 100 req/s). Request throughput may be too high.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/haproxy-exporter-v1/HaproxyBackendConnectionErrors.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the HaproxyBackendConnectionErrors alert:

## Meaning

The HaproxyBackendConnectionErrors alert indicates that the rate of connection errors to a specific HAProxy backend has exceeded 100 errors per second over the past minute. This may be caused by a variety of factors, including high request throughput, backend server issues, or network connectivity problems.

## Impact

This alert may indicate that the HAProxy backend is experiencing connectivity issues, which can lead to:

* Dropped requests and reduced system availability
* Increased latency and response times
* Potential errors and failures in dependent systems

## Diagnosis

To diagnose the cause of this alert, perform the following steps:

1. Check the HAProxy logs for errors and exceptions related to the affected backend.
2. Verify the status of the backend servers and ensure they are operational and responding to requests.
3. Review network connectivity and firewall rules to ensure there are no issues blocking traffic to the backend servers.
4. Check the request throughput and latency metrics to determine if there are any unusual patterns or spikes.
5. Review the HAProxy configuration to ensure it is properly configured and optimized for the current traffic load.

## Mitigation

To mitigate the effects of this alert, perform the following steps:

1. Investigate and resolve any underlying issues with the backend servers, such as resource constraints or software errors.
2. Adjust the HAProxy configuration to improve connection timeouts, retries, and queuing to handle high request throughput.
3. Implement traffic shaping or rate limiting to prevent overwhelming the backend servers.
4. Consider scaling out the backend servers or adding additional instances to handle increased traffic.
5. Monitor the situation closely and adjust the mitigation strategies as needed to ensure system stability and availability.