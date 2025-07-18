---
title: HighCPUUsage
description: Troubleshooting for alert HighCPUUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HighCPUUsage

CPU usage is above 80% for more than 5 minutes (current value: {{ $value }}%). This may indicate resource contention or overload.

<details>
  <summary>Alert Rule</summary>

{{% rule "netelastic/netelastic.yml" "HighCPUUsage" %}}

{{% comment %}}

```yaml
alert: HighCPUUsage
expr: vbng_cpuUsage > 80
for: 5m
labels:
    severity: critical
annotations:
    summary: High CPU usage on bng
    description: 'CPU usage is above 80% for more than 5 minutes (current value: {{ $value }}%). This may indicate resource contention or overload.'
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/netelastic/highcpuusage/

```

{{% /comment %}}

</details>


## Meaning

The HighCPUUsage alert is triggered when the CPU usage of a BNG (Broadband Network Gateway) exceeds 80% for more than 5 minutes. This alert indicates that the CPU is experiencing high load, which can lead to performance degradation, slow response times, and even outages.

## Impact

The impact of high CPU usage on a BNG can be significant, leading to:

* Slow or dropped packets, affecting user experience and quality of service
* Increased latency and response times, affecting real-time applications
* Potential outages or crashes, resulting in service unavailability
* Increased risk of security breaches due to inability to handle traffic efficiently

## Diagnosis

To diagnose the root cause of the HighCPUUsage alert, follow these steps:

1. Verify the CPU usage metrics to confirm the alert is valid
2. Check system logs for any errors, warnings, or unusual activity
3. Investigate recent configuration changes or software updates
4. Analyze network traffic patterns and utilization
5. Check for any signs of resource contention or overload
6. Consult with network operations and engineering teams to gather additional information

## Mitigation

To mitigate the HighCPUUsage alert, follow these steps:

1. Identify and terminate any unnecessary or resource-intensive processes
2. Adjust system configuration to optimize CPU utilization
3. Implement traffic management and QoS policies to reduce load
4. Consider upgrading hardware or increasing resources to meet growing demands
5. Schedule maintenance during off-peak hours to reduce impact on users
6. Collaborate with development teams to optimize software performance and efficiency

Note: This runbook provides general guidance and may require customization based on specific use cases and environments. It is essential to adapt and refine this process to fit your organization's specific needs and requirements.