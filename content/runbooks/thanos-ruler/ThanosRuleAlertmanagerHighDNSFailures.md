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

Thanos Rule {{$labels.instance}} has {{$value | humanize}}% of failing DNS queries for Alertmanager endpoints.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-ruler.yml" "ThanosRuleAlertmanagerHighDNSFailures" %}}

{{% comment %}}

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

{{% /comment %}}

</details>


## Meaning

The ThanosRuleAlertmanagerHighDNSFailures alert is triggered when the percentage of failing DNS queries for Alertmanager endpoints exceeds 1% for a Thanos Rule instance. This alert indicates a potential issue with the DNS resolution for Alertmanager endpoints, which may lead to alert delivery failures.

## Impact

* Delayed or failed alert delivery to notification channels
* Incomplete or inaccurate alerting and notifications
* Increased latency or timeouts in alert processing
* Potential impact on incident response and resolution times

## Diagnosis

1. Check the Thanos Rule instance logs for DNS resolution errors or timeouts
2. Verify the DNS configuration and settings for the Alertmanager endpoints
3. Check the network connectivity and latency between the Thanos Rule instance and the Alertmanager endpoints
4. Investigate any recent changes to the DNS infrastructure or Alertmanager configuration
5. Review the Prometheus metrics for other related errors or issues

## Mitigation

1. Investigate and resolve any DNS resolution errors or timeouts on the Thanos Rule instance
2. Verify and update the DNS configuration to ensure correct resolution of Alertmanager endpoints
3. Check and optimize the network connectivity and latency between the Thanos Rule instance and the Alertmanager endpoints
4. Consider implementing redundancy or fallback mechanisms for DNS resolution
5. Monitor the alerting system for any continued issues or errors and take corrective action as needed.