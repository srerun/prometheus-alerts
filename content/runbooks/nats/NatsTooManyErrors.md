---
title: NatsTooManyErrors
description: Troubleshooting for alert NatsTooManyErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsTooManyErrors

## Meaning
[//]: # "Short paragraph that explains what the alert means"
NATS server has encountered errors in the last 5 minutes

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NatsTooManyErrors
expr: increase(gnatsd_varz_jetstream_stats_api_errors[5m]) > 0
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats too many errors (instance {{ $labels.instance }})
    description: |-
        NATS server has encountered errors in the last 5 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/NatsTooManyErrors

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
