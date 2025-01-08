---
title: LokiRequestPanic
description: Troubleshooting for alert LokiRequestPanic
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# LokiRequestPanic

The {{ $labels.job }} is experiencing {{ printf "%.2f" $value }}% increase of panics

<details>
  <summary>Alert Rule</summary>

{{% rule "loki/loki-internal.yml" "LokiRequestPanic" %}}

{{% comment %}}

```yaml
alert: LokiRequestPanic
expr: sum(increase(loki_panic_total[10m])) by (namespace, job) > 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Loki request panic (instance {{ $labels.instance }})
    description: |-
        The {{ $labels.job }} is experiencing {{ printf "%.2f" $value }}% increase of panics
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/loki-internal/LokiRequestPanic.md

```

{{% /comment %}}

</details>


Here is a runbook for the LokiRequestPanic alert rule:

## Meaning

The LokiRequestPanic alert is triggered when there is a significant increase in panics within Loki requests within a 10-minute window, indicating a critical issue with the Loki service.

## Impact

The impact of this alert is high, as it can result in:

* Loki requests failing, leading to incomplete or missing log data
* Increased latency and errors in the logging pipeline
* Potential cascading failures in dependent systems

## Diagnosis

To diagnose the root cause of the LokiRequestPanic alert, follow these steps:

1. Check the Loki logs for errors related to the panic
2. Verify that the Loki service is running and configured correctly
3. Investigate recent changes to the Loki configuration or deployment
4. Review system metrics for resource utilization and potential bottlenecks
5. Check for any signs of resource exhaustion or overload

## Mitigation

To mitigate the LokiRequestPanic alert, follow these steps:

1. Roll back recent changes to the Loki configuration or deployment
2. Increase resources (e.g. CPU, memory) allocated to the Loki service
3. Implement rate limiting or queuing to handle high traffic volumes
4. Restart the Loki service to clear out any stuck requests
5. Consider implementing retry logic or circuit breakers to handle failed requests

Note: For more detailed steps and context-specific instructions, refer to the runbook URL provided in the alert annotation: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/loki-internal/LokiRequestPanic.md