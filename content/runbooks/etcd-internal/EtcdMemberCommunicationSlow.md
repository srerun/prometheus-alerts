---
title: EtcdMemberCommunicationSlow
description: Troubleshooting for alert EtcdMemberCommunicationSlow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdMemberCommunicationSlow

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Etcd member communication slowing down, 99th percentile is over 0.15s

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: EtcdMemberCommunicationSlow
expr: histogram_quantile(0.99, rate(etcd_network_peer_round_trip_time_seconds_bucket[1m])) > 0.15
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd member communication slow (instance {{ $labels.instance }})
    description: |-
        Etcd member communication slowing down, 99th percentile is over 0.15s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/EtcdMemberCommunicationSlow

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
