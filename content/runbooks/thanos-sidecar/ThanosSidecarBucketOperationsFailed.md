---
title: ThanosSidecarBucketOperationsFailed
description: Troubleshooting for alert ThanosSidecarBucketOperationsFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosSidecarBucketOperationsFailed

Thanos Sidecar {{$labels.instance}} bucket operations are failing

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-sidecar.yml" "ThanosSidecarBucketOperationsFailed" %}}

{{% comment %}}

```yaml
alert: ThanosSidecarBucketOperationsFailed
expr: sum by (job, instance) (rate(thanos_objstore_bucket_operation_failures_total{job=~".*thanos-sidecar.*"}[5m])) > 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Sidecar Bucket Operations Failed (instance {{ $labels.instance }})
    description: |-
        Thanos Sidecar {{$labels.instance}} bucket operations are failing
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-sidecar/ThanosSidecarBucketOperationsFailed.md

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
