---
title: ThanosRuleQueryHighDNSFailures
description: Troubleshooting for alert ThanosRuleQueryHighDNSFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleQueryHighDNSFailures

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Rule {{$labels.job}} has {{$value | humanize}}% of failing DNS queries for query endpoints.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-ruler.yml" "ThanosRuleQueryHighDNSFailures" %}}

{{% comment %}}

```yaml
alert: ThanosRuleQueryHighDNSFailures
expr: (sum by (job, instance) (rate(thanos_rule_query_apis_dns_failures_total{job=~".*thanos-rule.*"}[5m])) / sum by (job, instance) (rate(thanos_rule_query_apis_dns_lookups_total{job=~".*thanos-rule.*"}[5m])) * 100 > 1)
for: 15m
labels:
    severity: warning
annotations:
    summary: Thanos Rule Query High D N S Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Rule {{$labels.job}} has {{$value | humanize}}% of failing DNS queries for query endpoints.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-ruler/ThanosRuleQueryHighDNSFailures.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
