---
title: ThanosBucketReplicateRunLatency
description: Troubleshooting for alert ThanosBucketReplicateRunLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosBucketReplicateRunLatency

Thanos Replicate {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for the replicate operations.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-bucket-replicate.yml" "ThanosBucketReplicateRunLatency" %}}

{{% comment %}}

```yaml
alert: ThanosBucketReplicateRunLatency
expr: (histogram_quantile(0.99, sum by (job) (rate(thanos_replicate_replication_run_duration_seconds_bucket{job=~".*thanos-bucket-replicate.*"}[5m]))) > 20 and  sum by (job) (rate(thanos_replicate_replication_run_duration_seconds_bucket{job=~".*thanos-bucket-replicate.*"}[5m])) > 0)
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Bucket Replicate Run Latency (instance {{ $labels.instance }})
    description: |-
        Thanos Replicate {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for the replicate operations.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-bucket-replicate/ThanosBucketReplicateRunLatency.md

```

{{% /comment %}}

</details>


Here is the runbook for the ThanosBucketReplicateRunLatency alert:

## Meaning

The ThanosBucketReplicateRunLatency alert is triggered when the 99th percentile latency of Thanos replicate operations exceeds 20 seconds in a 5-minute window. This alert indicates that the Thanos replicate job is experiencing high latency, which can lead to data inconsistency and availability issues.

## Impact

The impact of this alert can be significant, as high latency in Thanos replicate operations can:

* Cause data inconsistencies between the primary and replicated buckets
* Lead to slower query performance and increased latency for end-users
* Result in failed replicate operations, leading to data loss or corruption
* Potentially trigger cascading failures in dependent systems

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the Thanos replicate job logs for errors or warnings that may indicate the cause of the high latency.
2. Verify that the Thanos replicate job is running with the expected configuration and resources.
3. Check the underlying storage system for signs of congestion or high latency.
4. Review the Thanos metrics to identify any trends or patterns that may indicate the cause of the high latency.
5. Check the network connectivity and bandwidth between the primary and replicated buckets.

## Mitigation

To mitigate the impact of this alert, follow these steps:

1. Immediately investigate and resolve any underlying issues causing the high latency, such as storage congestion or network connectivity issues.
2. Consider scaling up the resources allocated to the Thanos replicate job to improve its performance.
3. Review the Thanos configuration and optimize it for better performance and latency.
4. Implement additional monitoring and alerting to detect similar issues earlier.
5. Consider implementing redundancy and failover mechanisms to ensure data availability and consistency in case of replicate job failures.