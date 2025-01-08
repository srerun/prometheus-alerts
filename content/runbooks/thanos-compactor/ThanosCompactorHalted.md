---
title: ThanosCompactorHalted
description: Troubleshooting for alert ThanosCompactorHalted
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosCompactorHalted

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Compact {{$labels.job}} has failed to run and now is halted.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-compactor.yml" "ThanosCompactorHalted" %}}

{{% comment %}}

```yaml
alert: ThanosCompactorHalted
expr: thanos_compact_halted{job=~".*thanos-compact.*"} == 1
for: 5m
labels:
    severity: warning
annotations:
    summary: Thanos Compactor Halted (instance {{ $labels.instance }})
    description: |-
        Thanos Compact {{$labels.job}} has failed to run and now is halted.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-compactor/ThanosCompactorHalted.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
