---
title: PulsarReadOnlyBookies
description: Troubleshooting for alert PulsarReadOnlyBookies
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarReadOnlyBookies

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Observing Readonly Bookies

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PulsarReadOnlyBookies
expr: count(bookie_SERVER_STATUS{} == 0) by (pod)
for: 5m
labels:
    severity: critical
annotations:
    summary: Pulsar read only bookies (instance {{ $labels.instance }})
    description: |-
        Observing Readonly Bookies
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PulsarReadOnlyBookies

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
