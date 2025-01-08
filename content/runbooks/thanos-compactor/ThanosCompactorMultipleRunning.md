---
title: ThanosCompactorMultipleRunning
description: Troubleshooting for alert ThanosCompactorMultipleRunning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosCompactorMultipleRunning

No more than one Thanos Compact instance should be running at once. There are {{$value}} instances running.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-compactor.yml" "ThanosCompactorMultipleRunning" %}}

{{% comment %}}

```yaml
alert: ThanosCompactorMultipleRunning
expr: sum by (job) (up{job=~".*thanos-compact.*"}) > 1
for: 5m
labels:
    severity: warning
annotations:
    summary: Thanos Compactor Multiple Running (instance {{ $labels.instance }})
    description: |-
        No more than one Thanos Compact instance should be running at once. There are {{$value}} instances running.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-compactor/ThanosCompactorMultipleRunning.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the ThanosCompactorMultipleRunning alert:

## Meaning

The ThanosCompactorMultipleRunning alert is triggered when multiple Thanos Compactor instances are running simultaneously. This is not expected behavior, as Thanos Compactor is designed to run as a single instance. Running multiple instances can lead to inconsistencies and errors in the compacted data.

## Impact

The impact of running multiple Thanos Compactor instances can be significant, leading to:

* Data inconsistencies and errors
* Increased resource usage and load on the system
* Potential data loss or corruption
* Difficulty in troubleshooting and debugging issues

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus metrics to verify the number of running Thanos Compactor instances.
2. Check the Thanos Compactor logs to see if there are any errors or warnings related to multiple instances running.
3. Check the system configuration and deployment scripts to ensure that only one instance of Thanos Compactor is intended to be running.
4. Check for any automation or scheduling issues that may be causing multiple instances to start unintentionally.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately stop all but one of the running Thanos Compactor instances to prevent further data inconsistencies and errors.
2. Investigate and resolve the root cause of the multiple instances running, such as a configuration or deployment error.
3. Verify that the system configuration and deployment scripts are corrected to ensure only one instance of Thanos Compactor is running.
4. Monitor the system closely to ensure that the issue does not recur.
5. Consider implementing additional monitoring and alerting to detect and prevent similar issues in the future.