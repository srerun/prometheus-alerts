---
title: ThanosQueryHighDNSFailures
description: Troubleshooting for alert ThanosQueryHighDNSFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosQueryHighDNSFailures

Thanos Query {{$labels.job}} have {{$value | humanize}}% of failing DNS queries for store endpoints.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-query.yml" "ThanosQueryHighDNSFailures" %}}

{{% comment %}}

```yaml
alert: ThanosQueryHighDNSFailures
expr: (sum by (job) (rate(thanos_query_store_apis_dns_failures_total{job=~".*thanos-query.*"}[5m])) / sum by (job) (rate(thanos_query_store_apis_dns_lookups_total{job=~".*thanos-query.*"}[5m]))) * 100 > 1
for: 15m
labels:
    severity: warning
annotations:
    summary: Thanos Query High D N S Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Query {{$labels.job}} have {{$value | humanize}}% of failing DNS queries for store endpoints.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-query/ThanosQueryHighDNSFailures.md

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
