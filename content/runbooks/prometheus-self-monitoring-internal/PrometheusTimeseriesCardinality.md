---
title: PrometheusTimeseriesCardinality
description: Troubleshooting for alert PrometheusTimeseriesCardinality
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTimeseriesCardinality

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The "{{ $labels.name }}" timeseries cardinality is getting very high: {{ $value }}

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTimeseriesCardinality" %}}

<!-- Rule when generated

```yaml
alert: PrometheusTimeseriesCardinality
expr: label_replace(count by(__name__) ({__name__=~".+"}), "name", "$1", "__name__", "(.+)") > 10000
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus timeseries cardinality (instance {{ $labels.instance }})
    description: |-
        The "{{ $labels.name }}" timeseries cardinality is getting very high: {{ $value }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTimeseriesCardinality.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
