---
title: HaproxyServerHealthcheckFailure
description: Troubleshooting for alert HaproxyServerHealthcheckFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyServerHealthcheckFailure

Some server healthcheck are failing on {{ $labels.server }}

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/embedded-exporter-v2.yml" "HaproxyServerHealthcheckFailure" %}}

{{% comment %}}

```yaml
alert: HaproxyServerHealthcheckFailure
expr: increase(haproxy_server_check_failures_total[1m]) > 0
for: 1m
labels:
    severity: warning
annotations:
    summary: HAProxy server healthcheck failure (instance {{ $labels.instance }})
    description: |-
        Some server healthcheck are failing on {{ $labels.server }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v2/HaproxyServerHealthcheckFailure.md

```

{{% /comment %}}

</details>


Here is the runbook for the HaproxyServerHealthcheckFailure alert:

## Meaning

The HaproxyServerHealthcheckFailure alert is triggered when the HAProxy server healthcheck fails, indicating that one or more servers are not responding to healthcheck requests. This alert is raised when the rate of healthcheck failures increases over a 1-minute period, indicating a potential issue with the server or the HAProxy configuration.

## Impact

The impact of this alert can be significant, as it may indicate that one or more servers are not functioning correctly, which can lead to:

* Reduced system availability
* Increased error rates
* Degraded performance
* Potential data loss or corruption

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the HAProxy logs for errors or warnings related to the healthcheck failures.
2. Verify that the servers are configured correctly and are responding to healthcheck requests.
3. Check the server metrics to identify any potential issues, such as high CPU usage or memory consumption.
4. Review the HAProxy configuration to ensure that it is correct and up-to-date.
5. Perform a manual healthcheck to validate that the servers are responding correctly.

## Mitigation

To mitigate this issue, follow these steps:

1. Investigate and resolve any issues with the servers, such as restarting the server or applying configuration changes.
2. Update the HAProxy configuration to ensure it is correct and up-to-date.
3. Perform a rolling restart of the HAProxy servers to ensure that all servers are updated and functioning correctly.
4. Verify that the healthcheck failures have decreased or resolved after taking the above steps.
5. Consider implementing additional monitoring and alerting to detect potential issues with the servers or HAProxy configuration.