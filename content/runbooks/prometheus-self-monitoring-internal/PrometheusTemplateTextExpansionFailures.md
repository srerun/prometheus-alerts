---
title: PrometheusTemplateTextExpansionFailures
description: Troubleshooting for alert PrometheusTemplateTextExpansionFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTemplateTextExpansionFailures

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Prometheus encountered {{ $value }} template text expansion failures

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTemplateTextExpansionFailures" %}}

{{% comment %}}

```yaml
alert: PrometheusTemplateTextExpansionFailures
expr: increase(prometheus_template_text_expansion_failures_total[3m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus template text expansion failures (instance {{ $labels.instance }})
    description: |-
        Prometheus encountered {{ $value }} template text expansion failures
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTemplateTextExpansionFailures.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
