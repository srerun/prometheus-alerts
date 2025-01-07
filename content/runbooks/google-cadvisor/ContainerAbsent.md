---
title: ContainerAbsent
description: Troubleshooting for alert ContainerAbsent
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerAbsent

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A container is absent for 5 min

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ContainerAbsent
expr: absent(container_last_seen)
for: 5m
labels:
    severity: warning
annotations:
    summary: Container absent (instance {{ $labels.instance }})
    description: |-
        A container is absent for 5 min
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ContainerAbsent

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
