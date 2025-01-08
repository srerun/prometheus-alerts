---
title: NatsTooManyErrors
description: Troubleshooting for alert NatsTooManyErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsTooManyErrors

NATS server has encountered errors in the last 5 minutes

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsTooManyErrors" %}}

{{% comment %}}

```yaml
alert: NatsTooManyErrors
expr: increase(gnatsd_varz_jetstream_stats_api_errors[5m]) > 0
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats too many errors (instance {{ $labels.instance }})
    description: |-
        NATS server has encountered errors in the last 5 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsTooManyErrors.md

```

{{% /comment %}}

</details>


Here is the runbook for the NatsTooManyErrors alert rule:

## Meaning

The NatsTooManyErrors alert is triggered when the NATS server experiences an increasing number of errors in the Jetstream API within a 5-minute window. This alert indicates that the NATS server is not functioning as expected, and errors are occurring that may impact the system's overall reliability and performance.

## Impact

* Errors in the Jetstream API may cause data loss or corruption
* System reliability and performance may be degraded
* Applications that rely on NATS may experience errors or failures
* The overall system may become unstable or unresponsive

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the NATS server logs for error messages related to the Jetstream API
2. Verify that the NATS server is configured correctly and that there are no issues with the underlying infrastructure
3. Check the system's resource utilization (CPU, memory, disk space) to ensure that it is within normal operating ranges
4. Review the NATS server's configuration and verify that it is correctly configured to handle the current workload
5. Check for any recent changes or updates that may have caused the issue

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the NATS server to clear out any temporary errors
2. Check and adjust the NATS server's configuration to ensure it is correctly set up to handle the current workload
3. Verify that the underlying infrastructure is functioning correctly and that there are no issues with the system's resources
4. Check for any software updates or patches that may resolve the issue
5. Consider increasing the resources available to the NATS server or distributing the workload to multiple instances to improve reliability and performance.

Note: This is just a sample runbook, and the specific diagnosis and mitigation steps may vary depending on your specific environment and setup.