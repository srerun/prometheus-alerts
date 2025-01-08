---
title: NomadJobFailed
description: Troubleshooting for alert NomadJobFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NomadJobFailed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Nomad job failed

<details>
  <summary>Alert Rule</summary>

{{% rule "nomad/nomad-internal.yml" "NomadJobFailed" %}}

{{% comment %}}

```yaml
alert: NomadJobFailed
expr: nomad_nomad_job_summary_failed > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Nomad job failed (instance {{ $labels.instance }})
    description: |-
        Nomad job failed
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nomad-internal/NomadJobFailed.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
