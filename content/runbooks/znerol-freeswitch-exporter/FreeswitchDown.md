---
title: FreeswitchDown
description: Troubleshooting for alert FreeswitchDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeswitchDown

Freeswitch is unresponsive

<details>
  <summary>Alert Rule</summary>

{{% rule "freeswitch/znerol-freeswitch-exporter.yml" "FreeswitchDown" %}}

{{% comment %}}

```yaml
alert: FreeswitchDown
expr: freeswitch_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Freeswitch down (instance {{ $labels.instance }})
    description: |-
        Freeswitch is unresponsive
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/znerol-freeswitch-exporter/FreeswitchDown.md

```

{{% /comment %}}

</details>


Here is a runbook for the FreeswitchDown alert:

## Meaning

The FreeswitchDown alert is triggered when the Prometheus metric `freeswitch_up` has a value of 0, indicating that Freeswitch is not responding. This alert is critical, as it may impact voice communication services.

## Impact

The impact of Freeswitch being down can be severe, leading to:

* Disruption of voice communication services
* Loss of business continuity
* Potential revenue loss
* Negative impact on customer satisfaction

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Freeswitch server logs for any errors or warnings
2. Verify that the Freeswitch service is running and not stuck in a failed state
3. Check the network connectivity to the Freeswitch server
4. Verify that the Freeswitch configuration is correct and has not been changed recently
5. Check for any resource utilization issues (e.g., CPU, memory) on the Freeswitch server

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Freeswitch service
2. Check and resolve any network connectivity issues
3. Verify and correct any configuration issues
4. Investigate and resolve any resource utilization issues
5. Perform a rolling restart of the Freeswitch cluster (if applicable)
6. Notify the teams and stakeholders about the issue and the steps taken to resolve it

Note: This runbook is a general guideline, and you may need to tailor it to your specific environment and setup.