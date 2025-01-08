---
title: BlackboxProbeFailed
description: Troubleshooting for alert BlackboxProbeFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# BlackboxProbeFailed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Probe failed

<details>
  <summary>Alert Rule</summary>

{{% rule "blackbox/blackbox-exporter.yml" "BlackboxProbeFailed" %}}

<!-- Rule when generated

```yaml
alert: BlackboxProbeFailed
expr: probe_success == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Blackbox probe failed (instance {{ $labels.instance }})
    description: |-
        Probe failed
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/blackbox-exporter/BlackboxProbeFailed.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
