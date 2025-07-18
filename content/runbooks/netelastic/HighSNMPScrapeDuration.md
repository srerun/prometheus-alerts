---
title: HighSNMPScrapeDuration
description: Troubleshooting for alert HighSNMPScrapeDuration
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HighSNMPScrapeDuration

SNMP scrape duration is above 1 second for more than 5 minutes (current value: {{ $value }} seconds). This may indicate network latency or bng performance issues.

<details>
  <summary>Alert Rule</summary>

{{% rule "netelastic/netelastic.yml" "HighSNMPScrapeDuration" %}}

{{% comment %}}

```yaml
alert: HighSNMPScrapeDuration
expr: snmp_scrape_duration_seconds{module="netelastic"} > 1
for: 5m
labels:
    severity: warning
annotations:
    summary: High SNMP scrape duration
    description: 'SNMP scrape duration is above 1 second for more than 5 minutes (current value: {{ $value }} seconds). This may indicate network latency or bng performance issues.'
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/netelastic/highsnmpscrapeduration/

```

{{% /comment %}}

</details>


Here is a runbook for the HighSNMPScrapeDuration alert:

## Meaning

The HighSNMPScrapeDuration alert is triggered when the SNMP scrape duration for the netelastic module exceeds 1 second for more than 5 minutes. This alert indicates that there may be network latency or performance issues with the Border Network Gateway (BNG).

## Impact

The impact of this alert can be significant, as it may affect the overall performance and reliability of the network. Potential consequences include:

* Delayed or lost data due to slow scrape times
* Inaccurate monitoring data, leading to poor decision-making
* Increased load on the BNG, potentially leading to further performance degradation

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the network latency between the Prometheus server and the BNG using tools such as ping or traceroute.
2. Verify that the BNG is functioning correctly and that there are no issues with the SNMP agent or configuration.
3. Review the SNMP scrape logs to identify any patterns or errors that may indicate the cause of the slow scrape times.
4. Check the CPU and memory usage of the BNG to determine if it is experiencing high load or resource constraints.

## Mitigation

To mitigate the effects of this alert, follow these steps:

1. Investigate and address any network latency issues between the Prometheus server and the BNG.
2. Optimize the SNMP configuration and agent on the BNG to improve performance.
3. Consider implementing SNMP polling interval tuning to reduce the load on the BNG.
4. If the issue persists, consider upgrading the BNG hardware or implementing load balancing to improve performance.

Remember to update the alert annotations with any relevant information or findings during the diagnosis and mitigation process.