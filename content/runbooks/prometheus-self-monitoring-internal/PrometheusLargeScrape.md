---
title: PrometheusLargeScrape
description: Troubleshooting for alert PrometheusLargeScrape
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusLargeScrape

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Prometheus has many scrapes that exceed the sample limit

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusLargeScrape" %}}

<!-- Rule when generated

```yaml
alert: PrometheusLargeScrape
expr: increase(prometheus_target_scrapes_exceeded_sample_limit_total[10m]) > 10
for: 5m
labels:
    severity: warning
annotations:
    summary: Prometheus large scrape (instance {{ $labels.instance }})
    description: |-
        Prometheus has many scrapes that exceed the sample limit
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusLargeScrape.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
