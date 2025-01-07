---
title: CorednsPanicCount
description: Troubleshooting for alert CorednsPanicCount
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CorednsPanicCount

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Number of CoreDNS panics encountered

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CorednsPanicCount
expr: increase(coredns_panics_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: CoreDNS Panic Count (instance {{ $labels.instance }})
    description: |-
        Number of CoreDNS panics encountered
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CorednsPanicCount

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
