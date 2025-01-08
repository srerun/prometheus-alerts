---
title: ThanosSidecarBucketOperationsFailed
description: Troubleshooting for alert ThanosSidecarBucketOperationsFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosSidecarBucketOperationsFailed

Thanos Sidecar {{$labels.instance}} bucket operations are failing

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-sidecar.yml" "ThanosSidecarBucketOperationsFailed" %}}

{{% comment %}}

```yaml
alert: ThanosSidecarBucketOperationsFailed
expr: sum by (job, instance) (rate(thanos_objstore_bucket_operation_failures_total{job=~".*thanos-sidecar.*"}[5m])) > 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Sidecar Bucket Operations Failed (instance {{ $labels.instance }})
    description: |-
        Thanos Sidecar {{$labels.instance}} bucket operations are failing
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-sidecar/ThanosSidecarBucketOperationsFailed.md

```

{{% /comment %}}

</details>


## Meaning

The `ThanosSidecarBucketOperationsFailed` alert is triggered when the Thanos Sidecar instance encounters failures while performing bucket operations. This alert indicates that the Thanos Sidecar is experiencing issues while interacting with the object store, which can lead to data loss or inconsistencies.

## Impact

The impact of this alert is high, as it can result in:

* Data loss or corruption
* Inconsistent metrics and data
* Unavailability of the Thanos Sidecar instance
* Potential cascading failures in dependent systems

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Sidecar logs for error messages related to bucket operations
2. Verify the object store connection and credentials
3. Check the instance's resource utilization (CPU, memory, disk space) to ensure it is not overloaded
4. Investigate if there are any network connectivity issues between the Thanos Sidecar instance and the object store
5. Review the Thanos Sidecar configuration to ensure it is correctly set up

## Mitigation

To mitigate the issue, follow these steps:

1.Restart the Thanos Sidecar instance to attempt to recover from the failure
2. Investigate and resolve any underlying issues with the object store connection or credentials
3. Increase the resources (CPU, memory, disk space) allocated to the Thanos Sidecar instance if necessary
4. Implement additional logging and monitoring to detect similar issues in the future
5. Consider configuring Thanos Sidecar to use a more robust object store or to use a fallback storage solution.