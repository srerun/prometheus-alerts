---
title: NatsHighConnectionCount
description: Troubleshooting for alert NatsHighConnectionCount
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighConnectionCount

## Meaning
[//]: # "Short paragraph that explains what the alert means"
High number of NATS connections ({{ $value }}) for {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NatsHighConnectionCount
expr: gnatsd_varz_connections > 100
for: 3m
labels:
    severity: warning
annotations:
    summary: Nats high connection count (instance {{ $labels.instance }})
    description: |-
        High number of NATS connections ({{ $value }}) for {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/NatsHighConnectionCount

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
