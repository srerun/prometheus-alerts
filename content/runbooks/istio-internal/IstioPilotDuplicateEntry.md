---
title: IstioPilotDuplicateEntry
description: Troubleshooting for alert IstioPilotDuplicateEntry
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioPilotDuplicateEntry

Istio pilot duplicate entry error.

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioPilotDuplicateEntry" %}}

{{% comment %}}

```yaml
alert: IstioPilotDuplicateEntry
expr: sum(rate(pilot_duplicate_envoy_clusters{}[5m])) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Istio Pilot Duplicate Entry (instance {{ $labels.instance }})
    description: |-
        Istio pilot duplicate entry error.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioPilotDuplicateEntry.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the IstioPilotDuplicateEntry alert rule:

## Meaning

The IstioPilotDuplicateEntry alert is triggered when the Istio Pilot component detects duplicate entries in its Envoy cluster configuration. This can lead to inconsistent and unpredictable behavior in the service mesh, including traffic misrouting and errors.

## Impact

The impact of this alert can be severe, as it can cause:

* Traffic misrouting and errors
* Inconsistent behavior in the service mesh
* Downtime and unavailability of services
* Increased latency and decreased performance

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Istio Pilot logs for error messages related to duplicate entries.
2. Verify that the Envoy cluster configuration is correct and up-to-date.
3. Check for any recent changes to the Istio Pilot configuration or deployment.
4. Verify that the Istio Pilot instance is running with the correct permissions and access.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the affected Istio Pilot instance to reset the Envoy cluster configuration.
2. Verify that the Envoy cluster configuration is correct and up-to-date.
3. Check for any recent changes to the Istio Pilot configuration or deployment and revert if necessary.
4. Verify that the Istio Pilot instance is running with the correct permissions and access.
5. Consider increasing the logging level for Istio Pilot to gather more detailed information about the issue.

Note: If the issue persists, it may be necessary to involve the Istio development team or consult the Istio documentation for further guidance.