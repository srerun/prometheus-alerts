---
title: JuniperHighBandwidthUsage1gib
description: Troubleshooting for alert JuniperHighBandwidthUsage1gib
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JuniperHighBandwidthUsage1gib

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Interface is highly saturated. (> 0.90GiB/s)

<details>
  <summary>Alert Rule</summary>

{{% rule "juniper/czerwonk-junos-exporter.yml" "JuniperHighBandwidthUsage1gib" %}}

<!-- Rule when generated

```yaml
alert: JuniperHighBandwidthUsage1gib
expr: rate(junos_interface_transmit_bytes[1m]) * 8 > 1e+9 * 0.90
for: 1m
labels:
    severity: critical
annotations:
    summary: Juniper high Bandwidth Usage 1GiB (instance {{ $labels.instance }})
    description: |-
        Interface is highly saturated. (> 0.90GiB/s)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/czerwonk-junos-exporter/JuniperHighBandwidthUsage1gib.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
