---
title: ThanosCompactorHalted
description: Troubleshooting for alert ThanosCompactorHalted
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosCompactorHalted

Thanos Compact {{$labels.job}} has failed to run and now is halted.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-compactor.yml" "ThanosCompactorHalted" %}}

{{% comment %}}

```yaml
alert: ThanosCompactorHalted
expr: thanos_compact_halted{job=~".*thanos-compact.*"} == 1
for: 5m
labels:
    severity: warning
annotations:
    summary: Thanos Compactor Halted (instance {{ $labels.instance }})
    description: |-
        Thanos Compact {{$labels.job}} has failed to run and now is halted.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-compactor/ThanosCompactorHalted.md

```

{{% /comment %}}

</details>


Here is a runbook for the ThanosCompactorHalted alert:

## Meaning

The ThanosCompactorHalted alert indicates that the Thanos compactor has failed to run and is currently halted. This means that the compactor is not processing data and may lead to data inconsistencies and potential data loss.

## Impact

The impact of this alert is that the Thanos compactor is not functioning as expected, which can lead to:

* Data inconsistencies and potential data loss
* Increased storage usage due to unprocessed data
* Potential performance degradation of dependent systems

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos compactor logs for errors or exceptions that may indicate the cause of the halt.
2. Verify that the Thanos compactor is properly configured and that all required dependencies are available.
3. Check the system resources (e.g., CPU, memory, disk space) to ensure they are not causing the compactor to fail.
4. Review the Thanos compactor metrics to identify any trends or patterns that may indicate the cause of the halt.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Thanos compactor to attempt to recover from the halted state.
2. Investigate and resolve any underlying issues identified during diagnosis, such as configuration errors or system resource constraints.
3. Verify that the compactor is properly configured and all required dependencies are available.
4. Monitor the Thanos compactor metrics to ensure it is functioning correctly and processing data as expected.

Note: If the issue persists, consider seeking assistance from the Thanos community or a qualified expert.