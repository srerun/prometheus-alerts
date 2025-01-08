---
title: NatsHighPendingBytes
description: Troubleshooting for alert NatsHighPendingBytes
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighPendingBytes

## Meaning
[//]: # "Short paragraph that explains what the alert means"
High number of NATS pending bytes ({{ $value }}) for {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighPendingBytes" %}}

{{% comment %}}

```yaml
alert: NatsHighPendingBytes
expr: gnatsd_connz_pending_bytes > 100000
for: 3m
labels:
    severity: warning
annotations:
    summary: Nats high pending bytes (instance {{ $labels.instance }})
    description: |-
        High number of NATS pending bytes ({{ $value }}) for {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighPendingBytes.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
