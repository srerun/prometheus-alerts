---
title: LokiRequestErrors
description: Troubleshooting for alert LokiRequestErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# LokiRequestErrors

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The {{ $labels.job }} and {{ $labels.route }} are experiencing errors

<details>
  <summary>Alert Rule</summary>

{{% rule "loki/loki-internal.yml" "LokiRequestErrors" %}}

{{% comment %}}

```yaml
alert: LokiRequestErrors
expr: 100 * sum(rate(loki_request_duration_seconds_count{status_code=~"5.."}[1m])) by (namespace, job, route) / sum(rate(loki_request_duration_seconds_count[1m])) by (namespace, job, route) > 10
for: 15m
labels:
    severity: critical
annotations:
    summary: Loki request errors (instance {{ $labels.instance }})
    description: |-
        The {{ $labels.job }} and {{ $labels.route }} are experiencing errors
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/loki-internal/LokiRequestErrors.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
