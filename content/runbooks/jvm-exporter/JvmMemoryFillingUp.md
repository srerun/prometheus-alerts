---
title: JvmMemoryFillingUp
description: Troubleshooting for alert JvmMemoryFillingUp
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JvmMemoryFillingUp

## Meaning
[//]: # "Short paragraph that explains what the alert means"
JVM memory is filling up (> 80%)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: JvmMemoryFillingUp
expr: (sum by (instance)(jvm_memory_used_bytes{area="heap"}) / sum by (instance)(jvm_memory_max_bytes{area="heap"})) * 100 > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: JVM memory filling up (instance {{ $labels.instance }})
    description: |-
        JVM memory is filling up (> 80%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/JvmMemoryFillingUp

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
