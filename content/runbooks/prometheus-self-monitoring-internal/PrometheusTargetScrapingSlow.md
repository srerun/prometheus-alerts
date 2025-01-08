---
title: PrometheusTargetScrapingSlow
description: Troubleshooting for alert PrometheusTargetScrapingSlow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTargetScrapingSlow

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Prometheus is scraping exporters slowly since it exceeded the requested interval time. Your Prometheus server is under-provisioned.

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTargetScrapingSlow" %}}

{{% comment %}}

```yaml
alert: PrometheusTargetScrapingSlow
expr: prometheus_target_interval_length_seconds{quantile="0.9"} / on (interval, instance, job) prometheus_target_interval_length_seconds{quantile="0.5"} > 1.05
for: 5m
labels:
    severity: warning
annotations:
    summary: Prometheus target scraping slow (instance {{ $labels.instance }})
    description: |-
        Prometheus is scraping exporters slowly since it exceeded the requested interval time. Your Prometheus server is under-provisioned.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTargetScrapingSlow.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
