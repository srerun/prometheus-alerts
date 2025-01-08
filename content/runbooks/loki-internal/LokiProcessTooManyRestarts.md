---
title: LokiProcessTooManyRestarts
description: Troubleshooting for alert LokiProcessTooManyRestarts
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# LokiProcessTooManyRestarts

A loki process had too many restarts (target {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "loki/loki-internal.yml" "LokiProcessTooManyRestarts" %}}

{{% comment %}}

```yaml
alert: LokiProcessTooManyRestarts
expr: changes(process_start_time_seconds{job=~".*loki.*"}[15m]) > 2
for: 0m
labels:
    severity: warning
annotations:
    summary: Loki process too many restarts (instance {{ $labels.instance }})
    description: |-
        A loki process had too many restarts (target {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/loki-internal/LokiProcessTooManyRestarts.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule "LokiProcessTooManyRestarts":

## Meaning

The "LokiProcessTooManyRestarts" alert is triggered when a Loki process restarts more than 2 times within a 15-minute window. This indicates that the Loki process is experiencing instability or issues that are causing it to restart frequently.

## Impact

The impact of this alert is that Loki may not be able to collect and store log data correctly, leading to gaps in log data and potential issues with log-based alerting and monitoring. This can also lead to increased latency and decreased performance in the logging pipeline.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Loki process logs for errors or exceptions that may be causing the restarts.
2. Verify that the Loki configuration is correct and that there are no issues with the underlying infrastructure (e.g. disk space, network connectivity).
3. Check the system metrics (e.g. CPU, memory, disk usage) to see if there are any resource constraints that may be contributing to the restarts.
4. Check the Loki process status and verify that it is running correctly.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Loki process and verify that it is running correctly.
2. Check and update the Loki configuration to ensure that it is correct and up-to-date.
3. Investigate and resolve any underlying infrastructure issues (e.g. disk space, network connectivity).
4. Consider increasing the resources allocated to the Loki process (e.g. increasing the available memory or CPU).
5. Implement additional logging and monitoring to detect and alert on Loki process restarts.

Note: For more detailed steps and troubleshooting guides, please refer to the Loki documentation and the referenced runbook URL.