---
title: FreeswitchSessionsCritical
description: Troubleshooting for alert FreeswitchSessionsCritical
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeswitchSessionsCritical

High sessions usage on {{ $labels.instance }}: {{ $value | printf "%.2f"}}%

<details>
  <summary>Alert Rule</summary>

{{% rule "freeswitch/znerol-freeswitch-exporter.yml" "FreeswitchSessionsCritical" %}}

{{% comment %}}

```yaml
alert: FreeswitchSessionsCritical
expr: (freeswitch_session_active * 100 / freeswitch_session_limit) > 90
for: 5m
labels:
    severity: critical
annotations:
    summary: Freeswitch Sessions Critical (instance {{ $labels.instance }})
    description: |-
        High sessions usage on {{ $labels.instance }}: {{ $value | printf "%.2f"}}%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/znerol-freeswitch-exporter/FreeswitchSessionsCritical.md

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
