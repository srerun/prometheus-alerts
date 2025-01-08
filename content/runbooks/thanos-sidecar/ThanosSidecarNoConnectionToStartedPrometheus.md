---
title: ThanosSidecarNoConnectionToStartedPrometheus
description: Troubleshooting for alert ThanosSidecarNoConnectionToStartedPrometheus
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosSidecarNoConnectionToStartedPrometheus

Thanos Sidecar {{$labels.instance}} is unhealthy.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-sidecar.yml" "ThanosSidecarNoConnectionToStartedPrometheus" %}}

{{% comment %}}

```yaml
alert: ThanosSidecarNoConnectionToStartedPrometheus
expr: thanos_sidecar_prometheus_up{job=~".*thanos-sidecar.*"} == 0 and on (namespace, pod)prometheus_tsdb_data_replay_duration_seconds != 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Sidecar No Connection To Started Prometheus (instance {{ $labels.instance }})
    description: |-
        Thanos Sidecar {{$labels.instance}} is unhealthy.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-sidecar/ThanosSidecarNoConnectionToStartedPrometheus.md

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
