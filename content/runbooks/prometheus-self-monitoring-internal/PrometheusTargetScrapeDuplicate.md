---
title: PrometheusTargetScrapeDuplicate
description: Troubleshooting for alert PrometheusTargetScrapeDuplicate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTargetScrapeDuplicate

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Prometheus has many samples rejected due to duplicate timestamps but different values

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTargetScrapeDuplicate" %}}

{{% comment %}}

```yaml
alert: PrometheusTargetScrapeDuplicate
expr: increase(prometheus_target_scrapes_sample_duplicate_timestamp_total[5m]) > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus target scrape duplicate (instance {{ $labels.instance }})
    description: |-
        Prometheus has many samples rejected due to duplicate timestamps but different values
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTargetScrapeDuplicate.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
