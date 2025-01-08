---
title: ThanosStoreBucketHighOperationFailures
description: Troubleshooting for alert ThanosStoreBucketHighOperationFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosStoreBucketHighOperationFailures

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Store {{$labels.job}} Bucket is failing to execute {{$value | humanize}}% of operations.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-store.yml" "ThanosStoreBucketHighOperationFailures" %}}

{{% comment %}}

```yaml
alert: ThanosStoreBucketHighOperationFailures
expr: (sum by (job) (rate(thanos_objstore_bucket_operation_failures_total{job=~".*thanos-store.*"}[5m])) / sum by (job) (rate(thanos_objstore_bucket_operations_total{job=~".*thanos-store.*"}[5m])) * 100 > 5)
for: 15m
labels:
    severity: warning
annotations:
    summary: Thanos Store Bucket High Operation Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Store {{$labels.job}} Bucket is failing to execute {{$value | humanize}}% of operations.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-store/ThanosStoreBucketHighOperationFailures.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
