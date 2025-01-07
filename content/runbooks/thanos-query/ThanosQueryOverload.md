---
title: ThanosQueryOverload
description: Troubleshooting for alert ThanosQueryOverload
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosQueryOverload

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Query {{$labels.job}} has been overloaded for more than 15 minutes. This may be a symptom of excessive simultanous complex requests, low performance of the Prometheus API, or failures within these components. Assess the health of the Thanos query instances, the connnected Prometheus instances, look for potential senders of these requests and then contact support.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ThanosQueryOverload
expr: (max_over_time(thanos_query_concurrent_gate_queries_max[5m]) - avg_over_time(thanos_query_concurrent_gate_queries_in_flight[5m]) < 1)
for: 15m
labels:
    severity: warning
annotations:
    summary: Thanos Query Overload (instance {{ $labels.instance }})
    description: |-
        Thanos Query {{$labels.job}} has been overloaded for more than 15 minutes. This may be a symptom of excessive simultanous complex requests, low performance of the Prometheus API, or failures within these components. Assess the health of the Thanos query instances, the connnected Prometheus instances, look for potential senders of these requests and then contact support.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ThanosQueryOverload

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
