---
title: NomadJobQueued
description: Troubleshooting for alert NomadJobQueued
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NomadJobQueued

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Nomad job queued

<details>
  <summary>Alert Rule</summary>

{{% rule "nomad/nomad-internal.yml" "NomadJobQueued" %}}

<!-- Rule when generated

```yaml
alert: NomadJobQueued
expr: nomad_nomad_job_summary_queued > 0
for: 2m
labels:
    severity: warning
annotations:
    summary: Nomad job queued (instance {{ $labels.instance }})
    description: |-
        Nomad job queued
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nomad-internal/NomadJobQueued.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
