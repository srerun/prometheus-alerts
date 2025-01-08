---
title: HostOutOfMemory
description: Troubleshooting for alert HostOutOfMemory
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostOutOfMemory

Node memory is filling up (< 10% left)

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostOutOfMemory" %}}

{{% comment %}}

```yaml
alert: HostOutOfMemory
expr: (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes * 100 < 10) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host out of memory (instance {{ $labels.instance }})
    description: |-
        Node memory is filling up (< 10% left)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostOutOfMemory.md

```

{{% /comment %}}

</details>


Here is a runbook for the HostOutOfMemory alert:

## Meaning

The HostOutOfMemory alert is triggered when a node's available memory falls below 10% of its total memory capacity. This indicates that the node is running low on memory resources, which can cause performance issues, slow down applications, and even lead to crashes or errors.

## Impact

If left unaddressed, low memory can have significant impacts on the system and applications running on it. Some potential consequences include:

* Slow performance and response times
* Increased latency and timeouts
* Application crashes or errors
* Data loss or corruption
* Increased risk of security breaches

## Diagnosis

To diagnose the root cause of the HostOutOfMemory alert, follow these steps:

1. Check the node's memory usage patterns:
	* Use Prometheus queries to inspect the node's memory usage over time.
	* Look for trends, spikes, or anomalies in memory consumption.
2. Identify resource-intensive processes:
	* Use tools like `top`, `htop`, or `ps` to identify processes consuming high amounts of memory.
	* Check for any memory leaks or inefficient memory allocation.
3. Review system configuration and resource allocation:
	* Verify that the node has sufficient memory resources allocated.
	* Check for any misconfigured resources, such as swap space or memory limits.

## Mitigation

To mitigate the HostOutOfMemory alert, follow these steps:

1. Increase node memory capacity:
	* Add more physical RAM to the node, if possible.
	* Consider migrating to a more powerful instance type or virtual machine.
2. Optimize resource-intensive processes:
	* Identify and fix memory leaks or inefficient memory allocation in applications.
	* Optimize configuration and tuning for resource-intensive processes.
3. Implement memory-saving measures:
	* Consider implementing caching or buffering mechanisms to reduce memory usage.
	* Implement memory-saving features, such as compression or data deduplication.
4. Monitor and prevent future occurrences:
	* Set up regular memory usage checks and alerts.
	* Implement automated scripts to detect and respond to memory usage anomalies.

Remember to update the `runbook` annotation in the Prometheus alert rule to point to this runbook, so that operators can easily access these instructions when the alert is triggered.