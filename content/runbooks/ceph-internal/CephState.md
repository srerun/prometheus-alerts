---
title: CephState
description: Troubleshooting for alert CephState
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephState

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Ceph instance unhealthy

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephState" %}}

<!-- Rule when generated

```yaml
alert: CephState
expr: ceph_health_status != 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Ceph State (instance {{ $labels.instance }})
    description: |-
        Ceph instance unhealthy
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephState.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
