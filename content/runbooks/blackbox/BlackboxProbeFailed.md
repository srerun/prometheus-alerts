---
title: BlackboxProbeFailed
description: Troubleshooting for alert BlackboxProbeFailed
#published: true
date: 2023-12-13T22:43:31.513Z
tags: 
editor: markdown
dateCreated: 2023-12-13T16:44:36.452Z
---

# BlackboxProbeFailed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Probe failed 

<details>
  <summary>Alert Rule</summary>
  
```yaml
  - alert: BlackboxProbeFailed
    expr: probe_success == 0
    for: 0m
    labels:
      severity: critical
    annotations:
      summary: Blackbox probe failed (instance {{ $labels.instance }})
      description: "Probe failed\n  VALUE = {{ $value }}\n  LABELS = {{ $labels }}"
      runbook: https://wiki.ringsq.io/runbook/BlackboxProbeFailed
```
  </details>

## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
