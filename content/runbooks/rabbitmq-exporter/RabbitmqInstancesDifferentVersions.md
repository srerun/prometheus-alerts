---
title: RabbitmqInstancesDifferentVersions
description: Troubleshooting for alert RabbitmqInstancesDifferentVersions
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqInstancesDifferentVersions

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Running different version of RabbitMQ in the same cluster, can lead to failure.

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqInstancesDifferentVersions" %}}

{{% comment %}}

```yaml
alert: RabbitmqInstancesDifferentVersions
expr: count(count(rabbitmq_build_info) by (rabbitmq_version)) > 1
for: 1h
labels:
    severity: warning
annotations:
    summary: RabbitMQ instances different versions (instance {{ $labels.instance }})
    description: |-
        Running different version of RabbitMQ in the same cluster, can lead to failure.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqInstancesDifferentVersions.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
