---
title: SmartDeviceTemperatureCritical
description: Troubleshooting for alert SmartDeviceTemperatureCritical
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SmartDeviceTemperatureCritical

Device temperature critical  (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "smart-device-monitoring/smartctl-exporter.yml" "SmartDeviceTemperatureCritical" %}}

{{% comment %}}

```yaml
alert: SmartDeviceTemperatureCritical
expr: smartctl_device_temperature > 80
for: 2m
labels:
    severity: critical
annotations:
    summary: Smart device temperature critical (instance {{ $labels.instance }})
    description: |-
        Device temperature critical  (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/smartctl-exporter/SmartDeviceTemperatureCritical.md

```

{{% /comment %}}

</details>


## Meaning

The SmartDeviceTemperatureCritical alert is triggered when the temperature of a smart device, as reported by the `smartctl` exporter, exceeds 80 degrees Celsius. This indicates a critical temperature threshold has been breached, which can potentially lead to device failure, data loss, or even physical damage.

## Impact

The impact of this alert includes:

* Potential device failure or shutdown, leading to data unavailability and system downtime
* Increased risk of data loss or corruption due to overheating
* Possible physical damage to the device or surrounding components
* Decreased reliability and lifespan of the device

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the device's temperature reading using the `smartctl` exporter
2. Verify that the temperature reading is accurate and not a sensor error
3. Investigate possible causes of the high temperature, such as:
	* High ambient temperature
	* Poor cooling or ventilation
	* Overloaded or malfunctioning device
	* Faulty temperature sensor
4. Review system logs for any related errors or warnings
5. Consult device documentation and manufacturer guidelines for recommended operating temperatures

## Mitigation

To mitigate the issue, take the following steps:

1. Immediately shutdown the device to prevent further damage
2. Verify that the device is properly cooled and ventilated
3. Check for any blockages or obstructions in the device's airflow
4. Consider relocating the device to a cooler environment
5. Monitor the device's temperature closely and take corrective action if it continues to rise
6. Consider replacing the device if it is faulty or malfunctioning
7. Review and update device configuration and settings to prevent similar issues in the future