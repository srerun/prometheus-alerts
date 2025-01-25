---
title: HaproxyHasNoAliveBackends
description: Troubleshooting for alert HaproxyHasNoAliveBackends
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyHasNoAliveBackends

HAProxy has no alive active or backup backends for {{ $labels.proxy }}

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-embedded-exporter-v2.yml" "HaproxyHasNoAliveBackends" %}}

{{% comment %}}

```yaml
alert: HaproxyHasNoAliveBackends
expr: haproxy_backend_active_servers + haproxy_backend_backup_servers == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: HAproxy has no alive backends (instance {{ $labels.instance }})
    description: |-
        HAProxy has no alive active or backup backends for {{ $labels.proxy }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v2/HaproxyHasNoAliveBackends.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `HaproxyHasNoAliveBackends`:

## Meaning

The `HaproxyHasNoAliveBackends` alert is triggered when HAProxy has no alive active or backup backends for a specific proxy instance. This means that HAProxy is unable to forward incoming requests to any available backend servers, causing a complete outage of the service.

## Impact

The impact of this alert is critical, as it leads to a complete loss of service availability. Without any available backend servers, HAProxy is unable to process incoming requests, resulting in:

* Downtime for users
* Loss of revenue and business impact
* Increased latency and error rates

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the HAProxy logs for any error messages related to backend server connectivity or availability.
2. Verify the configuration of the HAProxy instance and ensure that the backend servers are correctly defined and configured.
3. Check the status of the backend servers to ensure they are running and accepting incoming requests.
4. Verify that there are no network connectivity issues between HAProxy and the backend servers.
5. Check the Prometheus metrics for any indication of high latency or error rates that may have contributed to the lack of available backend servers.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and address any underlying issues causing the unavailability of the backend servers, such as:
	* Server crashes or restarts
	* Network connectivity issues
	* Configuration errors
2. Immediately restart or redeploy any unavailable backend servers to restore service availability.
3. Perform a rolling update or rolling restart of the HAProxy instance to ensure that it can reconnect to the available backend servers.
4. Monitor the HAProxy and backend server metrics closely to ensure that the issue does not reoccur.
5. Consider implementing additional redundancy and failover mechanisms to minimize the impact of future outages.