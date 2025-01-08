---
title: HaproxyDown
description: Troubleshooting for alert HaproxyDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyDown

HAProxy down

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-exporter-v1.yml" "HaproxyDown" %}}

{{% comment %}}

```yaml
alert: HaproxyDown
expr: haproxy_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: HAProxy down (instance {{ $labels.instance }})
    description: |-
        HAProxy down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/haproxy-exporter-v1/HaproxyDown.md

```

{{% /comment %}}

</details>


## Meaning

The `HaproxyDown` alert is triggered when the `haproxy_up` metric has a value of 0, indicating that HAProxy is not responding or is down. This alert is critical, as it can cause disruptions to traffic and impact the availability of services.

## Impact

The impact of this alert is high, as it can result in:

* Loss of traffic and revenue
* Downtime and unavailability of services
* Increased latency and errors for users
* Potential security vulnerabilities if HAProxy is not functioning correctly

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the HAProxy logs for errors or issues that may be causing the downtime.
2. Verify that the HAProxy process is running and that there are no configuration issues.
3. Check the system resources (CPU, memory, disk space) to ensure they are within acceptable limits.
4. Investigate any recent changes or updates that may have caused the issue.
5. Use the `haproxy_info` metric to gather more information about the HAProxy instance and its configuration.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the HAProxy service to attempt to recover from the down state.
2. Check and fix any configuration issues that may be causing the downtime.
3. Verify that the system resources are within acceptable limits and take action to free up resources if necessary.
4. Roll back any recent changes or updates that may have caused the issue.
5. If the issue persists, consider escalating to a more senior team member or seeking external support.

Note: For more detailed steps and troubleshooting guides, refer to the linked runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/haproxy-exporter-v1/HaproxyDown.md