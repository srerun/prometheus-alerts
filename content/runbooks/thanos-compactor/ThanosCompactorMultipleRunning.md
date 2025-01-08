---
title: ThanosCompactorMultipleRunning
description: Troubleshooting for alert ThanosCompactorMultipleRunning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosCompactorMultipleRunning

No more than one Thanos Compact instance should be running at once. There are {{$value}} instances running.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-compactor.yml" "ThanosCompactorMultipleRunning" %}}

{{% comment %}}

```yaml
alert: ThanosCompactorMultipleRunning
expr: sum by (job) (up{job=~".*thanos-compact.*"}) > 1
for: 5m
labels:
    severity: warning
annotations:
    summary: Thanos Compactor Multiple Running (instance {{ $labels.instance }})
    description: |-
        No more than one Thanos Compact instance should be running at once. There are {{$value}} instances running.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-compactor/ThanosCompactorMultipleRunning.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
