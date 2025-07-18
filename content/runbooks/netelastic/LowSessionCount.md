---
title: LowSessionCount
description: Troubleshooting for alert LowSessionCount
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# LowSessionCount

Total sessions on slot 0 is below 1 for more than 5 minutes (current value: {{ $value }}). This may indicate a service disruption.

<details>
  <summary>Alert Rule</summary>

{{% rule "netelastic/netelastic.yml" "LowSessionCount" %}}

{{% comment %}}

```yaml
alert: LowSessionCount
expr: vbng_sessions_total{vbng_slotId="0"} < 1
for: 5m
labels:
    severity: critical
annotations:
    summary: Low session count
    description: 'Total sessions on slot 0 is below 1 for more than 5 minutes (current value: {{ $value }}). This may indicate a service disruption.'
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/netelastic/lowsessioncount/

```

{{% /comment %}}

</details>


## Meaning

The LowSessionCount alert is triggered when the total number of sessions on slot 0 of the VBN (Virtual Broadband Network) gateway falls below 1 for more than 5 minutes. This alert indicates a potential service disruption, as a normal functioning VBN gateway should have at least one active session.

## Impact

The impact of this alert is high, as it may indicate a complete loss of service for users connected to the VBN gateway. This could result in:

* Loss of internet access for users
* Disruption to critical services and applications
* Revenue loss due to service unavailability

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the VBN gateway logs for any errors or anomalies that may indicate the cause of the session count drop.
2. Verify that the VBN gateway is properly configured and that all necessary services are running.
3. Check for any network connectivity issues that may be preventing sessions from being established.
4. Verify that there are no software or firmware issues with the VBN gateway that may be causing the session count drop.
5. Check for any security-related issues, such as firewall rules or access control lists, that may be blocking sessions.

## Mitigation

To mitigate the impact of this alert, follow these steps:

1. Restart the VBN gateway services to try to re-establish sessions.
2. Verify that all necessary configuration files are present and valid.
3. Perform a rolling restart of the VBN gateway to ensure that all components are functioning correctly.
4. If the issue persists, consider performing a full reboot of the VBN gateway.
5. If the issue still persists, escalate to the network operations team for further investigation and resolution.

Remember to refer to the [LowSessionCount Runbook](https://srerun.github.io/prometheus-alerts/runbooks/netelastic/lowsessioncount/) for more detailed steps and procedures.