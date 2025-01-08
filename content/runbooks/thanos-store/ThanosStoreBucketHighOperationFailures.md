---
title: ThanosStoreBucketHighOperationFailures
description: Troubleshooting for alert ThanosStoreBucketHighOperationFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosStoreBucketHighOperationFailures

Thanos Store {{$labels.job}} Bucket is failing to execute {{$value | humanize}}% of operations.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-store.yml" "ThanosStoreBucketHighOperationFailures" %}}

{{% comment %}}

```yaml
alert: ThanosStoreBucketHighOperationFailures
expr: (sum by (job) (rate(thanos_objstore_bucket_operation_failures_total{job=~".*thanos-store.*"}[5m])) / sum by (job) (rate(thanos_objstore_bucket_operations_total{job=~".*thanos-store.*"}[5m])) * 100 > 5)
for: 15m
labels:
    severity: warning
annotations:
    summary: Thanos Store Bucket High Operation Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Store {{$labels.job}} Bucket is failing to execute {{$value | humanize}}% of operations.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-store/ThanosStoreBucketHighOperationFailures.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule:

## Meaning

The ThanosStoreBucketHighOperationFailures alert is triggered when the percentage of failed operations in a Thanos Store bucket exceeds 5% over a 5-minute window. This indicates that the bucket is experiencing an unusual number of failures, which may impact the availability and reliability of the system.

## Impact

* High operation failures in a Thanos Store bucket can lead to data loss, corruption, or inconsistency.
* This may cause issues with downstream systems that rely on the data stored in the bucket.
* Prolonged failures can lead to a buildup of undelivered data, causing further problems when the bucket recovers.

## Diagnosis

* Check the Thanos Store bucket logs for error messages indicating the cause of the failures.
* Investigate the network connectivity and storage system health to rule out infrastructure-related issues.
* Verify that the bucket is properly configured and that the storage capacity is sufficient.
* Check the system load and resource utilization to identify any resource constraints.

## Mitigation

* Investigate and address the root cause of the operation failures, such as network issues, storage system problems, or configuration errors.
* Temporarily increase the storage capacity or add more resources to alleviate any resource constraints.
* Implement retries or fall-back mechanisms to minimize the impact of failures on downstream systems.
* Consider scaling up or out the Thanos Store cluster to improve its resilience and availability.

Note: This runbook is a general guideline and may need to be tailored to your specific environment and use case.