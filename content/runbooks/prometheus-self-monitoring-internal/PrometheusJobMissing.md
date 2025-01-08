---
title: PrometheusJobMissing
description: Troubleshooting for alert PrometheusJobMissing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusJobMissing

A Prometheus job has disappeared

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusJobMissing" %}}

{{% comment %}}

```yaml
alert: PrometheusJobMissing
expr: absent(up{job="prometheus"})
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus job missing (instance {{ $labels.instance }})
    description: |-
        A Prometheus job has disappeared
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusJobMissing.md

```

{{% /comment %}}

</details>


Here is a runbook for the PrometheusJobMissing alert rule:

## Meaning
The PrometheusJobMissing alert is triggered when Prometheus is unable to scrape metrics from its own internal job. This alert is critical because Prometheus is responsible for monitoring and alerting on other services, and if it is not functioning correctly, it may prevent detecting issues in other parts of the system.

## Impact
The impact of this alert is that Prometheus will not be able to collect metrics from its own internal job, which may lead to:

* Incomplete or missing metrics
* Delayed or missed alerts for critical issues in the system
* Increased mean time to detect (MTTD) and mean time to resolve (MTTR) for incidents

## Diagnosis
To diagnose the issue, follow these steps:

* Check the Prometheus server logs for any errors or warnings related to the job
* Verify that the Prometheus instance is running and healthy
* Check the Prometheus configuration to ensure that the job is correctly configured
* Verify that the instance label in the alert matches the expected instance name
* Check the Prometheus targets page to ensure that the job is not being scraped correctly

## Mitigation
To mitigate the issue, follow these steps:

* Restart the Prometheus instance to ensure it is running correctly
* Check the Prometheus configuration and update it if necessary to ensure the job is correctly configured
* Verify that the instance label in the alert matches the expected instance name
* Check the Prometheus targets page and update the targets configuration if necessary
* If the issue persists, escalate to the Prometheus administrator for further assistance