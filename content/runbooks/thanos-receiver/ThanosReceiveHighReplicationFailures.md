---
title: ThanosReceiveHighReplicationFailures
description: Troubleshooting for alert ThanosReceiveHighReplicationFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveHighReplicationFailures

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Receive {{$labels.job}} is failing to replicate {{$value | humanize}}% of requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-receiver.yml" "ThanosReceiveHighReplicationFailures" %}}

<!-- Rule when generated

```yaml
alert: ThanosReceiveHighReplicationFailures
expr: thanos_receive_replication_factor > 1 and ((sum by (job) (rate(thanos_receive_replications_total{result="error", job=~".*thanos-receive.*"}[5m])) / sum by (job) (rate(thanos_receive_replications_total{job=~".*thanos-receive.*"}[5m]))) > (max by (job) (floor((thanos_receive_replication_factor{job=~".*thanos-receive.*"}+1)/ 2)) / max by (job) (thanos_receive_hashring_nodes{job=~".*thanos-receive.*"}))) * 100
for: 5m
labels:
    severity: warning
annotations:
    summary: Thanos Receive High Replication Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Receive {{$labels.job}} is failing to replicate {{$value | humanize}}% of requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-receiver/ThanosReceiveHighReplicationFailures.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
