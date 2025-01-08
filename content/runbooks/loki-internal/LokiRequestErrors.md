---
title: LokiRequestErrors
description: Troubleshooting for alert LokiRequestErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# LokiRequestErrors

The {{ $labels.job }} and {{ $labels.route }} are experiencing errors

<details>
  <summary>Alert Rule</summary>

{{% rule "loki/loki-internal.yml" "LokiRequestErrors" %}}

{{% comment %}}

```yaml
alert: LokiRequestErrors
expr: 100 * sum(rate(loki_request_duration_seconds_count{status_code=~"5.."}[1m])) by (namespace, job, route) / sum(rate(loki_request_duration_seconds_count[1m])) by (namespace, job, route) > 10
for: 15m
labels:
    severity: critical
annotations:
    summary: Loki request errors (instance {{ $labels.instance }})
    description: |-
        The {{ $labels.job }} and {{ $labels.route }} are experiencing errors
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/loki-internal/LokiRequestErrors.md

```

{{% /comment %}}

</details>


## Meaning

The LokiRequestErrors alert is triggered when the rate of errors for Loki requests exceeds 10% of the total requests over a 1-minute window, sustained for 15 minutes. This indicates that there is a significant issue with Loki request processing, which may impact the reliability and accuracy of log data.

## Impact

* Log data may not be properly processed, leading to incomplete or inaccurate logs.
* This can impact the ability to troubleshoot issues, detect security threats, and monitor system performance.
* The affected namespace, job, and route may experience errors, leading to downstream impacts on dependent services and applications.

## Diagnosis

1. Check the Loki request logs for errors and exceptions.
2. Investigate the root cause of the errors, such as:
	* Network connectivity issues between Loki and the client.
	* Configuration errors in Loki or the client.
	* Resource constraints or overload on Loki or the client.
3. Verify that the Loki instance is properly configured and healthy.
4. Check for any recent changes or deployments that may have introduced the issue.
5. Review the Loki request metrics to identify patterns or trends that may indicate the source of the issue.

## Mitigation

1. Investigate and resolve the root cause of the errors, as identified during diagnosis.
2. Implement temporary workarounds to reduce the error rate, such as:
	* Load balancing or distributing traffic to healthy instances.
	* Increasing resource capacity or scaling Loki instances.
	* Implementing retries or circuit breakers to handle temporary errors.
3. Perform a rolling restart of Loki instances to ensure all instances are healthy and up-to-date.
4. Verify that the error rate has decreased and the alert has cleared.
5. Implement long-term fixes to prevent similar issues from occurring in the future, such as:
	* Improving Loki instance configuration and resource allocation.
	* Enhancing logging and monitoring to detect issues earlier.
	* Developing automated recovery procedures to minimize downtime.