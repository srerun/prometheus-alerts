---
title: CortexRulerConfigurationReloadFailure
description: Troubleshooting for alert CortexRulerConfigurationReloadFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CortexRulerConfigurationReloadFailure

Cortex ruler configuration reload failure (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "cortex/coretex-internal.yml" "CortexRulerConfigurationReloadFailure" %}}

{{% comment %}}

```yaml
alert: CortexRulerConfigurationReloadFailure
expr: cortex_ruler_config_last_reload_successful != 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Cortex ruler configuration reload failure (instance {{ $labels.instance }})
    description: |-
        Cortex ruler configuration reload failure (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/coretex-internal/CortexRulerConfigurationReloadFailure.md

```

{{% /comment %}}

</details>


Here is a runbook for the CortexRulerConfigurationReloadFailure alert:

## Meaning

The CortexRulerConfigurationReloadFailure alert is triggered when the Cortex ruler configuration fails to reload successfully. This alert is critical because it indicates that the ruler configuration has not been updated, which can lead to rules not being executed correctly, resulting in incorrect alerts and notifications.

## Impact

The impact of this alert is high, as it can lead to:

* Incorrect alerting and notification
* Delayed or missed alerts
* Incomplete or inaccurate monitoring data

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cortex ruler logs for errors related to configuration reload.
2. Verify that the ruler configuration file is correct and up-to-date.
3. Check the instance {{ $labels.instance }} for any issues or errors.
4. Verify that the Cortex ruler service is running and healthy.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Cortex ruler service to attempt a configuration reload.
2. Check and update the ruler configuration file if necessary.
3. Investigate and resolve any issues or errors on the instance {{ $labels.instance }}.
4. Verify that the Cortex ruler service is running and healthy after the restart.
5. If the issue persists, escalate to the Cortex administrator or developer team for further assistance.

Remember to monitor the Cortex ruler service and instance {{ $labels.instance }} closely after mitigation to ensure the issue is resolved.