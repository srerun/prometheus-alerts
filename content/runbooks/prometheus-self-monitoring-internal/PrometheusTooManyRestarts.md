---
title: PrometheusTooManyRestarts
description: Troubleshooting for alert PrometheusTooManyRestarts
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTooManyRestarts

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Prometheus has restarted more than twice in the last 15 minutes. It might be crashlooping.

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTooManyRestarts" %}}

{{% comment %}}

```yaml
alert: PrometheusTooManyRestarts
expr: changes(process_start_time_seconds{job=~"prometheus|pushgateway|alertmanager"}[15m]) > 2
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus too many restarts (instance {{ $labels.instance }})
    description: |-
        Prometheus has restarted more than twice in the last 15 minutes. It might be crashlooping.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTooManyRestarts.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
