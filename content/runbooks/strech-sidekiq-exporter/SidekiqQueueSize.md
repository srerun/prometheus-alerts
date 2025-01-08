---
title: SidekiqQueueSize
description: Troubleshooting for alert SidekiqQueueSize
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SidekiqQueueSize

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Sidekiq queue {{ $labels.name }} is growing

<details>
  <summary>Alert Rule</summary>

{{% rule "sidekiq/strech-sidekiq-exporter.yml" "SidekiqQueueSize" %}}

<!-- Rule when generated

```yaml
alert: SidekiqQueueSize
expr: sidekiq_queue_size > 100
for: 1m
labels:
    severity: warning
annotations:
    summary: Sidekiq queue size (instance {{ $labels.instance }})
    description: |-
        Sidekiq queue {{ $labels.name }} is growing
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/strech-sidekiq-exporter/SidekiqQueueSize.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
