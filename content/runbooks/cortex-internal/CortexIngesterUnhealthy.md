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

{{% rule "cortex/cortex-internal.yml" "CortexIngesterUnhealthy" %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/cortex-internal/CortexIngesterUnhealthy.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the CortexIngesterUnhealthy alert:

## Meaning

The CortexIngesterUnhealthy alert is triggered when one or more Cortex ingesters are reported as unhealthy by the Cortex ring. This indicates a critical issue with the Cortex cluster, as unhealthy ingesters can lead to data loss and errors in the system.

## Impact

The impact of an unhealthy ingester can be severe, leading to:

* Data loss or corruption
* Errors in query results
* Increased latency and timeouts
* Reduced system reliability and availability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cortex ring membership to identify the unhealthy ingester(s)
2. Review the ingester logs to determine the cause of the unhealthy state
3. Check the system metrics (e.g. CPU, memory, disk usage) to identify any resource issues
4. Verify that the ingester is properly configured and running with the correct version
5. Check for any network connectivity issues between the ingester and other Cortex components

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the unhealthy ingester(s) to attempt to recover
2. If the issue persists, investigate and resolve any underlying causes (e.g. resource issues, configuration errors)
3. If necessary, replace the unhealthy ingester with a new instance
4. Verify that the Cortex cluster is functioning correctly and data is being ingested properly
5. Monitor the system closely to ensure the issue does not recur

Note: This is just a sample runbook, and you should tailor it to your specific environment and requirements.