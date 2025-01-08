---
title: BlackboxProbeFailed
description: Troubleshooting for alert BlackboxProbeFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# BlackboxProbeFailed

Probe failed

<details>
  <summary>Alert Rule</summary>

{{% rule "blackbox/blackbox-exporter.yml" "BlackboxProbeFailed" %}}

{{% comment %}}

```yaml
alert: BlackboxProbeFailed
expr: probe_success == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Blackbox probe failed (instance {{ $labels.instance }})
    description: |-
        Probe failed
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/blackbox-exporter/BlackboxProbeFailed.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "BlackboxProbeFailed":

## Meaning

The BlackboxProbeFailed alert is triggered when a blackbox probe fails to return a successful response. This probe is used to monitor the availability and responsiveness of a specific endpoint or service. When this alert is triggered, it indicates that the probe was unable to connect to the endpoint or receive a valid response, which may indicate a problem with the service or endpoint being monitored.

## Impact

The impact of this alert varies depending on the specific use case and service being monitored. However, in general, a failed blackbox probe can indicate:

* Loss of service availability or responsiveness
* Potential outage or downtime for users
* Increased latency or errors for dependent systems
* Possible security concerns if the probe is monitoring a critical security endpoint

## Diagnosis

To diagnose the root cause of the BlackboxProbeFailed alert, follow these steps:

1. Check the blackbox exporter logs for errors or issues related to the probe
2. Verify that the endpoint or service being monitored is reachable and responding correctly
3. Check for any network connectivity issues or firewall rules that may be blocking the probe
4. Verify that the probe configuration is correct and up-to-date
5. Check for any recent changes or updates to the service or endpoint being monitored that may have caused the probe to fail

## Mitigation

To mitigate the effects of the BlackboxProbeFailed alert, follow these steps:

1. Restart the blackbox exporter service to retry the probe
2. Investigate and resolve any network connectivity issues or firewall rules that may be blocking the probe
3. Update the probe configuration to ensure it is correct and up-to-date
4. Verify that the service or endpoint being monitored is available and responsive
5. Notify relevant teams or stakeholders of the issue and work to resolve it as quickly as possible to minimize impact to users.

Additionally, consider implementing measures to prevent similar issues in the future, such as:

* Implementing redundancy or failover mechanisms for critical services or endpoints
* Improving monitoring and logging for the blackbox exporter and dependent systems
* Conducting regular maintenance and testing of the probe and endpoint configurations