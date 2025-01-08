---
title: NetdataHighMemoryUsage
description: Troubleshooting for alert NetdataHighMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataHighMemoryUsage

Netdata high memory usage (> 80%)

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "NetdataHighMemoryUsage" %}}

{{% comment %}}

```yaml
alert: NetdataHighMemoryUsage
expr: 100 / netdata_system_ram_MiB_average * netdata_system_ram_MiB_average{dimension=~"free|cached"} < 20
for: 5m
labels:
    severity: warning
annotations:
    summary: Netdata high memory usage (instance {{ $labels.instance }})
    description: |-
        Netdata high memory usage (> 80%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/NetdataHighMemoryUsage.md

```

{{% /comment %}}

</details>


Here is a runbook for the NetdataHighMemoryUsage alert:

## Meaning

The NetdataHighMemoryUsage alert is triggered when the available memory (free and cached) on a Netdata instance falls below 20% of the total system RAM. This alert indicates that the Netdata instance is experiencing high memory usage, which can lead to performance issues, slow response times, and even crashes.

## Impact

High memory usage on a Netdata instance can have significant impacts on the system and its users:

* Slow response times: High memory usage can cause Netdata to respond slowly to requests, leading to delays and frustration for users.
* Performance issues: Insufficient memory can cause Netdata to fail to process data in a timely manner, leading to errors and inconsistencies.
* Crashes: In extreme cases, high memory usage can cause Netdata to crash, resulting in data loss and system downtime.

## Diagnosis

To diagnose the root cause of the NetdataHighMemoryUsage alert, follow these steps:

1. Check the Netdata dashboard for signs of high memory usage, such as high RAM usage graphs or error messages.
2. Review the Netdata logs to identify any errors or warnings related to memory usage.
3. Check for any resource-intensive processes or plugins running on the Netdata instance that may be contributing to high memory usage.
4. Verify that the Netdata instance is configured to use the correct amount of memory for its workload.

## Mitigation

To mitigate the NetdataHighMemoryUsage alert, follow these steps:

1. Restart the Netdata service to free up memory and restart any resource-intensive processes.
2. Identify and terminate any unnecessary or resource-intensive processes or plugins running on the Netdata instance.
3. Adjust the Netdata configuration to optimize memory usage, such as by reducing the sampling rate or disabling unnecessary plugins.
4. Consider upgrading the Netdata instance to a more powerful machine with more RAM to handle the workload.
5. Monitor the Netdata instance closely to ensure that memory usage returns to normal levels and does not continue to increase over time.