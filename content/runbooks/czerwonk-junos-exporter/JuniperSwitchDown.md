---
title: JuniperSwitchDown
description: Troubleshooting for alert JuniperSwitchDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JuniperSwitchDown

The switch appears to be down

<details>
  <summary>Alert Rule</summary>

{{% rule "juniper/czerwonk-junos-exporter.yml" "JuniperSwitchDown" %}}

{{% comment %}}

```yaml
alert: JuniperSwitchDown
expr: junos_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Juniper switch down (instance {{ $labels.instance }})
    description: |-
        The switch appears to be down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/czerwonk-junos-exporter/JuniperSwitchDown.md

```

{{% /comment %}}

</details>


Here is a runbook for the JuniperSwitchDown alert rule:

## Meaning

The JuniperSwitchDown alert is triggered when the `junos_up` metric, which monitors the availability of a Juniper switch, returns a value of 0. This indicates that the switch is currently down and not responding.

## Impact

The impact of a Juniper switch being down can be significant, as it may cause network connectivity issues, disrupt critical services, and affect business operations. The duration of the outage will depend on the speed and effectiveness of the mitigation efforts.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the switch's power status and ensure it is properly powered on.
2. Verify that the Junos exporter is running and configured correctly.
3. Review switch logs for any error messages or indications of a hardware or software failure.
4. Check for any recent changes or updates that may have caused the issue.
5. Attempt to ping the switch to confirm its unavailability.

## Mitigation

To mitigate the issue, follow these steps:

1. If the switch is not powered on, power it on and verify it is functioning correctly.
2. If the Junos exporter is not running, start it and verify it is configured correctly.
3. Investigate and address any underlying hardware or software issues identified in the switch logs.
4. If the issue persists, consider escalating to a network engineering team for further assistance.
5. Once the switch is back online, verify that network connectivity has been restored and that critical services are functioning correctly.