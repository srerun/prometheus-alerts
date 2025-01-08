---
title: NatsMaxPingOperationsExceeded
description: Troubleshooting for alert NatsMaxPingOperationsExceeded
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsMaxPingOperationsExceeded

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The maximum number of ping operations in NATS has exceeded 50

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsMaxPingOperationsExceeded" %}}

{{% comment %}}

```yaml
alert: NatsMaxPingOperationsExceeded
expr: gnatsd_varz_ping_max > 50
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats max ping operations exceeded (instance {{ $labels.instance }})
    description: |-
        The maximum number of ping operations in NATS has exceeded 50
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsMaxPingOperationsExceeded.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
