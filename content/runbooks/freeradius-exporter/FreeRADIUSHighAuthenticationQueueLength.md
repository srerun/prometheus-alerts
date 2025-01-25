---
title: FreeRADIUSHighAuthenticationQueueLength
description: Troubleshooting for alert FreeRADIUSHighAuthenticationQueueLength
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeRADIUSHighAuthenticationQueueLength

The authentication queue length ({{ $value }}) is above the threshold of 100.

<details>
  <summary>Alert Rule</summary>

{{% rule "freeradius/freeradius-exporter.yml" "FreeRADIUSHighAuthenticationQueueLength" %}}

{{% comment %}}

```yaml
alert: FreeRADIUSHighAuthenticationQueueLength
expr: freeradius_queue_len_auth > 100
for: 5m
labels:
    severity: warning
annotations:
    summary: High Authentication Queue Length
    description: The authentication queue length ({{ $value }}) is above the threshold of 100.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/freeradius-exporter/freeradiushighauthenticationqueuelength/

```

{{% /comment %}}

</details>


Here is a sample runbook for the FreeRADIUSHighAuthenticationQueueLength alert:

## Meaning

The FreeRADIUSHighAuthenticationQueueLength alert is triggered when the authentication queue length in FreeRADIUS exceeds 100 for more than 5 minutes. This indicates that there is a high number of authentication requests waiting to be processed, which can lead to delays and failures in the authentication process.

## Impact

A high authentication queue length can cause:

* Delays in authentication requests, leading to poor user experience
* Increased risk of authentication failures, resulting in denied access to resources
* Potential security risks, as unauthenticated requests may be stuck in the queue
* Performance issues with the FreeRADIUS server, leading to increased latency and potential crashes

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the FreeRADIUS server logs for any errors or issues that may be causing the authentication queue to grow.
2. Verify that the FreeRADIUS server is properly configured and that all necessary services are running.
3. Check the system resources (CPU, memory, disk space) to ensure that the server is not overwhelmed.
4. Investigate any recent changes to the network or system that may be contributing to the increase in authentication requests.
5. Review the authentication queue length graph to identify any patterns or trends that may indicate the root cause of the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Increase the resources (CPU, memory, disk space) available to the FreeRADIUS server to improve its performance and ability to process authentication requests.
2. Optimize the FreeRADIUS server configuration to improve its efficiency and reduce the authentication queue length.
3. Implement rate limiting or traffic shaping to reduce the number of authentication requests and prevent overwhelming the server.
4. Investigate and address any underlying issues that may be causing the increase in authentication requests, such as a denial-of-service attack or a misconfigured client.
5. Consider implementing a queue management system to prioritize and manage authentication requests more effectively.

Remember to update the alert annotation and runbook URL to reflect the changes made to this runbook.