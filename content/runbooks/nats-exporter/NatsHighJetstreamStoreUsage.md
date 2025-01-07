---
title: NatsHighJetstreamStoreUsage
description: Troubleshooting for alert NatsHighJetstreamStoreUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighJetstreamStoreUsage

## Meaning
[//]: # "Short paragraph that explains what the alert means"
JetStream store usage is over 80%

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NatsHighJetstreamStoreUsage
expr: gnatsd_varz_jetstream_stats_storage / gnatsd_varz_jetstream_config_max_storage > 0.8
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats high JetStream store usage (instance {{ $labels.instance }})
    description: |-
        JetStream store usage is over 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/NatsHighJetstreamStoreUsage

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
