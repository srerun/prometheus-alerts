---
title: EtcdHighCommitDurations
description: Troubleshooting for alert EtcdHighCommitDurations
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHighCommitDurations

Etcd commit duration increasing, 99th percentile is over 0.25s

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdHighCommitDurations" %}}

{{% comment %}}

```yaml
alert: EtcdHighCommitDurations
expr: histogram_quantile(0.99, rate(etcd_disk_backend_commit_duration_seconds_bucket[1m])) > 0.25
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd high commit durations (instance {{ $labels.instance }})
    description: |-
        Etcd commit duration increasing, 99th percentile is over 0.25s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdHighCommitDurations.md

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
