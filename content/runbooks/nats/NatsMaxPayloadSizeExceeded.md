---
title: NatsMaxPayloadSizeExceeded
description: Troubleshooting for alert NatsMaxPayloadSizeExceeded
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsMaxPayloadSizeExceeded

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The max payload size allowed by NATS has been exceeded (1MB)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NatsMaxPayloadSizeExceeded
expr: max(gnatsd_varz_max_payload) > 1024 * 1024
for: 5m
labels:
    severity: critical
annotations:
    summary: Nats max payload size exceeded (instance {{ $labels.instance }})
    description: |-
        The max payload size allowed by NATS has been exceeded (1MB)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/NatsMaxPayloadSizeExceeded

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
