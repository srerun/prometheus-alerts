---
title: GrafanaAlloyServiceDown
description: Troubleshooting for alert GrafanaAlloyServiceDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# GrafanaAlloyServiceDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Alloy on (instance {{ $labels.instance }}) is not responding or has stopped running.

<details>
  <summary>Alert Rule</summary>

{{% rule "grafana-alloy/grafana-alloy-internal.yml" "GrafanaAlloyServiceDown" %}}

<!-- Rule when generated

```yaml
alert: GrafanaAlloyServiceDown
expr: 'count by (instance) (alloy_build_info) unless count by (instance) (alloy_build_info offset 2m)  '
for: 0m
labels:
    severity: critical
annotations:
    summary: Grafana Alloy service down (instance {{ $labels.instance }})
    description: |-
        Alloy on (instance {{ $labels.instance }}) is not responding or has stopped running.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/grafana-alloy-internal/GrafanaAlloyServiceDown.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
