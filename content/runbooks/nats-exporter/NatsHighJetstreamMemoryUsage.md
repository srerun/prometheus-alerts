---
title: NatsHighJetstreamMemoryUsage
description: Troubleshooting for alert NatsHighJetstreamMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighJetstreamMemoryUsage

## Meaning
[//]: # "Short paragraph that explains what the alert means"
JetStream memory usage is over 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighJetstreamMemoryUsage" %}}

{{% comment %}}

```yaml
alert: NatsHighJetstreamMemoryUsage
expr: gnatsd_varz_jetstream_stats_memory / gnatsd_varz_jetstream_config_max_memory > 0.8
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats high JetStream memory usage (instance {{ $labels.instance }})
    description: |-
        JetStream memory usage is over 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighJetstreamMemoryUsage.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
