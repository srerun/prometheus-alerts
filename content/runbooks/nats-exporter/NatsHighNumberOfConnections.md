---
title: NatsHighNumberOfConnections
description: Troubleshooting for alert NatsHighNumberOfConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighNumberOfConnections

## Meaning
[//]: # "Short paragraph that explains what the alert means"
NATS server has more than 1000 active connections

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NatsHighNumberOfConnections
expr: gnatsd_connz_num_connections > 1000
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats high number of connections (instance {{ $labels.instance }})
    description: |-
        NATS server has more than 1000 active connections
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/NatsHighNumberOfConnections

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
