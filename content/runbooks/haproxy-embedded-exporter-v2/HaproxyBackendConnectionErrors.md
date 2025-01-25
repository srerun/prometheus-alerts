---
title: HaproxyBackendConnectionErrors
description: Troubleshooting for alert HaproxyBackendConnectionErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyBackendConnectionErrors

Too many connection errors to {{ $labels.fqdn }}/{{ $labels.backend }} backend (> 100 req/s). Request throughput may be too high.

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-embedded-exporter-v2.yml" "HaproxyBackendConnectionErrors" %}}

{{% comment %}}

```yaml
alert: HaproxyBackendConnectionErrors
expr: (sum by (proxy) (rate(haproxy_backend_connection_errors_total[1m]))) > 100
for: 1m
labels:
    severity: critical
annotations:
    summary: HAProxy backend connection errors (instance {{ $labels.instance }})
    description: |-
        Too many connection errors to {{ $labels.fqdn }}/{{ $labels.backend }} backend (> 100 req/s). Request throughput may be too high.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v2/HaproxyBackendConnectionErrors.md

```

{{% /comment %}}

</details>


Here is a runbook for the HaproxyBackendConnectionErrors alert:

## Meaning

This alert is triggered when the rate of HAProxy backend connection errors exceeds 100 errors per second for a given backend proxy. This indicates that there are issues with the backend server(s) or the network connectivity, leading to failed connections.

## Impact

* Request throughput may be impacted, resulting in slow or failed responses to users.
* Increased errors can lead to increased latency and decreased overall system performance.
* In extreme cases, this can cause a cascading failure of dependent systems or services.

## Diagnosis

To diagnose the root cause of the issue:

1. Check the HAProxy logs for errors related to the specific backend proxy.
2. Verify the health and connectivity of the backend servers.
3. Check network connectivity and routing between HAProxy and the backend servers.
4. Review recent changes to the HAProxy configuration or backend server deployments.
5. Utilize tools like `haproxy -s` or `haproxyctl` to inspect the current HAProxy configuration and statistics.

## Mitigation

To mitigate the issue:

1. **Immediately**: Investigate and address any issues with the backend servers, such as high CPU usage, memory issues, or disk space constraints.
2. **Short-term**: Adjust the HAProxy configuration to reduce the load on the backend servers, such as by increasing the timeout or reducing the number of connections.
3. **Long-term**: Consider scaling out the backend servers or implementing additional load balancing strategies to handle increased traffic demands.
4. **Monitoring**: Closely monitor the HAProxy backend connection errors and adjust thresholds or alerting mechanisms as needed to ensure timely detection of future issues.

Remember to update the alert annotation with the root cause and mitigation steps taken to resolve the issue.