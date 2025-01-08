---
title: ThanosStoreObjstoreOperationLatencyHigh
description: Troubleshooting for alert ThanosStoreObjstoreOperationLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosStoreObjstoreOperationLatencyHigh

Thanos Store {{$labels.job}} Bucket has a 99th percentile latency of {{$value}} seconds for the bucket operations.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-store.yml" "ThanosStoreObjstoreOperationLatencyHigh" %}}

{{% comment %}}

```yaml
alert: ThanosStoreObjstoreOperationLatencyHigh
expr: (histogram_quantile(0.99, sum by (job, le) (rate(thanos_objstore_bucket_operation_duration_seconds_bucket{job=~".*thanos-store.*"}[5m]))) > 2 and  sum by (job) (rate(thanos_objstore_bucket_operation_duration_seconds_count{job=~".*thanos-store.*"}[5m])) > 0)
for: 10m
labels:
    severity: warning
annotations:
    summary: Thanos Store Objstore Operation Latency High (instance {{ $labels.instance }})
    description: |-
        Thanos Store {{$labels.job}} Bucket has a 99th percentile latency of {{$value}} seconds for the bucket operations.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-store/ThanosStoreObjstoreOperationLatencyHigh.md

```

{{% /comment %}}

</details>


Here is the runbook for the ThanosStoreObjstoreOperationLatencyHigh alert:

## Meaning

The ThanosStoreObjstoreOperationLatencyHigh alert is triggered when the 99th percentile latency of Thanos Store objstore bucket operations exceeds 2 seconds and there are ongoing operations. This alert indicates that the Thanos Store is experiencing high latency when performing objstore bucket operations, which may impact the overall performance and reliability of the system.

## Impact

The high latency of objstore bucket operations can have a significant impact on the system, including:

* Delayed query responses and slower data ingest
* Increased load on the system, leading to potential crashes or failures
* Decreased system reliability and availability
* Potential data loss or corruption

## Diagnosis

To diagnose the root cause of the high latency, follow these steps:

1. Check the Thanos Store logs for any errors or warnings related to objstore bucket operations.
2. Investigate the current load and resource utilization of the Thanos Store nodes.
3. Check the network connectivity and latency between the Thanos Store nodes and the objstore.
4. Verify that the objstore is properly configured and functioning correctly.
5. Check the objstore bucket operation metrics to identify any patterns or anomalies.

## Mitigation

To mitigate the high latency of objstore bucket operations, follow these steps:

1. Scale up the Thanos Store nodes to increase resources and handle the load.
2. Optimize the objstore configuration for better performance.
3. Implement caching or other optimization techniques to reduce the load on the objstore.
4. Check for any software or configuration issues and apply updates or patches as needed.
5. Consider implementing a load balancing or queuing mechanism to handle high volumes of objstore bucket operations.

By following these steps, you should be able to identify and resolve the root cause of the high latency and restore normal operation of the Thanos Store.