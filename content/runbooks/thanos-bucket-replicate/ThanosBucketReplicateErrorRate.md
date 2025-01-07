---
title: ThanosBucketReplicateErrorRate
description: Troubleshooting for alert ThanosBucketReplicateErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosBucketReplicateErrorRate

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Replicate is failing to run, {{$value | humanize}}% of attempts failed.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ThanosBucketReplicateErrorRate
expr: (sum by (job) (rate(thanos_replicate_replication_runs_total{result="error", job=~".*thanos-bucket-replicate.*"}[5m]))/ on (job) group_left  sum by (job) (rate(thanos_replicate_replication_runs_total{job=~".*thanos-bucket-replicate.*"}[5m]))) * 100 >= 10
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Bucket Replicate Error Rate (instance {{ $labels.instance }})
    description: |-
        Thanos Replicate is failing to run, {{$value | humanize}}% of attempts failed.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ThanosBucketReplicateErrorRate

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
