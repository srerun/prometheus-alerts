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

{{% rule "cortex/cortex-internal.yml" "CortexRulerConfigurationReloadFailure" %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/cortex-internal/CortexRulerConfigurationReloadFailure.md

```

{{% /comment %}}

</details>


Here is the runbook for the CortexRulerConfigurationReloadFailure alert:

## Meaning

The CortexRulerConfigurationReloadFailure alert is triggered when the Cortex ruler configuration reload is unsuccessful. This means that the Cortex ruler, responsible for running rules and sending notifications, is unable to reload its configuration. This can lead to issues with rule execution, notification delivery, and overall system reliability.

## Impact

The impact of this alert is moderate to high, as it can affect the functionality of the Cortex ruler and potentially cause:

* Rules not being executed correctly
* Notifications not being sent
* Inaccurate or incomplete data being collected
* Increased latency or errors in the system

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cortex ruler logs for errors related to configuration reload.
2. Verify that the configuration file is valid and correctly formatted.
3. Check the permissions and access rights for the configuration file.
4. Restart the Cortex ruler service to attempt to reload the configuration.
5. Check the value of the `cortex_ruler_config_last_reload_successful` metric to see if it has returned to 1.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve any underlying issues causing the configuration reload failure.
2. Verify that the configuration file is correctly formatted and valid.
3. Restart the Cortex ruler service to reload the configuration.
4. Monitor the `cortex_ruler_config_last_reload_successful` metric to ensure it remains at 1.
5. Consider implementing additional logging and monitoring to detect configuration reload failures earlier.

Note: This runbook is a general guide, and the specific steps may vary depending on your environment and setup.