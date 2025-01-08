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

The ThanosQueryHighDNSFailures alert is triggered when the percentage of failing DNS queries for Thanos Query store endpoints exceeds 1%. This alert indicates that there is an issue with DNS resolution for Thanos Query instances, which may impact the overall performance and reliability of the system.

## Impact

* High DNS failure rates can lead to timeouts and errors in Thanos Query, resulting in delayed or failed queries.
* This may cause:
	+ Increased latency and decreased performance for dependent systems.
	+ Incomplete or incorrect data being returned to users.
	+ Potential data loss or corruption.

## Diagnosis

To diagnose the issue, follow these steps:

1. **Check DNS resolution**: Verify that DNS resolution is working correctly for the affected Thanos Query instance(s).
2. **Investigate DNS server logs**: Analyze DNS server logs to identify any errors or issues that may be causing the high failure rate.
3. **Verify network connectivity**: Ensure that there are no network connectivity issues between the Thanos Query instance and the DNS servers.
4. **Check Thanos Query configuration**: Review the Thanos Query configuration to ensure that it is correctly configured for DNS resolution.

## Mitigation

To mitigate the issue, follow these steps:

1. **Restart DNS service**: Restart the DNS service to ensure it is running correctly.
2. **Check and update DNS server configuration**: Review and update the DNS server configuration to resolve any issues identified during diagnosis.
3. **Increase DNS timeout**: Temporarily increase the DNS timeout to allow for more time to resolve DNS queries.
4. **Implement DNS caching**: Consider implementing DNS caching to reduce the load on DNS servers and improve performance.
5. **Monitor and investigate**: Continuously monitor the situation and investigate the root cause of the issue to prevent future occurrences.