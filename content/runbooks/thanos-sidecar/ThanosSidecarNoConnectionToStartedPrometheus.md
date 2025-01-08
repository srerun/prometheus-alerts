---
title: ThanosSidecarNoConnectionToStartedPrometheus
description: Troubleshooting for alert ThanosSidecarNoConnectionToStartedPrometheus
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosSidecarNoConnectionToStartedPrometheus

Thanos Sidecar {{$labels.instance}} is unhealthy.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-sidecar.yml" "ThanosSidecarNoConnectionToStartedPrometheus" %}}

{{% comment %}}

```yaml
alert: ThanosSidecarNoConnectionToStartedPrometheus
expr: thanos_sidecar_prometheus_up{job=~".*thanos-sidecar.*"} == 0 and on (namespace, pod)prometheus_tsdb_data_replay_duration_seconds != 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Sidecar No Connection To Started Prometheus (instance {{ $labels.instance }})
    description: |-
        Thanos Sidecar {{$labels.instance}} is unhealthy.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-sidecar/ThanosSidecarNoConnectionToStartedPrometheus.md

```

{{% /comment %}}

</details>


Here is a runbook for the ThanosSidecarNoConnectionToStartedPrometheus alert:

## Meaning

The ThanosSidecarNoConnectionToStartedPrometheus alert is triggered when a Thanos Sidecar instance is unable to connect to a started Prometheus instance. This alert is critical because it indicates a failure in the monitoring pipeline, which can lead to incomplete or inaccurate metrics.

## Impact

The impact of this alert is that metrics from the affected Prometheus instance will not be ingested into Thanos, resulting in incomplete or inaccurate monitoring data. This can lead to:

* Incomplete visibility into system performance and health
* Inaccurate alerting and reporting
* Delayed or missed detection of critical issues

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Sidecar logs for errors or warnings related to connecting to the Prometheus instance.
2. Verify that the Prometheus instance is running and healthy.
3. Check the network connectivity between the Thanos Sidecar and Prometheus instances.
4. Verify that the Prometheus instance is correctly configured to allow connections from the Thanos Sidecar.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Thanos Sidecar instance to attempt to re-establish the connection to the Prometheus instance.
2. Check and correct any configuration issues with the Prometheus instance or Thanos Sidecar.
3. Verify that the network connectivity between the instances is stable and working as expected.
4. If the issue persists, consider increasing the logging level of the Thanos Sidecar to gather more detailed information about the connection issue.

Additional resources:

* Refer to the Thanos Sidecar documentation for more information on configuring and troubleshooting connections to Prometheus instances.
* If you are unsure about the root cause of the issue or need further assistance, consult with your monitoring team or a Thanos expert.