---
title: ThanosRuleConfigReloadFailure
description: Troubleshooting for alert ThanosRuleConfigReloadFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleConfigReloadFailure

Thanos Rule {{$labels.job}} has not been able to reload its configuration.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-ruler.yml" "ThanosRuleConfigReloadFailure" %}}

{{% comment %}}

```yaml
alert: ThanosRuleConfigReloadFailure
expr: avg by (job, instance) (thanos_rule_config_last_reload_successful{job=~".*thanos-rule.*"}) != 1
for: 5m
labels:
    severity: info
annotations:
    summary: Thanos Rule Config Reload Failure (instance {{ $labels.instance }})
    description: |-
        Thanos Rule {{$labels.job}} has not been able to reload its configuration.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-ruler/ThanosRuleConfigReloadFailure.md

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
