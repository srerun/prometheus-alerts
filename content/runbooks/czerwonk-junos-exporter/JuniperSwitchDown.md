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


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
