---
title: ThanosCompactHasNotRun
description: Troubleshooting for alert ThanosCompactHasNotRun
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosCompactHasNotRun

Thanos Compact {{$labels.job}} has not uploaded anything for 24 hours.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-compactor.yml" "ThanosCompactHasNotRun" %}}

{{% comment %}}

```yaml
alert: ThanosCompactHasNotRun
expr: (time() - max by (job) (max_over_time(thanos_objstore_bucket_last_successful_upload_time{job=~".*thanos-compact.*"}[24h]))) / 60 / 60 > 24
for: 0m
labels:
    severity: warning
annotations:
    summary: Thanos Compact Has Not Run (instance {{ $labels.instance }})
    description: |-
        Thanos Compact {{$labels.job}} has not uploaded anything for 24 hours.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-compactor/ThanosCompactHasNotRun.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "ThanosCompactHasNotRun":

## Meaning

The "ThanosCompactHasNotRun" alert is triggered when a Thanos Compact instance has not uploaded data to the object store for more than 24 hours. This indicates that the compaction process is not functioning correctly, which can lead to data inconsistencies and potential data loss.

## Impact

If this alert is not addressed, it can lead to:

* Data inconsistencies and potential data loss
* Increased storage usage due to uncompact data
* Performance degradation of the Thanos system
* Inability to query or retrieve data from the object store

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Compact instance logs for errors or exceptions related to the upload process.
2. Verify that the Thanos Compact instance is running and has not crashed or terminated abnormally.
3. Check the object store configuration and credentials to ensure they are correct and up-to-date.
4. Verify that the network connectivity between the Thanos Compact instance and the object store is stable and functional.
5. Check the system resources (CPU, memory, disk space) to ensure they are not overutilized or exhausted.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Thanos Compact instance to ensure it is running correctly.
2. Check and update the object store configuration and credentials as necessary.
3. Verify that the network connectivity between the Thanos Compact instance and the object store is stable and functional.
4. Check the system resources (CPU, memory, disk space) and allocate additional resources if necessary.
5. Manually trigger a compaction upload to the object store to ensure the process is functioning correctly.

Note: If the issue persists after following these steps, it may be necessary to escalate the issue to a senior engineer or expert in Thanos Compact for further assistance.