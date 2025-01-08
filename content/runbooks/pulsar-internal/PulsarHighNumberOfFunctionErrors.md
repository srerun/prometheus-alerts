---
title: PulsarHighNumberOfFunctionErrors
description: Troubleshooting for alert PulsarHighNumberOfFunctionErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarHighNumberOfFunctionErrors

Observing more than 10 Function errors per minute

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarHighNumberOfFunctionErrors" %}}

{{% comment %}}

```yaml
alert: PulsarHighNumberOfFunctionErrors
expr: sum((rate(pulsar_function_user_exceptions_total{}[1m]) + rate(pulsar_function_system_exceptions_total{}[1m])) > 10) by (name)
for: 1m
labels:
    severity: critical
annotations:
    summary: Pulsar high number of function errors (instance {{ $labels.instance }})
    description: |-
        Observing more than 10 Function errors per minute
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarHighNumberOfFunctionErrors.md

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
