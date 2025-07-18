---
title: LowCallCompletionRate
description: Troubleshooting for alert LowCallCompletionRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# LowCallCompletionRate

Call completion rate for IPOE users is below 95% for more than 10 minutes (current value: {{ $value }}%). This may indicate connectivity issues.

<details>
  <summary>Alert Rule</summary>

{{% rule "netelastic/netelastic.yml" "LowCallCompletionRate" %}}

{{% comment %}}

```yaml
alert: LowCallCompletionRate
expr: vbng_call_rate_percent{vbng_usertype="ipoe"} < 95
for: 10m
labels:
    severity: warning
annotations:
    summary: Low call completion rate
    description: 'Call completion rate for IPOE users is below 95% for more than 10 minutes (current value: {{ $value }}%). This may indicate connectivity issues.'
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/netelastic/lowcallcompletionrate/

```

{{% /comment %}}

</details>


## Meaning

The `LowCallCompletionRate` alert is triggered when the call completion rate for IPOE users falls below 95% for more than 10 minutes. This alert indicates a potential issue with call connectivity, which may be affecting user experience.

## Impact

* Users may experience dropped or failed calls, leading to frustration and potential loss of business.
* Poor call quality may damage the reputation of the network provider.
* Failure to address the issue promptly may result in a larger-scale outage.

## Diagnosis

To diagnose the root cause of the low call completion rate, follow these steps:

1. Check the network infrastructure for any signs of congestion, packet loss, or high latency.
2. Review call logs to identify patterns of failed or dropped calls.
3. Verify that the IPOE user base is not experiencing any unusual traffic patterns.
4. Consult with network operations teams to determine if any recent changes or maintenance activities may be contributing to the issue.
5. Analyze metrics from other Prometheus alert rules to identify any potential correlations or root causes.

## Mitigation

To mitigate the impact of the low call completion rate, follow these steps:

1. Investigate and address any underlying network infrastructure issues, such as congestion or packet loss.
2. Implement temporary workarounds to improve call quality, such as traffic shaping or priority queuing.
3. Collaborate with network operations teams to perform emergency maintenance or upgrades to resolve the issue.
4. Communicate with stakeholders and users about the ongoing issue and provide updates on the mitigation efforts.
5. Monitor call completion rates closely to ensure that the issue is fully resolved and does not recur.

Note: This runbook provides general guidance and may need to be tailored to specific environments and use cases.