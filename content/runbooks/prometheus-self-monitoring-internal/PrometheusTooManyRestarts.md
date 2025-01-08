---
title: PrometheusTooManyRestarts
description: Troubleshooting for alert PrometheusTooManyRestarts
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTooManyRestarts
Prometheus has restarted more than twice in the last 15 minutes. It might be crashlooping.

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTooManyRestarts" %}}

{{% comment %}}

```yaml
alert: PrometheusTooManyRestarts
expr: changes(process_start_time_seconds{job=~"prometheus|pushgateway|alertmanager"}[15m]) > 2
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus too many restarts (instance {{ $labels.instance }})
    description: |-
        Prometheus has restarted more than twice in the last 15 minutes. It might be crashlooping.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTooManyRestarts.md

```

{{% /comment %}}

</details>


## Meaning
This alert triggers when a Prometheus component (Prometheus server, Pushgateway, or Alertmanager) restarts more than twice within a 15-minute window. Frequent restarts can indicate underlying issues such as configuration errors, resource constraints, or software bugs.

## Impact
Frequent restarts of Prometheus components can lead to:
- Gaps in metrics collection, causing incomplete data.
- Delayed or missed alerts, impacting monitoring reliability.
- Potential instability in the monitoring infrastructure.

## Diagnosis
1. **Confirm the Alert:**
   - Check the alert details to identify the affected instance (`{{ $labels.instance }}`) and the specific component (`job`).
   - Note the value of `process_start_time_seconds` changes to confirm frequent restarts.

2. **Review Logs:**
   - Access the logs of the affected component (Prometheus server, Pushgateway, or Alertmanager).
   - Look for recurring errors or exceptions that might indicate the cause of the restarts.

3. **Check Resource Utilization:**
   - Monitor CPU, memory, and disk usage on the affected instance.
   - Look for spikes or exhaustion of resources.

4. **Review Configuration:**
   - Check for recent changes to the component’s configuration files.
   - Validate configurations for syntax errors or invalid settings.

5. **Inspect Dependencies:**
   - Verify network connectivity and DNS resolution for any external dependencies.
   - Ensure that required storage backends or remote write endpoints are operational.

## Mitigation
1. **Resolve Immediate Issues:**
   - If resource constraints are identified, scale up the resources (e.g., increase CPU, memory, or disk).
   - If configuration errors are found, correct and validate the configurations before restarting the component.

2. **Apply Fixes:**
   - Address any software bugs by upgrading to a stable and supported version of the component.
   - Fix network or dependency issues if applicable.

3. **Restart the Component:**
   - After addressing the root cause, restart the affected component manually to stabilize its state.

4. **Monitor Post-Mitigation:**
   - Ensure the component stabilizes and the alert clears.
   - Monitor for recurring issues to validate the effectiveness of the fix.

5. **Document Findings:**
   - Record the root cause, mitigation steps, and any follow-up actions in your incident management system.

## References
- [Prometheus Troubleshooting Guide](https://prometheus.io/docs/operating/troubleshooting/)
- [Alerting Best Practices](https://prometheus.io/docs/practices/alerting/)
- [Runbook Link](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTooManyRestarts.md)

