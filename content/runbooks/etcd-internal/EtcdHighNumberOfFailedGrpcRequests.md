---
title: EtcdHighNumberOfFailedGrpcRequests
description: Troubleshooting for alert EtcdHighNumberOfFailedGrpcRequests
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHighNumberOfFailedGrpcRequests

More than 1% GRPC request failure detected in Etcd

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdHighNumberOfFailedGrpcRequests" %}}

{{% comment %}}

```yaml
alert: EtcdHighNumberOfFailedGrpcRequests
expr: sum(rate(grpc_server_handled_total{grpc_code!="OK"}[1m])) BY (grpc_service, grpc_method) / sum(rate(grpc_server_handled_total[1m])) BY (grpc_service, grpc_method) > 0.01
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd high number of failed GRPC requests (instance {{ $labels.instance }})
    description: |-
        More than 1% GRPC request failure detected in Etcd
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdHighNumberOfFailedGrpcRequests.md

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
