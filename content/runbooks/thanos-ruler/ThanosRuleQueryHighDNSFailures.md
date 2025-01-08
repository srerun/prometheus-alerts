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


## Meaning

The `ThanosRuleQueryHighDNSFailures` alert is triggered when the percentage of failed DNS queries for Thanos Rule query endpoints exceeds 1% over a 5-minute period. This indicates that there may be issues with DNS resolution or connectivity to the query endpoints, which can impact the availability and performance of Thanos Rule.

## Impact

Failed DNS queries can lead to:

* Delays or failures in querying rules, causing issues with alerting and notification systems
* Increased latency and errors in Thanos Rule, potentially impacting downstream systems
* Difficulty in troubleshooting issues due to incomplete or missing data

If left unchecked, high DNS failure rates can lead to a cascading failure of the entire monitoring and alerting system.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Rule logs for errors related to DNS resolution or query endpoint connectivity.
2. Verify that the DNS servers are functioning correctly and reachable from the Thanos Rule instances.
3. Investigate any recent changes to the network configuration, DNS setup, or query endpoint URLs.
4. Check the system resources (e.g., CPU, memory, and disk space) to ensure they are within acceptable limits.
5. Review the Thanos Rule configuration to ensure it is correct and up-to-date.

## Mitigation

To mitigate the issue, follow these steps:

1. Check and update the DNS server configuration to ensure it is correct and reachable.
2. Verify that the query endpoint URLs are correct and reachable.
3. Implement retry mechanisms or circuit breakers to handle temporary DNS failures.
4. Consider implementing a DNS cache or proxy to reduce the load on the DNS servers.
5. Review and optimize the Thanos Rule configuration to reduce the load on the system resources.

Remember to investigate and address any underlying causes of the DNS failures to prevent future occurrences.