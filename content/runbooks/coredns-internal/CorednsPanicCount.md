---
title: CorednsPanicCount
description: Troubleshooting for alert CorednsPanicCount
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CorednsPanicCount

Number of CoreDNS panics encountered

<details>
  <summary>Alert Rule</summary>

{{% rule "coredns/coredns-internal.yml" "CorednsPanicCount" %}}

{{% comment %}}

```yaml
alert: CorednsPanicCount
expr: increase(coredns_panics_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: CoreDNS Panic Count (instance {{ $labels.instance }})
    description: |-
        Number of CoreDNS panics encountered
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/coredns-internal/CorednsPanicCount.md

```

{{% /comment %}}

</details>


Here is a runbook for the CorednsPanicCount alert rule:

## Meaning

The CorednsPanicCount alert is triggered when the CoreDNS server experiences one or more panic events within a 1-minute window. This indicates that the CoreDNS process has encountered a critical error that causes it to crash or terminate abnormally.

## Impact

The impact of this alert is high, as a CoreDNS panic can lead to:

* Loss of DNS resolution functionality
* Increased latency or timeouts for DNS queries
* Potential disruption to dependent applications and services
* Decreased overall system reliability and availability

## Diagnosis

To diagnose the root cause of the CoreDNS panic, follow these steps:

1. **Check the CoreDNS logs**: Review the CoreDNS logs to identify the specific error or condition that led to the panic.
2. **Verify system resource utilization**: Check the system resources (CPU, memory, disk space) to ensure they are within healthy limits.
3. **Investigate recent changes**: Review recent changes to the CoreDNS configuration, plugins, or underlying system to identify potential causes of the panic.
4. **Run a CoreDNS debug session**: Enable debug logging for CoreDNS and reproduce the issue to gather more detailed information.

## Mitigation

To mitigate the effects of the CoreDNS panic, follow these steps:

1. **Restart the CoreDNS service**: Immediately restart the CoreDNS service to restore DNS functionality.
2. **Investigate and address root cause**: Perform a thorough investigation to identify and address the root cause of the panic, following the diagnosis steps above.
3. **Implement additional monitoring and alerting**: Enhance monitoring and alerting to detect potential issues before they lead to a CoreDNS panic.
4. **Consider implementing a CoreDNS cluster**: If the CoreDNS instance is critical to the system, consider implementing a CoreDNS cluster to provide redundancy and high availability.