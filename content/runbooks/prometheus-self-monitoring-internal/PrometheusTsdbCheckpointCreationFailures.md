---
title: PrometheusTsdbCheckpointCreationFailures
description: Troubleshooting for alert PrometheusTsdbCheckpointCreationFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTsdbCheckpointCreationFailures

Prometheus encountered {{ $value }} checkpoint creation failures

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTsdbCheckpointCreationFailures" %}}

{{% comment %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTsdbCheckpointCreationFailures.md

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
