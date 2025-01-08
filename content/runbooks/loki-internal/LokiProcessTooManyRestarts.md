---
title: LokiProcessTooManyRestarts
description: Troubleshooting for alert LokiProcessTooManyRestarts
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# LokiProcessTooManyRestarts

A loki process had too many restarts (target {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "loki/loki-internal.yml" "LokiProcessTooManyRestarts" %}}

{{% comment %}}

```yaml
alert: LokiProcessTooManyRestarts
expr: changes(process_start_time_seconds{job=~".*loki.*"}[15m]) > 2
for: 0m
labels:
    severity: warning
annotations:
    summary: Loki process too many restarts (instance {{ $labels.instance }})
    description: |-
        A loki process had too many restarts (target {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/loki-internal/LokiProcessTooManyRestarts.md

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
