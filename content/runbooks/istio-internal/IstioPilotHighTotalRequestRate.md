---
title: IstioPilotHighTotalRequestRate
description: Troubleshooting for alert IstioPilotHighTotalRequestRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioPilotHighTotalRequestRate

Number of Istio Pilot push errors is too high (> 5%). Envoy sidecars might have outdated configuration.

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioPilotHighTotalRequestRate" %}}

{{% comment %}}

```yaml
alert: IstioPilotHighTotalRequestRate
expr: sum(rate(pilot_xds_push_errors[1m])) / sum(rate(pilot_xds_pushes[1m])) * 100 > 5
for: 1m
labels:
    severity: warning
annotations:
    summary: Istio Pilot high total request rate (instance {{ $labels.instance }})
    description: |-
        Number of Istio Pilot push errors is too high (> 5%). Envoy sidecars might have outdated configuration.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioPilotHighTotalRequestRate.md

```

{{% /comment %}}

</details>


Here is a runbook for the IstioPilotHighTotalRequestRate alert:

## Meaning

The IstioPilotHighTotalRequestRate alert is triggered when the total request rate of Istio Pilot is high, indicating that Envoy sidecars may have outdated configuration. This can lead to performance issues and errors in the service mesh.

## Impact

* High request rates can cause performance degradation and increased latency in the service mesh.
* Outdated configuration on Envoy sidecars can lead to errors and inconsistencies in the service mesh.
* If left unchecked, this issue can cause cascading failures and impact the overall reliability of the system.

## Diagnosis

1. Check the Pilot logs for any errors or warnings related to configuration pushes.
2. Verify the version of Istio and Envoy to ensure they are up-to-date.
3. Check the number of active Envoy sidecars and their configuration status.
4. Investigate any recent changes to the service mesh configuration or deployment.

## Mitigation

1. Investigate and resolve any underlying issues causing the high request rate, such as network connectivity problems or misconfigured services.
2. Verify that the Pilot is properly configured and deployed.
3. Check for any stuck or failed Envoy sidecars and restart them if necessary.
4. Consider increasing the Pilot's resource allocation (e.g., CPU, memory) to handle the increased request rate.
5. Implement measures to prevent similar issues in the future, such as:
	* Regularly checking the Pilot's performance and configuration.
	* Implementing monitoring and alerting for Pilot errors and performance issues.
	* Ensuring that Envoy sidecars are properly configured and up-to-date.