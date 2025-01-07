---
title: NatsServerDown
description: Troubleshooting for alert NatsServerDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsServerDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
NATS server has been down for more than 5 minutes

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NatsServerDown
expr: absent(up{job="nats"})
for: 5m
labels:
    severity: critical
annotations:
    summary: Nats server down (instance {{ $labels.instance }})
    description: |-
        NATS server has been down for more than 5 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/NatsServerDown

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
