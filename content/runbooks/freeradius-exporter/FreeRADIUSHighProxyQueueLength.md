---
title: FreeRADIUSHighProxyQueueLength
description: Troubleshooting for alert FreeRADIUSHighProxyQueueLength
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeRADIUSHighProxyQueueLength

The proxy queue length ({{ $value }}) is above the threshold of 100.

<details>
  <summary>Alert Rule</summary>

{{% rule "freeradius/freeradius-exporter.yml" "FreeRADIUSHighProxyQueueLength" %}}

{{% comment %}}

```yaml
alert: FreeRADIUSHighProxyQueueLength
expr: freeradius_queue_len_proxy > 100
for: 5m
labels:
    severity: warning
annotations:
    summary: High Proxy Queue Length
    description: The proxy queue length ({{ $value }}) is above the threshold of 100.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/freeradius-exporter/freeradiushighproxyqueuelength/

```

{{% /comment %}}

</details>


Here is a runbook for the FreeRADIUSHighProxyQueueLength alert rule:

## Meaning

The FreeRADIUSHighProxyQueueLength alert is triggered when the length of the FreeRADIUS proxy queue exceeds 100 for a period of 5 minutes or more. This indicates that the FreeRADIUS server is unable to process proxy requests efficiently, leading to a backlog of requests in the queue.

## Impact

A high proxy queue length can cause delays in authentication and authorization requests, leading to a poor user experience. This can result in:

* Slow or failed logins
* Delays in accessing network resources
* Increased latency in authentication responses
* Potential security risks due to delayed authentication decisions

## Diagnosis

To diagnose the root cause of the high proxy queue length, follow these steps:

1. Check the FreeRADIUS server logs for any errors or warnings related to proxy requests.
2. Verify that the FreeRADIUS server is running with sufficient resources (CPU, memory, and disk space).
3. Check the network connectivity and latency between the FreeRADIUS server and the proxy endpoint.
4. Review the FreeRADIUS configuration to ensure that the proxy queue settings are correctly configured.
5. Check for any recent changes or updates to the FreeRADIUS server or its dependencies.

## Mitigation

To mitigate the effects of a high proxy queue length, follow these steps:

1. **Increase the proxy queue size**: Temporarily increase the proxy queue size to reduce the backlog of requests.
2. **Optimize FreeRADIUS server resources**: Ensure that the FreeRADIUS server is running with sufficient resources (CPU, memory, and disk space) to process proxy requests efficiently.
3. **Check and resolve network connectivity issues**: Verify that the network connectivity and latency between the FreeRADIUS server and the proxy endpoint are within acceptable limits.
4. **Tune FreeRADIUS configuration**: Review and optimize the FreeRADIUS configuration to improve proxy request processing efficiency.
5. **Implement queue monitoring and alerting**: Set up monitoring and alerting for the proxy queue length to detect potential issues earlier.
6. **Consider implementing a load balancer or additional FreeRADIUS servers**: If the proxy queue length remains consistently high, consider implementing a load balancer or adding additional FreeRADIUS servers to distribute the load and improve responsiveness.

Remember to investigate and address the root cause of the high proxy queue length to prevent future occurrences of this alert.