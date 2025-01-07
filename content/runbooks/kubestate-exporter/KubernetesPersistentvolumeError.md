---
title: KubernetesPersistentvolumeError
description: Troubleshooting for alert KubernetesPersistentvolumeError
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesPersistentvolumeError

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Persistent volume {{ $labels.persistentvolume }} is in bad state

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: KubernetesPersistentvolumeError
expr: kube_persistentvolume_status_phase{phase=~"Failed|Pending", job="kube-state-metrics"} > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Kubernetes PersistentVolumeClaim pending ({{ $labels.namespace }}/{{ $labels.persistentvolumeclaim }})
    description: |-
        Persistent volume {{ $labels.persistentvolume }} is in bad state
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/KubernetesPersistentvolumeError

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
