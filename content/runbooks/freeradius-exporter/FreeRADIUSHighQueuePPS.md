---
title: FreeRADIUSHighQueuePPS
description: Troubleshooting for alert FreeRADIUSHighQueuePPS
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeRADIUSHighQueuePPS

The FreeRADIUS server is processing a high number of packets per second ({{ $value }}). Check server load.

<details>
  <summary>Alert Rule</summary>

{{% rule "freeradius/freeradius-exporter.yml" "FreeRADIUSHighQueuePPS" %}}

{{% comment %}}

```yaml
alert: FreeRADIUSHighQueuePPS
expr: freeradius_queue_pps_in > 1000 or freeradius_queue_pps_out > 1000
for: 5m
labels:
    severity: warning
annotations:
    summary: High Queue Packets Per Second
    description: The FreeRADIUS server is processing a high number of packets per second ({{ $value }}). Check server load.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/freeradius-exporter/freeradiushighqueuepps/

```

{{% /comment %}}

</details>


Here is a runbook for the FreeRADIUSHighQueuePPS alert rule:

## Meaning

The FreeRADIUSHighQueuePPS alert is triggered when the FreeRADIUS server is processing a high number of packets per second, either incoming or outgoing. This can indicate a potential performance issue or a sudden increase in traffic.

## Impact

If not addressed, a high queue packets per second can lead to:

* Delayed or dropped packets, resulting in authentication failures or service disruptions
* Increased latency and response times for RADIUS requests
* Overloaded FreeRADIUS servers, potentially leading to crashes or instability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the FreeRADIUS server logs for any errors or warnings related to high traffic or packet processing
2. Investigate the current traffic patterns and identify any unusual spikes or trends
3. Verify that the FreeRADIUS server is properly configured and optimized for the current workload
4. Review system resources (CPU, memory, disk usage) to ensure they are not overwhelmed
5. Check for any recent changes or updates to the FreeRADIUS configuration or environment

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and address any underlying causes of the high traffic, such as a misconfigured client or an attack
2. Implement rate limiting or traffic shaping to reduce the load on the FreeRADIUS server
3. Consider scaling up or distributing the FreeRADIUS server to handle increased traffic demands
4. Optimize the FreeRADIUS server configuration for better performance and packet processing efficiency
5. Monitor the situation closely and adjust as needed to prevent further issues