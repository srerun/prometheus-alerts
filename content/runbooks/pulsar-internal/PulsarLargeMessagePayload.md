---
title: PulsarLargeMessagePayload
description: Troubleshooting for alert PulsarLargeMessagePayload
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarLargeMessagePayload

Observing large message payload (> 1MB)

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarLargeMessagePayload" %}}

{{% comment %}}

```yaml
alert: PulsarLargeMessagePayload
expr: sum(pulsar_entry_size_overflow > 0) by (topic)
for: 1h
labels:
    severity: warning
annotations:
    summary: Pulsar large message payload (instance {{ $labels.instance }})
    description: |-
        Observing large message payload (> 1MB)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarLargeMessagePayload.md

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
