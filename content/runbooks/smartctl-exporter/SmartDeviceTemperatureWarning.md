---
title: SmartDeviceTemperatureWarning
description: Troubleshooting for alert SmartDeviceTemperatureWarning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SmartDeviceTemperatureWarning

Device temperature  warning (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "smart-device-monitoring/smartctl-exporter.yml" "SmartDeviceTemperatureWarning" %}}

{{% comment %}}

```yaml
alert: SmartDeviceTemperatureWarning
expr: smartctl_device_temperature > 60
for: 2m
labels:
    severity: warning
annotations:
    summary: Smart device temperature warning (instance {{ $labels.instance }})
    description: |-
        Device temperature  warning (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/smartctl-exporter/SmartDeviceTemperatureWarning.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule `SmartDeviceTemperatureWarning`:

## Meaning

The `SmartDeviceTemperatureWarning` alert is triggered when the temperature of a smart device exceeds 60 degrees Celsius. This alert is raised to warn of a potential overheating issue that could lead to device failure or data loss.

## Impact

If left unchecked, an overheating smart device can lead to:

* Device failure or shutdown, resulting in data loss or unavailability
* Permanent damage to the device, requiring costly repairs or replacement
* Compromised data integrity or security
* Disruption to critical business operations or services

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the device's temperature reading in Prometheus using the `smartctl_device_temperature` metric.
2. Verify that the device is not in a high-temperature environment or experiencing unusual workload patterns.
3. Review device logs for any error messages or warnings related to temperature or overheating.
4. Check the device's cooling system or fan functionality to ensure proper operation.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately shutdown the device to prevent further overheating and potential damage.
2. Investigate and address any underlying causes of the overheating issue, such as:
	* High ambient temperature
	* Inadequate cooling or airflow
	* Overworked or malfunctioning device components
3. Consider relocating the device to a cooler environment or providing additional cooling mechanisms.
4. Monitor the device's temperature closely and implement preventative measures to avoid future overheating issues.

Note: This runbook is a sample and may require modification to fit your specific use case or environment.