---
title: ThanosSidecarIsDown
description: Troubleshooting for alert ThanosSidecarIsDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosSidecarIsDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
ThanosSidecar has disappeared. Prometheus target for the component cannot be discovered.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ThanosSidecarIsDown
expr: absent(up{job=~".*thanos-sidecar.*"} == 1)
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Sidecar Is Down (instance {{ $labels.instance }})
    description: |-
        ThanosSidecar has disappeared. Prometheus target for the component cannot be discovered.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ThanosSidecarIsDown

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
