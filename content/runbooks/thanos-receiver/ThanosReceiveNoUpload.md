---
title: ThanosReceiveNoUpload
description: Troubleshooting for alert ThanosReceiveNoUpload
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveNoUpload

Thanos Receive {{$labels.instance}} has not uploaded latest data to object storage.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-receiver.yml" "ThanosReceiveNoUpload" %}}

{{% comment %}}

```yaml
alert: ThanosReceiveNoUpload
expr: (up{job=~".*thanos-receive.*"} - 1) + on (job, instance) (sum by (job, instance) (increase(thanos_shipper_uploads_total{job=~".*thanos-receive.*"}[3h])) == 0)
for: 3h
labels:
    severity: critical
annotations:
    summary: Thanos Receive No Upload (instance {{ $labels.instance }})
    description: |-
        Thanos Receive {{$labels.instance}} has not uploaded latest data to object storage.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-receiver/ThanosReceiveNoUpload.md

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
