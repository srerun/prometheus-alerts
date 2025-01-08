---
title: PulsarHighNumberOfSinkErrors
description: Troubleshooting for alert PulsarHighNumberOfSinkErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarHighNumberOfSinkErrors

Observing more than 10 Sink errors per minute

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarHighNumberOfSinkErrors" %}}

{{% comment %}}

```yaml
alert: PulsarHighNumberOfSinkErrors
expr: sum(rate(pulsar_sink_sink_exceptions_total{}[1m]) > 10) by (name)
for: 1m
labels:
    severity: critical
annotations:
    summary: Pulsar high number of sink errors (instance {{ $labels.instance }})
    description: |-
        Observing more than 10 Sink errors per minute
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarHighNumberOfSinkErrors.md

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
