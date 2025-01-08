---
title: OpenebsUsedPoolCapacity
description: Troubleshooting for alert OpenebsUsedPoolCapacity
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OpenebsUsedPoolCapacity

OpenEBS Pool use more than 80% of his capacity

<details>
  <summary>Alert Rule</summary>

{{% rule "openebs/openebs-internal.yml" "OpenebsUsedPoolCapacity" %}}

{{% comment %}}

```yaml
alert: OpenebsUsedPoolCapacity
expr: openebs_used_pool_capacity_percent > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: OpenEBS used pool capacity (instance {{ $labels.instance }})
    description: |-
        OpenEBS Pool use more than 80% of his capacity
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/openebs-internal/OpenebsUsedPoolCapacity.md

```

{{% /comment %}}

</details>


Here is a runbook for the OpenEBSUsedPoolCapacity alert rule:

## Meaning

The OpenEBSUsedPoolCapacity alert is triggered when the used pool capacity of an OpenEBS instance exceeds 80%. This indicates that the storage capacity of the pool is almost full, which can lead to performance issues, data loss, or even complete system unavailability.

## Impact

The impact of this alert can be significant, as it can affect the availability and performance of the entire OpenEBS cluster. If the pool capacity continues to increase, it can lead to:

* Slower write performance
* Increased risk of data loss
* Unavailability of the OpenEBS cluster
* Downtime of dependent applications and services

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the OpenEBS dashboard or CLI to identify the pool that is experiencing high utilization.
2. Verify the current used capacity and available capacity of the pool.
3. Check the write performance and IOPS of the pool to determine if it is affected.
4. Review the OpenEBS logs to identify any errors or warnings related to the pool.
5. Check the application and service logs to identify if there are any issues related to data writes or reads.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the root cause of the increased usage: Check if there is an unexpected surge in data writes or if there is an underlying issue with the application or service using the OpenEBS cluster.
2. Free up capacity: Consider deleting unnecessary data, compressing data, or evacuating data to a different pool or storage system.
3. Add more capacity: Expand the OpenEBS cluster by adding more nodes or disks to increase the available capacity.
4. Implement capacity planning: Review the OpenEBS usage patterns and implement a capacity planning strategy to prevent similar issues in the future.
5. Monitor and adjust: Continuously monitor the OpenEBS pool capacity and adjust the mitigation strategy as needed.