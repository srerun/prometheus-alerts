---
title: KubernetesCronjobSuspended
description: Troubleshooting for alert KubernetesCronjobSuspended
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesCronjobSuspended

CronJob {{ $labels.namespace }}/{{ $labels.cronjob }} is suspended

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesCronjobSuspended" %}}

{{% comment %}}

```yaml
alert: KubernetesCronjobSuspended
expr: kube_cronjob_spec_suspend != 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Kubernetes CronJob suspended ({{ $labels.namespace }}/{{ $labels.cronjob }})
    description: |-
        CronJob {{ $labels.namespace }}/{{ $labels.cronjob }} is suspended
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesCronjobSuspended.md

```

{{% /comment %}}

</details>


Here is a runbook for the KubernetesCronjobSuspended alert:

## Meaning

The KubernetesCronjobSuspended alert is triggered when a CronJob in a Kubernetes cluster is suspended. A suspended CronJob will not execute its scheduled tasks, which can impact the application or system that relies on it. This alert is classified as a warning, indicating a potential issue that requires attention.

## Impact

The impact of a suspended CronJob can vary depending on the specific use case and application. Some possible consequences include:

* Missed scheduled tasks, leading to data inconsistency or staleness
* Delays in processing critical workflows or jobs
* Increased latency or errors in dependent systems
* Potential security risks if the suspended CronJob is responsible for security-related tasks

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the CronJob's configuration and status using `kubectl` commands:
```
kubectl describe cronjob <cronjob-name> -n <namespace>
```
2. Review the CronJob's suspension history:
```
kubectl get cronjob <cronjob-name> -n <namespace> -o yaml
```
3. Verify if there are any errors or warnings in the CronJob's logs:
```
kubectl logs <cronjob-name> -n <namespace>
```
4. Check the Kubernetes cluster's event logs for any relevant errors or warnings:
```
kubectl get events -n <namespace>
```
5. Verify if there are any changes to the CronJob's configuration or deployment that may have caused the suspension.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the reason for the CronJob's suspension and address it accordingly:
	* If the suspension is intentional, ensure that the CronJob is correctly configured and rescheduled as needed.
	* If the suspension is unintentional, investigate and resolve the underlying issue that caused the suspension.
2. Unsuspend the CronJob using `kubectl`:
```
kubectl patch cronjob <cronjob-name> -n <namespace> -p='{"spec":{"suspend": false}}'
```
3. Verify that the CronJob is running successfully and not suspended:
```
kubectl get cronjob <cronjob-name> -n <namespace> -o yaml
```
4. Monitor the CronJob's status and logs to ensure it continues to run as expected.
5. Update the KubernetesCronjobSuspended alert rule to reflect the resolved status.