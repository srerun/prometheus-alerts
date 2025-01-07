---
title: RabbitmqFileDescriptorsUsage
description: Troubleshooting for alert RabbitmqFileDescriptorsUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqFileDescriptorsUsage

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A node use more than 90% of file descriptors

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: RabbitmqFileDescriptorsUsage
expr: rabbitmq_process_open_fds / rabbitmq_process_max_fds * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ file descriptors usage (instance {{ $labels.instance }})
    description: |-
        A node use more than 90% of file descriptors
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/RabbitmqFileDescriptorsUsage

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
