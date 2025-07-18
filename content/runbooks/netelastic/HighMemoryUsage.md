---
title: HighMemoryUsage
description: Troubleshooting for alert HighMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HighMemoryUsage

Memory usage is above 80% for more than 5 minutes (current value: {{ $value }}%). This may lead to performance degradation.

<details>
  <summary>Alert Rule</summary>

{{% rule "netelastic/netelastic.yml" "HighMemoryUsage" %}}

{{% comment %}}

```yaml
alert: HighMemoryUsage
expr: vbng_memUsage > 80
for: 5m
labels:
    severity: critical
annotations:
    summary: High memory usage on bng
    description: 'Memory usage is above 80% for more than 5 minutes (current value: {{ $value }}%). This may lead to performance degradation.'
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/netelastic/highmemoryusage/

```

{{% /comment %}}

</details>


## Meaning

The HighMemoryUsage alert is triggered when the memory usage of the BNG (Broadband Network Gateway) exceeds 80% for more than 5 minutes. This indicates a potential issue with memory allocation or utilization on the BNG, which can lead to performance degradation and negatively impact network reliability.

## Impact

* Performance degradation: High memory usage can cause the BNG to slow down, leading to delayed or dropped packets, and ultimately affecting the quality of service for end-users.
* Network reliability: Prolonged high memory usage can cause the BNG to become unstable, leading to crashes or reboots, which can result in network outages and downtime.
* Increased latency: High memory usage can cause the BNG to take longer to process packets, leading to increased latency and affecting real-time applications such as voice and video.

## Diagnosis

To diagnose the HighMemoryUsage alert, follow these steps:

1. **Check memory usage graphs**: Review the memory usage graphs in Prometheus to identify the trend and pattern of memory usage leading up to the alert.
2. **Verify system logs**: Check the system logs for any indications of memory-related errors, such as out-of-memory errors or memory allocation failures.
3. **Investigate running processes**: Use tools like `top` or `ps` to identify any resource-intensive processes that may be contributing to high memory usage.
4. **Check system configuration**: Review the system configuration to ensure that it is optimized for memory usage, including settings like swap space and memory limits.

## Mitigation

To mitigate the HighMemoryUsage alert, follow these steps:

1. **Investigate and address root cause**: Identify the underlying cause of high memory usage and address it, whether it be a software bug, a misconfigured service, or a hardware issue.
2. **Optimize system configuration**: Review and optimize system configuration to ensure it is optimized for memory usage, including settings like swap space and memory limits.
3. **Implement memory-saving measures**: Implement measures to reduce memory usage, such as enabling memory compression, adjusting buffer sizes, or optimizing application performance.
4. **Monitor and trend memory usage**: Continuously monitor and trend memory usage to ensure that the mitigation steps have been effective and to catch any potential regressions.

Note: This runbook provides general guidance for diagnosing and mitigating the HighMemoryUsage alert. The specific steps and procedures may vary depending on the specific environment and infrastructure.