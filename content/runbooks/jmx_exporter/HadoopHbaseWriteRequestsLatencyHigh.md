---
title: HadoopHbaseWriteRequestsLatencyHigh
description: Troubleshooting for alert HadoopHbaseWriteRequestsLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopHbaseWriteRequestsLatencyHigh

HBase Write Requests are experiencing high latency.

<details>
  <summary>Alert Rule</summary>

{{% rule "hadoop/jmx_exporter.yml" "HadoopHbaseWriteRequestsLatencyHigh" %}}

{{% comment %}}

```yaml
alert: HadoopHbaseWriteRequestsLatencyHigh
expr: hadoop_hbase_write_requests_latency_seconds > 0.5
for: 10m
labels:
    severity: warning
annotations:
    summary: Hadoop HBase Write Requests Latency High (instance {{ $labels.instance }})
    description: |-
        HBase Write Requests are experiencing high latency.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/jmx_exporter/HadoopHbaseWriteRequestsLatencyHigh.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule:

## Meaning

The `HadoopHbaseWriteRequestsLatencyHigh` alert is triggered when the average latency of HBase write requests exceeds 0.5 seconds over a 10-minute period. This indicates that HBase is experiencing performance issues, which can impact the overall system's ability to write data efficiently.

## Impact

High latency in HBase write requests can have a significant impact on the system, leading to:

* Delayed data ingestion and processing
* Increased queue sizes and memory utilization
* Decreased system throughput and responsiveness
* Potential data loss or corruption

## Diagnosis

To diagnose the root cause of the high latency, follow these steps:

1. Investigate the HBase cluster's resource utilization:
	* Check the CPU, memory, and disk usage of the HBase nodes.
	* Verify that the nodes are not experiencing high resource contention.
2. Examine the HBase write request patterns:
	* Analyze the write request volume and distribution.
	* Identify any sudden spikes or changes in write request patterns.
3. Check the HBase configuration and tuning:
	* Verify that the HBase configuration is optimal for the workload.
	* Check the region server configuration, block cache settings, and other tunable parameters.
4. Investigate potential external factors:
	* Check for any network issues or congestion.
	* Verify that the underlying storage system is performing optimally.

## Mitigation

To mitigate the high latency, consider the following steps:

1. Scale out the HBase cluster:
	* Add more nodes to the cluster to distribute the write load.
	* Ensure that the new nodes are properly configured and balanced.
2. Optimize HBase configuration and tuning:
	* Adjust the region server configuration, block cache settings, and other tunable parameters.
	* Implement efficient compression and caching strategies.
3. Implement write request buffering or queuing:
	* Introduce a queuing mechanism to handle sudden spikes in write requests.
	* Implement a buffering mechanism to absorb bursts of write traffic.
4. Investigate and address underlying storage issues:
	* Optimize the underlying storage system for better performance.
	* Consider upgrading or replacing the storage system if necessary.

Remember to monitor the system closely after applying these mitigations to ensure that the latency issue is resolved and the system is performing optimally.