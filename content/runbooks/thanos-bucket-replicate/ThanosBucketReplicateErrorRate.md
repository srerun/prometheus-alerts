---
title: ThanosBucketReplicateErrorRate
description: Troubleshooting for alert ThanosBucketReplicateErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosBucketReplicateErrorRate

Thanos Replicate is failing to run, {{$value | humanize}}% of attempts failed.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-bucket-replicate.yml" "ThanosBucketReplicateErrorRate" %}}

{{% comment %}}

```yaml
alert: ThanosBucketReplicateErrorRate
expr: (sum by (job) (rate(thanos_replicate_replication_runs_total{result="error", job=~".*thanos-bucket-replicate.*"}[5m]))/ on (job) group_left  sum by (job) (rate(thanos_replicate_replication_runs_total{job=~".*thanos-bucket-replicate.*"}[5m]))) * 100 >= 10
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Bucket Replicate Error Rate (instance {{ $labels.instance }})
    description: |-
        Thanos Replicate is failing to run, {{$value | humanize}}% of attempts failed.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-bucket-replicate/ThanosBucketReplicateErrorRate.md

```

{{% /comment %}}

</details>


## Meaning

The ThanosBucketReplicateErrorRate alert is triggered when the error rate of Thanos bucket replication runs exceeds 10% over a 5-minute period. This indicates that Thanos Replicate is experiencing issues while trying to replicate data from one storage bucket to another.

## Impact

A high error rate in Thanos bucket replication can lead to:

* Data inconsistencies between buckets
* Increased storage costs due to duplicated or redundant data
* Delayed or failed data backups
* Potential data loss in case of bucket failures or corruption

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Replicate logs for errors related to bucket replication.
2. Verify the bucket credentials and permissions to ensure they are correct and up-to-date.
3. Check the network connectivity and firewall rules to ensure they are not blocking the replication process.
4. Investigate any recent changes to the Thanos Replicate configuration or bucket settings.
5. Check the disk space and storage capacity of the affected buckets to ensure they are not full or near capacity.

## Mitigation

To mitigate the issue, take the following steps:

1. Check the Thanos Replicate configuration and adjust the retry policy or backoff strategy to handle temporary errors.
2. Verify that the bucket credentials and permissions are correct and update them if necessary.
3. Ensure network connectivity and firewall rules are correctly configured to allow replication traffic.
4. Consider increasing the disk space or storage capacity of the affected buckets to prevent replication failures.
5. Implement monitoring and alerting for Thanos Replicate to catch issues early and prevent data inconsistencies.

Additional resources:

* Refer to the Thanos documentation for troubleshooting replication issues: [Thanos Replicate Troubleshooting](https://thanos.io/tip/thanos/docs/components/replicate.md#troubleshooting)
* Review the Thanos Replicate configuration and adjust settings as needed: [Thanos Replicate Configuration](https://thanos.io/tip/thanos/docs/components/replicate.md#configuration)