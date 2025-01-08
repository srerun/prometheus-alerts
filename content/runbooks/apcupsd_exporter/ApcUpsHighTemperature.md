---
title: ApcUpsHighTemperature
description: Troubleshooting for alert ApcUpsHighTemperature
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApcUpsHighTemperature

Internal temperature is high ({{$value}}°C)

<details>
  <summary>Alert Rule</summary>

{{% rule "apc-ups/apcupsd_exporter.yml" "ApcUpsHighTemperature" %}}

{{% comment %}}

```yaml
alert: ApcUpsHighTemperature
expr: apcupsd_internal_temperature_celsius >= 40
for: 2m
labels:
    severity: warning
annotations:
    summary: APC UPS high temperature (instance {{ $labels.instance }})
    description: |-
        Internal temperature is high ({{$value}}°C)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/apcupsd_exporter/ApcUpsHighTemperature.md

```

{{% /comment %}}

</details>


## Meaning
The ApcUpsHighTemperature alert is triggered when the internal temperature of the APC UPS device exceeds 40°C, as reported by the apcupsd exporter. This alert indicates a potential issue with the UPS's cooling system or environmental conditions, which can lead to device failure or reduced lifespan if left unaddressed.

## Impact
If left unchecked, high internal temperatures can cause:

* Reduced lifespan of the UPS device
* Increased risk of component failure
* Potential data loss or system downtime
* Increased maintenance costs

## Diagnosis
To diagnose the root cause of the high internal temperature, follow these steps:

1. Verify the UPS device's environmental conditions, such as room temperature and humidity.
2. Check the UPS's ventilation and airflow to ensure it is not obstructed.
3. Review the UPS's event logs to identify any errors or warnings related to temperature or cooling system failures.
4. Validate the accuracy of the temperature reading by checking the apcupsd exporter's configuration and sensor data.

## Mitigation
To mitigate the high internal temperature, take the following steps:

1. Ensure the UPS device is installed in a well-ventilated area, away from direct sunlight and heat sources.
2. Verify that the UPS's cooling system is functioning correctly, and clean or replace filters as necessary.
3. Consider relocating the UPS to a cooler area or providing additional cooling measures, such as air conditioning or fans.
4. Monitor the UPS's temperature and event logs closely to detect any potential issues before they become critical.

Remember to refer to the apcupsd exporter documentation and the UPS device's user manual for specific guidance on troubleshooting and maintenance procedures.