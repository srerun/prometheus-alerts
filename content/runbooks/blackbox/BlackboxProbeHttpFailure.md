---
title: BlackboxProbeHttpFailure
description: Troubleshooting for alert BlackboxProbeHttpFailure
#published: true
date: 2023-12-13T19:40:22.063Z
tags: 
editor: markdown
dateCreated: 2023-12-13T16:44:39.061Z
---

# BlackboxProbeHttpFailure

## Meaning
[//]: # "Short paragraph that explains what the alert means"
HTTP status code is not 200-399 


<details>
  <summary>Alert Rule</summary>

```yaml
alert: BlackboxProbeFailed
expr: probe_success == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Blackbox probe failed (instance {{ $labels.instance }})
    description: |-
        Probe failed
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/BlackboxProbeFailed

```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
