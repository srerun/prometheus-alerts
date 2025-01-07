---
title: ContainerKilled
description: Troubleshooting for alert ContainerKilled
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerKilled

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A container has disappeared

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ContainerKilled
expr: time() - container_last_seen > 60
for: 0m
labels:
    severity: warning
annotations:
    summary: Container killed (instance {{ $labels.instance }})
    description: |-
        A container has disappeared
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ContainerKilled

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
