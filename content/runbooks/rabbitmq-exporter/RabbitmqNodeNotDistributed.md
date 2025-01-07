---
title: RabbitmqNodeNotDistributed
description: Troubleshooting for alert RabbitmqNodeNotDistributed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqNodeNotDistributed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Distribution link state is not 'up'

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: RabbitmqNodeNotDistributed
expr: erlang_vm_dist_node_state < 3
for: 0m
labels:
    severity: critical
annotations:
    summary: RabbitMQ node not distributed (instance {{ $labels.instance }})
    description: |-
        Distribution link state is not 'up'
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/RabbitmqNodeNotDistributed

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
