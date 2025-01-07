---
title: HostOutOfMemory
description: Troubleshooting for alert HostOutOfMemory
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostOutOfMemory

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Node memory is filling up (< 10% left)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HostOutOfMemory
expr: (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes * 100 < 10) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host out of memory (instance {{ $labels.instance }})
    description: |-
        Node memory is filling up (< 10% left)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HostOutOfMemory

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
