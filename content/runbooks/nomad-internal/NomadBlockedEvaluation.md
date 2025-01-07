---
title: NomadBlockedEvaluation
description: Troubleshooting for alert NomadBlockedEvaluation
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NomadBlockedEvaluation

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Nomad blocked evaluation

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NomadBlockedEvaluation
expr: nomad_nomad_blocked_evals_total_blocked > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Nomad blocked evaluation (instance {{ $labels.instance }})
    description: |-
        Nomad blocked evaluation
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/NomadBlockedEvaluation

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
