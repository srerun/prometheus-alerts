---
title: KubernetesHpaMetricsUnavailability
description: Troubleshooting for alert KubernetesHpaMetricsUnavailability
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesHpaMetricsUnavailability

HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} is unable to collect metrics

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesHpaMetricsUnavailability" %}}

{{% comment %}}

```yaml
alert: KubernetesHpaMetricsUnavailability
expr: kube_horizontalpodautoscaler_status_condition{status="false", condition="ScalingActive"} == 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Kubernetes HPA metrics unavailability (instance {{ $labels.instance }})
    description: |-
        HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} is unable to collect metrics
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesHpaMetricsUnavailability.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `KubernetesHpaMetricsUnavailability`:

## Meaning

The `KubernetesHpaMetricsUnavailability` alert is triggered when a Kubernetes Horizontal Pod Autoscaler (HPA) is unable to collect metrics. This means that the HPA is not able to make informed scaling decisions, which can lead to unpredictable behavior and potential issues with application performance and resource utilization.

## Impact

The impact of this alert is that the affected HPA will not be able to scale the target deployment or replica set based on current metrics, which can lead to:

* Insufficient resources being allocated to the application, resulting in performance issues or errors
* Over-allocation of resources, resulting in wasted resources and increased costs
* Unstable or unpredictable application behavior

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the HPA configuration to ensure that it is correctly configured to collect metrics from the target deployment or replica set.
2. Verify that the metric collection mechanism (e.g. metrics-server, kube-state-metrics) is functioning correctly and able to provide metrics to the HPA.
3. Check the HPA logs for any errors or issues related to metric collection.
4. Verify that the HPA is able to connect to the Kubernetes API server and retrieve the necessary metrics.

## Mitigation

To mitigate the issue, follow these steps:

1. Verify that the HPA configuration is correct and that the target deployment or replica set is correctly labeled.
2. Restart the HPA controller to force a re-sync of the metrics.
3. Check the metric collection mechanism and troubleshoot any issues found.
4. If the issue persists, consider increasing the verbosity of the HPA logs to gather more information about the issue.
5. If necessary, involve the Kubernetes or application team to assist with troubleshooting and resolution.

Note: This runbook provides general guidance and may need to be tailored to the specific environment and deployment.