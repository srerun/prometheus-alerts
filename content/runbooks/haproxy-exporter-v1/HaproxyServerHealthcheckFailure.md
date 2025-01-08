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

{{% rule "haproxy/haproxy-exporter-v1.yml" "HaproxyServerHealthcheckFailure" %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/haproxy-exporter-v1/HaproxyServerHealthcheckFailure.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule `HaproxyServerHealthcheckFailure`:

## Meaning

The `HaproxyServerHealthcheckFailure` alert is triggered when the number of healthcheck failures for a HAProxy server increases within a 1-minute window. This indicates that one or more servers behind the load balancer are not responding correctly to health checks, which may lead to reduced availability or incorrect routing of traffic.

## Impact

The impact of this alert is that some servers may be taken out of rotation by HAProxy, leading to:

* Reduced capacity and potential service disruption
* Incorrect traffic routing, which may cause errors or timeouts for clients
* Increased latency or connection failures for clients accessing the affected servers

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the HAProxy logs for errors related to healthcheck failures
2. Verify the healthcheck configuration for the affected server(s)
3. Check the server logs for any errors or issues that may be causing the healthcheck failures
4. Verify that the server(s) are properly configured and running correctly
5. Check for any network connectivity issues between the HAProxy instance and the affected server(s)

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve the underlying cause of the healthcheck failures
2. Verify that the affected server(s) are properly configured and running correctly
3. If necessary, update the HAProxy configuration to temporarily remove the affected server(s) from rotation
4. Implement additional logging or monitoring to detect similar issues in the future
5. Consider implementing automated remediation steps, such as restarting the affected server(s) or running a healthcheck script to verify server health.