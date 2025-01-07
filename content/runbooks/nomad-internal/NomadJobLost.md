---
title: NomadJobLost
description: Troubleshooting for alert NomadJobLost
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NomadJobLost

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Nomad job lost

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NomadJobLost
expr: nomad_nomad_job_summary_lost > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Nomad job lost (instance {{ $labels.instance }})
    description: |-
        Nomad job lost
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/NomadJobLost

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
