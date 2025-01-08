---
title: EtcdHttpRequestsSlow
description: Troubleshooting for alert EtcdHttpRequestsSlow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHttpRequestsSlow

## Meaning
[//]: # "Short paragraph that explains what the alert means"
HTTP requests slowing down, 99th percentile is over 0.15s

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdHttpRequestsSlow" %}}

{{% comment %}}

```yaml
alert: EtcdHttpRequestsSlow
expr: histogram_quantile(0.99, rate(etcd_http_successful_duration_seconds_bucket[1m])) > 0.15
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd HTTP requests slow (instance {{ $labels.instance }})
    description: |-
        HTTP requests slowing down, 99th percentile is over 0.15s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdHttpRequestsSlow.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
