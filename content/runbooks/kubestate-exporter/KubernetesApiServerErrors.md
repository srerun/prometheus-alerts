---
title: KubernetesApiServerErrors
description: Troubleshooting for alert KubernetesApiServerErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesApiServerErrors

Kubernetes API server is experiencing high error rate

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesApiServerErrors" %}}

{{% comment %}}

```yaml
alert: KubernetesApiServerErrors
expr: sum(rate(apiserver_request_total{job="apiserver",code=~"(?:5..)"}[1m])) by (instance, job) / sum(rate(apiserver_request_total{job="apiserver"}[1m])) by (instance, job) * 100 > 3
for: 2m
labels:
    severity: critical
annotations:
    summary: Kubernetes API server errors (instance {{ $labels.instance }})
    description: |-
        Kubernetes API server is experiencing high error rate
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesApiServerErrors.md

```

{{% /comment %}}

</details>


## Meaning

The `KubernetesApiServerErrors` alert is triggered when the rate of errors in the Kubernetes API server exceeds 3% of the total requests over a 1-minute period. This alert is critical and indicates that the API server is experiencing a high error rate, which can impact the overall reliability and performance of the Kubernetes cluster.

## Impact

The impact of this alert can be significant, as it may indicate:

* Increased latency or timeouts for API requests
* Failure to deploy or manage resources in the cluster
* Increased error rates for applications and services running in the cluster
* Potential data loss or corruption due to failed API requests

## Diagnosis

To diagnose the root cause of the alert, follow these steps:

1. Check the API server logs for errors and exceptions
2. Investigate the cụause of the errors (e.g., network issues, configuration problems, etc.)
3. Verify that the API server is running and healthy
4. Check the cluster's resource utilization (CPU, memory, disk) to ensure it's not overwhelmed
5. Review the Kubernetes cluster configuration to ensure it's correctly set up

## Mitigation

To mitigate the impact of the alert, take the following steps:

1. Investigate and resolve the underlying cause of the errors (e.g., fix network issues, update configurations, etc.)
2. Restart the API server if necessary
3. Scale up the API server to handle increased load (if necessary)
4. Implement retry mechanisms for failed API requests
5. Monitor the API server performance and adjust resource allocations as needed

Remember to refer to the [runbook](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesApiServerErrors.md) for more detailed steps and guidelines specific to your environment.