---
title: ThanosReceiveConfigReloadFailure
description: Troubleshooting for alert ThanosReceiveConfigReloadFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveConfigReloadFailure

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Receive {{$labels.job}} has not been able to reload hashring configurations.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-receiver.yml" "ThanosReceiveConfigReloadFailure" %}}

<!-- Rule when generated

```yaml
alert: ThanosReceiveConfigReloadFailure
expr: avg by (job) (thanos_receive_config_last_reload_successful{job=~".*thanos-receive.*"}) != 1
for: 5m
labels:
    severity: warning
annotations:
    summary: Thanos Receive Config Reload Failure (instance {{ $labels.instance }})
    description: |-
        Thanos Receive {{$labels.job}} has not been able to reload hashring configurations.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-receiver/ThanosReceiveConfigReloadFailure.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
