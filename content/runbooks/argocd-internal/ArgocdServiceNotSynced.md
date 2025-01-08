---
title: ArgocdServiceNotSynced
description: Troubleshooting for alert ArgocdServiceNotSynced
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ArgocdServiceNotSynced

Service {{ $labels.name }} run by argo is currently not in sync.

<details>
  <summary>Alert Rule</summary>

{{% rule "argocd/argocd-internal.yml" "ArgocdServiceNotSynced" %}}

{{% comment %}}

```yaml
alert: ArgocdServiceNotSynced
expr: argocd_app_info{sync_status!="Synced"} != 0
for: 15m
labels:
    severity: warning
annotations:
    summary: ArgoCD service not synced (instance {{ $labels.instance }})
    description: |-
        Service {{ $labels.name }} run by argo is currently not in sync.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/argocd-internal/ArgocdServiceNotSynced.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `ArgocdServiceNotSynced`:

## Meaning

The `ArgocdServiceNotSynced` alert is triggered when an ArgoCD service is not in sync. This means that the service's desired state and actual state do not match, which may indicate a problem with the deployment or configuration of the service.

## Impact

If an ArgoCD service is not in sync, it may lead to:

* Inconsistent behavior of the service
* Errors or exceptions in the service
* Downtime or unavailability of the service
* Data inconsistencies or losses

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the ArgoCD logs for errors or warnings related to the service
2. Verify that the service configuration is correct and up-to-date
3. Check the service's deployment status and history
4. Review the ArgoCD application YAML file to ensure it is correctly configured
5. Verify that the service is correctly registered with ArgoCD

## Mitigation

To mitigate the issue, follow these steps:

1. Check the ArgoCD UI for any errors or warnings related to the service
2. Run the command `argocd app sync <app-name>` to sync the service
3. Verify that the service is now in sync and running correctly
4. Investigate and fix any underlying issues that may have caused the service to become out of sync
5. Consider setting up automated syncing of the service to prevent this issue from happening again in the future.