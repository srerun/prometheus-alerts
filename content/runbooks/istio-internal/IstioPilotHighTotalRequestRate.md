---
title: IstioPilotHighTotalRequestRate
description: Troubleshooting for alert IstioPilotHighTotalRequestRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioPilotHighTotalRequestRate

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Number of Istio Pilot push errors is too high (> 5%). Envoy sidecars might have outdated configuration.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: IstioPilotHighTotalRequestRate
expr: sum(rate(pilot_xds_push_errors[1m])) / sum(rate(pilot_xds_pushes[1m])) * 100 > 5
for: 1m
labels:
    severity: warning
annotations:
    summary: Istio Pilot high total request rate (instance {{ $labels.instance }})
    description: |-
        Number of Istio Pilot push errors is too high (> 5%). Envoy sidecars might have outdated configuration.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/IstioPilotHighTotalRequestRate

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
