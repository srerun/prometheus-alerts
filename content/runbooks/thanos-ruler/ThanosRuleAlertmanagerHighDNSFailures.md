---
title: ThanosRuleAlertmanagerHighDNSFailures
description: Troubleshooting for alert ThanosRuleAlertmanagerHighDNSFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleAlertmanagerHighDNSFailures

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Rule {{$labels.instance}} has {{$value | humanize}}% of failing DNS queries for Alertmanager endpoints.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-ruler.yml" "ThanosRuleAlertmanagerHighDNSFailures" %}}

<!-- Rule when generated

```yaml
alert: ThanosRuleAlertmanagerHighDNSFailures
expr: (sum by (job, instance) (rate(thanos_rule_alertmanagers_dns_failures_total{job=~".*thanos-rule.*"}[5m])) / sum by (job, instance) (rate(thanos_rule_alertmanagers_dns_lookups_total{job=~".*thanos-rule.*"}[5m])) * 100 > 1)
for: 15m
labels:
    severity: warning
annotations:
    summary: Thanos Rule Alertmanager High D N S Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Rule {{$labels.instance}} has {{$value | humanize}}% of failing DNS queries for Alertmanager endpoints.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-ruler/ThanosRuleAlertmanagerHighDNSFailures.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
