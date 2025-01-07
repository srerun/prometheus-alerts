---
title: NatsJetstreamConsumersExceeded
description: Troubleshooting for alert NatsJetstreamConsumersExceeded
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsJetstreamConsumersExceeded

## Meaning
[//]: # "Short paragraph that explains what the alert means"
JetStream has more than 100 active consumers

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NatsJetstreamConsumersExceeded
expr: sum(gnatsd_varz_jetstream_stats_accounts) > 100
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats JetStream consumers exceeded (instance {{ $labels.instance }})
    description: |-
        JetStream has more than 100 active consumers
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/NatsJetstreamConsumersExceeded

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
