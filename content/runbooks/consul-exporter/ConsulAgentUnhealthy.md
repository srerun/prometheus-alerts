---
title: ConsulAgentUnhealthy
description: Troubleshooting for alert ConsulAgentUnhealthy
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ConsulAgentUnhealthy

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A Consul agent is down

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ConsulAgentUnhealthy
expr: consul_health_node_status{status="critical"} == 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Consul agent unhealthy (instance {{ $labels.instance }})
    description: |-
        A Consul agent is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ConsulAgentUnhealthy

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
