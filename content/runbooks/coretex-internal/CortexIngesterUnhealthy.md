---
title: CortexIngesterUnhealthy
description: Troubleshooting for alert CortexIngesterUnhealthy
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CortexIngesterUnhealthy

Cortex has an unhealthy ingester

<details>
  <summary>Alert Rule</summary>

{{% rule "cortex/coretex-internal.yml" "CortexIngesterUnhealthy" %}}

{{% comment %}}

```yaml
alert: CortexIngesterUnhealthy
expr: cortex_ring_members{state="Unhealthy", name="ingester"} > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Cortex ingester unhealthy (instance {{ $labels.instance }})
    description: |-
        Cortex has an unhealthy ingester
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/coretex-internal/CortexIngesterUnhealthy.md

```

{{% /comment %}}

</details>


Here is a runbook for the CortexIngesterUnhealthy alert rule:

## Meaning

The CortexIngesterUnhealthy alert is triggered when one or more Cortex ingester instances are reported as unhealthy by the Cortex ring. This means that the ingester is not functioning correctly and may be unable to process requests or send data to storage.

## Impact

An unhealthy ingester can have a significant impact on the overall health and functionality of the Cortex system. This can lead to:

* Data loss or incomplete data
* Increased latency or timeouts for requests
* Decreased system performance and reliability
* Inaccurate or incomplete metrics and alerting

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cortex ingester logs for errors or exceptions that may indicate the cause of the unhealthy state.
2. Verify that the ingester instance is running and that the process is not stuck or dead.
3. Check the system resources (CPU, memory, disk space) to ensure they are within acceptable limits.
4. Investigate if there are any network connectivity issues or firewall rules blocking communication between the ingester and other Cortex components.
5. Review the Cortex ring configuration to ensure it is correctly configured and that the ingester is properly registered.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the unhealthy ingester instance to attempt to recover it.
2. If the issue persists, try to redeploy the ingester instance with a new version or configuration.
3. If the issue is related to system resources, consider scaling up the instance or adding more resources to the system.
4. If the issue is related to network connectivity, investigate and resolve any firewall or network configuration issues.
5. If none of the above steps resolve the issue, consider reaching out to the Cortex support team or filing a bug report.

Remember to always follow proper incident response procedures and to investigate and resolve the root cause of the issue to prevent it from happening again in the future.