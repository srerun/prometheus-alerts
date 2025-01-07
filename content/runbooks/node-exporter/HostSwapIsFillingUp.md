---
title: HostSwapIsFillingUp
description: Troubleshooting for alert HostSwapIsFillingUp
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostSwapIsFillingUp

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Swap is filling up (>80%)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HostSwapIsFillingUp
expr: ((1 - (node_memory_SwapFree_bytes / node_memory_SwapTotal_bytes)) * 100 > 80) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host swap is filling up (instance {{ $labels.instance }})
    description: |-
        Swap is filling up (>80%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HostSwapIsFillingUp

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
