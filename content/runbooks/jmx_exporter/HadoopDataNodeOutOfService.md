---
title: HadoopDataNodeOutOfService
description: Troubleshooting for alert HadoopDataNodeOutOfService
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopDataNodeOutOfService

The Hadoop DataNode is not sending heartbeats.

<details>
  <summary>Alert Rule</summary>

{{% rule "hadoop/jmx_exporter.yml" "HadoopDataNodeOutOfService" %}}

{{% comment %}}

```yaml
alert: HadoopDataNodeOutOfService
expr: hadoop_datanode_last_heartbeat == 0
for: 10m
labels:
    severity: warning
annotations:
    summary: Hadoop Data Node Out Of Service (instance {{ $labels.instance }})
    description: |-
        The Hadoop DataNode is not sending heartbeats.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/jmx_exporter/HadoopDataNodeOutOfService.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `HadoopDataNodeOutOfService`:

## Meaning

The `HadoopDataNodeOutOfService` alert is triggered when a Hadoop DataNode fails to send heartbeats to the Hadoop cluster for an extended period (10 minutes). This indicates that the DataNode is not operating properly and may be unavailable for data storage and retrieval.

## Impact

The impact of a DataNode being out of service can be significant, leading to:

* Reduced data availability and accessibility
* Increased load on other DataNodes in the cluster
* Potential data loss or corruption
* Slower query performance and increased latency

## Diagnosis

To diagnose the issue, perform the following steps:

1. Check the DataNode logs for any errors or exceptions that may indicate the cause of the failure.
2. Verify that the DataNode process is running and configured correctly.
3. Check the network connectivity between the DataNode and the Namenode to ensure that there are no communication issues.
4. Check the disk usage and available space on the DataNode to ensure that it is not running out of space.
5. Check the JMX metrics for any indications of high CPU usage, memory issues, or other resource constraints.

## Mitigation

To mitigate the issue, perform the following steps:

1. Restart the DataNode process to attempt to recover from any temporary issues.
2. Check and resolve any underlying hardware or software issues that may be causing the failure.
3. Verify that the DataNode is properly configured and registered with the Namenode.
4. Consider adding additional DataNodes to the cluster to increase redundancy and availability.
5. Implement monitoring and alerting to detect similar issues earlier in the future.

Note: The above runbook is a general guideline and may need to be tailored to your specific Hadoop cluster configuration and deployment.