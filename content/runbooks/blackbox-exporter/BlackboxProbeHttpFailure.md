---
title: BlackboxProbeHttpFailure
description: Troubleshooting for alert BlackboxProbeHttpFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# BlackboxProbeHttpFailure

HTTP status code is not 200-399

<details>
  <summary>Alert Rule</summary>

{{% rule "blackbox/blackbox-exporter.yml" "BlackboxProbeHttpFailure" %}}

{{% comment %}}

```yaml
alert: BlackboxProbeHttpFailure
expr: probe_http_status_code <= 199 OR probe_http_status_code >= 400
for: 0m
labels:
    severity: critical
annotations:
    summary: Blackbox probe HTTP failure (instance {{ $labels.instance }})
    description: |-
        HTTP status code is not 200-399
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/blackbox-exporter/BlackboxProbeHttpFailure.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule:

## Meaning

The BlackboxProbeHttpFailure alert is triggered when a blackbox probe (a synthetic HTTP request) fails to receive a successful HTTP response (200-399) from a target instance. This alert indicates that the target instance is not responding as expected, which can impact the overall health and availability of the system.

## Impact

The impact of this alert can be significant, as it may indicate:

* Loss of visibility into system performance and health
* Unavailability of critical services or applications
* Potential data loss or corruption
* Degraded user experience

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the blackbox exporter logs for errors or exceptions related to the failing probe.
2. Verify that the target instance is reachable and responding to HTTP requests.
3. Check the HTTP response code and body for any errors or clues.
4. Review system and application logs for any signs of errors or issues.
5. Check for any network connectivity issues or firewall rules that may be blocking the probe.

## Mitigation

To mitigate this alert, follow these steps:

1. Restart the blackbox exporter service to reset the probe.
2. Verify that the target instance is running and responding to HTTP requests.
3. Check for any updates or patches to the blackbox exporter or target instance.
4. Investigate and resolve any underlying issues causing the probe failure.
5. Consider increasing the timeout or retries for the probe to make it more resilient to temporary failures.

Remember to update the runbook with specific details and procedures relevant to your environment and systems.