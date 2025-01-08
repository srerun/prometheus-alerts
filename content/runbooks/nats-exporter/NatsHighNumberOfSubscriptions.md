---
title: NatsHighNumberOfSubscriptions
description: Troubleshooting for alert NatsHighNumberOfSubscriptions
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighNumberOfSubscriptions

## Meaning
[//]: # "Short paragraph that explains what the alert means"
NATS server has more than 1000 active subscriptions

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighNumberOfSubscriptions" %}}

{{% comment %}}

```yaml
alert: NatsHighNumberOfSubscriptions
expr: gnatsd_connz_subscriptions > 1000
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats high number of subscriptions (instance {{ $labels.instance }})
    description: |-
        NATS server has more than 1000 active subscriptions
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighNumberOfSubscriptions.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
