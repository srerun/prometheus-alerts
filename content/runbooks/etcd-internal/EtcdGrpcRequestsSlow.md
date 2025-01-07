---
title: EtcdGrpcRequestsSlow
description: Troubleshooting for alert EtcdGrpcRequestsSlow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdGrpcRequestsSlow

## Meaning
[//]: # "Short paragraph that explains what the alert means"
GRPC requests slowing down, 99th percentile is over 0.15s

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: EtcdGrpcRequestsSlow
expr: histogram_quantile(0.99, sum(rate(grpc_server_handling_seconds_bucket{grpc_type="unary"}[1m])) by (grpc_service, grpc_method, le)) > 0.15
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd GRPC requests slow (instance {{ $labels.instance }})
    description: |-
        GRPC requests slowing down, 99th percentile is over 0.15s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/EtcdGrpcRequestsSlow

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
