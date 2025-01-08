---
title: BlackboxProbeSlowPing
description: Troubleshooting for alert BlackboxProbeSlowPing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# BlackboxProbeSlowPing

Blackbox ping took more than 1s

<details>
  <summary>Alert Rule</summary>

{{% rule "blackbox/blackbox-exporter.yml" "BlackboxProbeSlowPing" %}}

{{% comment %}}

```yaml
alert: BlackboxProbeSlowPing
expr: avg_over_time(probe_icmp_duration_seconds[1m]) > 1
for: 1m
labels:
    severity: warning
annotations:
    summary: Blackbox probe slow ping (instance {{ $labels.instance }})
    description: |-
        Blackbox ping took more than 1s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/blackbox-exporter/BlackboxProbeSlowPing.md

```

{{% /comment %}}

</details>


Here is a runbook for the BlackboxProbeSlowPing alert:

## Meaning

The BlackboxProbeSlowPing alert indicates that the Blackbox exporter's ICMP pings are taking longer than 1 second to complete, on average, over a 1-minute period. This could be a sign of network congestion, high latency, or other issues that may impact the ability to monitor network connectivity.

## Impact

The impact of this alert is that Blackbox exporter's ICMP pings may not be accurately reflecting the current state of network connectivity, potentially leading to false positives or false negatives in monitoring and alerting. This could result in delayed or missed detection of network issues, which could have significant consequences for application performance and reliability.

## Diagnosis

To diagnose the root cause of this issue, you should:

1. Check the Blackbox exporter's logs for any errors or issues related to ICMP pings.
2. Investigate network congestion or high latency issues in the environment.
3. Verify that the Blackbox exporter is configured correctly and that ICMP pings are not being blocked by firewalls or other security controls.
4. Check the instance labels to identify which specific instance is experiencing the slow pings.
5. Review the VALUE and LABELS annotations in the alert to understand the severity and scope of the issue.

## Mitigation

To mitigate the impact of this issue, you can:

1. Investigate and resolve any network congestion or high latency issues in the environment.
2. Optimize the Blackbox exporter's configuration to improve ICMP ping performance.
3. Consider increasing the timeout value for ICMP pings to allow for slower responses.
4. Implement additional monitoring and alerting for Blackbox exporter performance to detect issues earlier.
5. Follow up with the relevant teams to ensure that the issue is resolved and that monitoring and alerting are functioning correctly.