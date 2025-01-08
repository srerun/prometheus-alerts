---
title: HostMemoryUnderMemoryPressure
description: Troubleshooting for alert HostMemoryUnderMemoryPressure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostMemoryUnderMemoryPressure

The node is under heavy memory pressure. High rate of major page faults

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostMemoryUnderMemoryPressure" %}}

{{% comment %}}

```yaml
alert: HostMemoryUnderMemoryPressure
expr: (rate(node_vmstat_pgmajfault[1m]) > 1000) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host memory under memory pressure (instance {{ $labels.instance }})
    description: |-
        The node is under heavy memory pressure. High rate of major page faults
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostMemoryUnderMemoryPressure.md

```

{{% /comment %}}

</details>


## Meaning

The HostMemoryUnderMemoryPressure alert is triggered when a node is experiencing high memory pressure, indicated by a high rate of major page faults. This can lead to performance degradation, slowdowns, and even crashes if left unchecked.

## Impact

* Performance degradation: High memory pressure can cause the node to slow down, leading to decreased performance and responsiveness.
* Increased latency: As the system struggles to allocate memory, requests may take longer to process, leading to increased latency.
* Potential crashes: If memory pressure is severe, the node may crash or become unresponsive, leading to downtime and potential data loss.

## Diagnosis

To diagnose the root cause of the HostMemoryUnderMemoryPressure alert, follow these steps:

1. Investigate the node's memory usage: Check the node's memory usage patterns to identify any unusual spikes or trends.
2. Review system logs: Analyze system logs to identify any errors or warnings related to memory allocation or page faults.
3. Check running processes: Identify any memory-intensive processes that may be contributing to the high memory pressure.
4. Verify node configuration: Review the node's configuration to ensure that it is properly sized and configured for the workload.

## Mitigation

To mitigate the HostMemoryUnderMemoryPressure alert, follow these steps:

1. Reduce memory allocation: Identify and optimize memory-intensive processes or applications to reduce their memory footprint.
2. Increase node resources: Consider upgrading the node's hardware or adding more resources (e.g., RAM) to alleviate memory pressure.
3. Implement memory-efficient practices: Implement best practices for memory management, such as using caching mechanisms or optimizing data structures.
4. Monitor and adjust: Continuously monitor the node's memory usage and adjust resource allocation or configuration as needed to prevent future occurrences of memory pressure.

Additional resources:

* For more information on diagnosing and mitigating memory pressure, refer to the [Node Exporter documentation](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostMemoryUnderMemoryPressure.md).
* Consider implementing memory-related metrics and dashboards to provide better visibility into node memory usage.