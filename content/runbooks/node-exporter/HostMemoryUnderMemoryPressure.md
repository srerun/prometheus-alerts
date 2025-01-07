---
title: HostMemoryUnderMemoryPressure
description: Troubleshooting for alert HostMemoryUnderMemoryPressure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostMemoryUnderMemoryPressure

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The node is under heavy memory pressure. High rate of major page faults

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HostMemoryUnderMemoryPressure
expr: (rate(node_vmstat_pgmajfault[1m]) > 1000) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host memory under memory pressure (instance {{ $labels.instance }})
    description: |-
        The node is under heavy memory pressure. High rate of major page faults
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HostMemoryUnderMemoryPressure

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
