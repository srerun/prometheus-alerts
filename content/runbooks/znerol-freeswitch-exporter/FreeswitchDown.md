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


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
