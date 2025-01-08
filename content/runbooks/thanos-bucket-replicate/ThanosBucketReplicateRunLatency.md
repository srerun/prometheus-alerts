---
title: ThanosBucketReplicateRunLatency
description: Troubleshooting for alert ThanosBucketReplicateRunLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosBucketReplicateRunLatency

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Replicate {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for the replicate operations.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-bucket-replicate.yml" "ThanosBucketReplicateRunLatency" %}}

<!-- Rule when generated

```yaml
alert: ThanosBucketReplicateRunLatency
expr: (histogram_quantile(0.99, sum by (job) (rate(thanos_replicate_replication_run_duration_seconds_bucket{job=~".*thanos-bucket-replicate.*"}[5m]))) > 20 and  sum by (job) (rate(thanos_replicate_replication_run_duration_seconds_bucket{job=~".*thanos-bucket-replicate.*"}[5m])) > 0)
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Bucket Replicate Run Latency (instance {{ $labels.instance }})
    description: |-
        Thanos Replicate {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for the replicate operations.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-bucket-replicate/ThanosBucketReplicateRunLatency.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
