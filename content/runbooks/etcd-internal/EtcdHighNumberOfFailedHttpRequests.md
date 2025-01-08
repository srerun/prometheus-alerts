---
title: EtcdHighNumberOfFailedHttpRequests
description: Troubleshooting for alert EtcdHighNumberOfFailedHttpRequests
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHighNumberOfFailedHttpRequests

More than 1% HTTP failure detected in Etcd

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdHighNumberOfFailedHttpRequests" %}}

{{% comment %}}

```yaml
alert: EtcdHighNumberOfFailedHttpRequests
expr: sum(rate(etcd_http_failed_total[1m])) BY (method) / sum(rate(etcd_http_received_total[1m])) BY (method) > 0.01
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd high number of failed HTTP requests (instance {{ $labels.instance }})
    description: |-
        More than 1% HTTP failure detected in Etcd
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdHighNumberOfFailedHttpRequests.md

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
