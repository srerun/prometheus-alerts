---
title: PrometheusTsdbCheckpointCreationFailures
description: Troubleshooting for alert PrometheusTsdbCheckpointCreationFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTsdbCheckpointCreationFailures

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Prometheus encountered {{ $value }} checkpoint creation failures

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PrometheusTsdbCheckpointCreationFailures
expr: increase(prometheus_tsdb_checkpoint_creations_failed_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus TSDB checkpoint creation failures (instance {{ $labels.instance }})
    description: |-
        Prometheus encountered {{ $value }} checkpoint creation failures
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PrometheusTsdbCheckpointCreationFailures

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
