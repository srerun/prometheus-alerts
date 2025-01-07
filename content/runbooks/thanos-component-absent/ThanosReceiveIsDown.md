---
title: ThanosReceiveIsDown
description: Troubleshooting for alert ThanosReceiveIsDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveIsDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
ThanosReceive has disappeared. Prometheus target for the component cannot be discovered.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ThanosReceiveIsDown
expr: absent(up{job=~".*thanos-receive.*"} == 1)
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Receive Is Down (instance {{ $labels.instance }})
    description: |-
        ThanosReceive has disappeared. Prometheus target for the component cannot be discovered.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ThanosReceiveIsDown

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
