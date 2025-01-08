---
title: NatsHighMemoryUsage
description: Troubleshooting for alert NatsHighMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighMemoryUsage

## Meaning
[//]: # "Short paragraph that explains what the alert means"
NATS server memory usage is above 200MB for {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighMemoryUsage" %}}

{{% comment %}}

```yaml
alert: NatsHighMemoryUsage
expr: gnatsd_varz_mem > 200 * 1024 * 1024
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats high memory usage (instance {{ $labels.instance }})
    description: |-
        NATS server memory usage is above 200MB for {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighMemoryUsage.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
