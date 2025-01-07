---
title: HostOomKillDetected
description: Troubleshooting for alert HostOomKillDetected
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostOomKillDetected

## Meaning
[//]: # "Short paragraph that explains what the alert means"
OOM kill detected

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HostOomKillDetected
expr: (increase(node_vmstat_oom_kill[1m]) > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: warning
annotations:
    summary: Host OOM kill detected (instance {{ $labels.instance }})
    description: |-
        OOM kill detected
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HostOomKillDetected

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
