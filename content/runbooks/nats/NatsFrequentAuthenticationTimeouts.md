---
title: NatsFrequentAuthenticationTimeouts
description: Troubleshooting for alert NatsFrequentAuthenticationTimeouts
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsFrequentAuthenticationTimeouts

## Meaning
[//]: # "Short paragraph that explains what the alert means"
There have been more than 5 authentication timeouts in the last 5 minutes

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NatsFrequentAuthenticationTimeouts
expr: increase(gnatsd_varz_auth_timeout[5m]) > 5
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats frequent authentication timeouts (instance {{ $labels.instance }})
    description: |-
        There have been more than 5 authentication timeouts in the last 5 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/NatsFrequentAuthenticationTimeouts

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
