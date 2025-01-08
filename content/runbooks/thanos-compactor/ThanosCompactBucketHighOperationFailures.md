---
title: ThanosCompactBucketHighOperationFailures
description: Troubleshooting for alert ThanosCompactBucketHighOperationFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosCompactBucketHighOperationFailures

Thanos Compact {{$labels.job}} Bucket is failing to execute {{$value | humanize}}% of operations.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-compactor.yml" "ThanosCompactBucketHighOperationFailures" %}}

{{% comment %}}

```yaml
alert: ThanosCompactBucketHighOperationFailures
expr: (sum by (job) (rate(thanos_objstore_bucket_operation_failures_total{job=~".*thanos-compact.*"}[5m])) / sum by (job) (rate(thanos_objstore_bucket_operations_total{job=~".*thanos-compact.*"}[5m])) * 100 > 5)
for: 15m
labels:
    severity: warning
annotations:
    summary: Thanos Compact Bucket High Operation Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Compact {{$labels.job}} Bucket is failing to execute {{$value | humanize}}% of operations.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-compactor/ThanosCompactBucketHighOperationFailures.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `ThanosCompactBucketHighOperationFailures`:

## Meaning

The `ThanosCompactBucketHighOperationFailures` alert is triggered when the percentage of failed operations in a Thanos compact bucket exceeds 5% over a 5-minute period. This indicates that there is an issue with the compact bucket, which may lead to data inconsistencies or loss.

## Impact

The impact of this alert is that the Thanos compact bucket may not be functioning correctly, leading to potential data inconsistencies or loss. This can have a significant impact on the overall reliability and accuracy of the monitoring system.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos compact bucket logs for any error messages or exceptions that may indicate the cause of the failures.
2. Verify that the compact bucket is properly configured and that there are no issues with the underlying storage system.
3. Check the network connectivity between the compact bucket and the object store to ensure that there are no connectivity issues.
4. Review the Thanos compactor metrics to identify any trends or patterns that may indicate the root cause of the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve any underlying issues with the compact bucket configuration, storage system, or network connectivity.
2. Restart the Thanos compact bucket to ensure that it is functioning correctly.
3. Monitor the Thanos compactor metrics to ensure that the issue is resolved and that the compact bucket is functioning correctly.
4. Consider increasing the logging level for the Thanos compact bucket to gather more detailed information about the issue.

Additional resources:

* Refer to the Thanos documentation for more information on troubleshooting and configuring the compact bucket.
* Review the Thanos compactor metrics to gain a deeper understanding of the system's behavior.

By following these steps, you should be able to diagnose and mitigate the issue, ensuring that the Thanos compact bucket is functioning correctly and that data inconsistencies or loss are minimized.