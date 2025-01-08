---
title: PulsarHighNumberOfFunctionErrors
description: Troubleshooting for alert PulsarHighNumberOfFunctionErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarHighNumberOfFunctionErrors

Observing more than 10 Function errors per minute

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarHighNumberOfFunctionErrors" %}}

{{% comment %}}

```yaml
alert: PulsarHighNumberOfFunctionErrors
expr: sum((rate(pulsar_function_user_exceptions_total{}[1m]) + rate(pulsar_function_system_exceptions_total{}[1m])) > 10) by (name)
for: 1m
labels:
    severity: critical
annotations:
    summary: Pulsar high number of function errors (instance {{ $labels.instance }})
    description: |-
        Observing more than 10 Function errors per minute
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarHighNumberOfFunctionErrors.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule:

## Meaning

The `PulsarHighNumberOfFunctionErrors` alert is triggered when the rate of function errors in Pulsar exceeds 10 errors per minute. This indicates a significant issue with the Pulsar functions, potentially causing failures, data loss, or performance degradation.

## Impact

* Function execution failures may lead to data loss or inconsistencies.
* Increased error rates can cause performance degradation, leading to slower processing times or even complete system halts.
* High error rates can also indicate potential security vulnerabilities or configuration issues.

## Diagnosis

To diagnose the issue:

1. Check the Pulsar function logs for error messages and exceptions to identify the root cause of the errors.
2. Verify the function configuration and ensure that it is correct and up-to-date.
3. Check the system resources (e.g., CPU, memory, and disk space) to ensure they are not overwhelmed.
4. Review the Pulsar cluster's overall health and performance metrics to identify any underlying issues.

## Mitigation

To mitigate the issue:

1. **Immediate Action**: Pause or roll back the affected Pulsar functions to prevent further errors and data loss.
2. **Short-term Fix**: Investigate and resolve the root cause of the errors, which may involve:
	* Updating function code or configuration.
	* Adjusting system resources or scaling the Pulsar cluster.
	* Implementing retry mechanisms or circuit breakers to handle transient errors.
3. **Long-term Solution**: Implement proactive measures to prevent similar issues in the future, such as:
	* Enhancing function monitoring and logging.
	* Implementing automated testing and validation for functions.
	* Conducting regular Pulsar cluster maintenance and upgrades.
4. **Post-Incident Activities**:
	* Perform a thorough post-mortem analysis to identify areas for improvement.
	* Update the runbook to reflect new knowledge and best practices.
	* Schedule a review of the incident with the teams involved to discuss lessons learned and prevent similar incidents in the future.