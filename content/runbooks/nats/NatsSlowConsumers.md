---
title: NatsSlowConsumers
description: Troubleshooting for alert NatsSlowConsumers
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsSlowConsumers

## Meaning
[//]: # "Short paragraph that explains what the alert means"
There are slow consumers in NATS for {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NatsSlowConsumers
expr: gnatsd_varz_slow_consumers > 0
for: 3m
labels:
    severity: critical
annotations:
    summary: Nats slow consumers (instance {{ $labels.instance }})
    description: |-
        There are slow consumers in NATS for {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/NatsSlowConsumers

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
