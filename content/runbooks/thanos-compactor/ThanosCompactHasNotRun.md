---
title: ThanosCompactHasNotRun
description: Troubleshooting for alert ThanosCompactHasNotRun
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosCompactHasNotRun

Thanos Compact {{$labels.job}} has not uploaded anything for 24 hours.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-compactor.yml" "ThanosCompactHasNotRun" %}}

{{% comment %}}

```yaml
alert: ThanosCompactHasNotRun
expr: (time() - max by (job) (max_over_time(thanos_objstore_bucket_last_successful_upload_time{job=~".*thanos-compact.*"}[24h]))) / 60 / 60 > 24
for: 0m
labels:
    severity: warning
annotations:
    summary: Thanos Compact Has Not Run (instance {{ $labels.instance }})
    description: |-
        Thanos Compact {{$labels.job}} has not uploaded anything for 24 hours.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-compactor/ThanosCompactHasNotRun.md

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
