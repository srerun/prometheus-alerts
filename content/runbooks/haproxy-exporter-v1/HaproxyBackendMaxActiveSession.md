---
title: HaproxyBackendMaxActiveSession
description: Troubleshooting for alert HaproxyBackendMaxActiveSession
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyBackendMaxActiveSession

HAproxy backend {{ $labels.fqdn }}/{{ $labels.backend }} is reaching session limit (> 80%).

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-exporter-v1.yml" "HaproxyBackendMaxActiveSession" %}}

{{% comment %}}

```yaml
alert: HaproxyBackendMaxActiveSession
expr: ((sum by (backend) (avg_over_time(haproxy_backend_current_sessions[2m]) * 100) / sum by (backend) (avg_over_time(haproxy_backend_limit_sessions[2m])))) > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: HAProxy backend max active session (instance {{ $labels.instance }})
    description: |-
        HAproxy backend {{ $labels.fqdn }}/{{ $labels.backend }} is reaching session limit (> 80%).
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/haproxy-exporter-v1/HaproxyBackendMaxActiveSession.md

```

{{% /comment %}}

</details>


Here is a runbook for the HaproxyBackendMaxActiveSession alert rule:

## Meaning

The HAProxyBackendMaxActiveSession alert is triggered when the average number of active sessions on a HAProxy backend exceeds 80% of the configured session limit. This indicates that the backend is approaching its capacity, which can lead to performance issues, slow response times, and potentially even service unavailability.

## Impact

If left unaddressed, this issue can cause:

* Increased latency and response times for users
* Decreased throughput and productivity
* Potential service unavailability or downtime
* Increased load on the HAProxy instance, leading to resource exhaustion

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the HAProxy logs for any errors or warnings related to the backend in question
2. Verify the current number of active sessions on the backend using the `haproxy_backend_current_sessions` metric
3. Check the configured session limit for the backend using the `haproxy_backend_limit_sessions` metric
4. Investigate the cause of the increased session usage (e.g., increased traffic, slow application responses, etc.)

## Mitigation

To mitigate the issue, follow these steps:

1. **Immediate Action**: Temporarily increase the session limit for the affected backend to avoid service disruption
2. **Root Cause Analysis**: Identify and address the underlying cause of the increased session usage (e.g., optimize application performance, adjust load balancing configuration, etc.)
3. **Short-Term Solution**: Implement traffic shedding or rate limiting to reduce the load on the HAProxy instance
4. **Long-Term Solution**: Upgrade the HAProxy instance or add additional resources to increase its capacity
5. **Monitoring**: Closely monitor the HAProxy backend and adjust the session limit as needed to ensure the service remains stable and performant.