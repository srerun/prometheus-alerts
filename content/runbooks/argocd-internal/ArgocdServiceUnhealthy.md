---
title: ArgocdServiceUnhealthy
description: Troubleshooting for alert ArgocdServiceUnhealthy
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ArgocdServiceUnhealthy

Service {{ $labels.name }} run by argo is currently not healthy.

<details>
  <summary>Alert Rule</summary>

{{% rule "argocd/argocd-internal.yml" "ArgocdServiceUnhealthy" %}}

{{% comment %}}

```yaml
alert: ArgocdServiceUnhealthy
expr: argocd_app_info{health_status!="Healthy"} != 0
for: 15m
labels:
    severity: warning
annotations:
    summary: ArgoCD service unhealthy (instance {{ $labels.instance }})
    description: |-
        Service {{ $labels.name }} run by argo is currently not healthy.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/argocd-internal/ArgocdServiceUnhealthy.md

```

{{% /comment %}}

</details>


## Meaning
The `ArgocdServiceUnhealthy` alert is triggered when the `argocd_app_info` metric indicates that an Argo CD service is not in a healthy state. This metric is used to monitor the health of Argo CD services, and a non-healthy state can indicate a problem with the service, such as a deployment failure or a misconfiguration.

## Impact
If this alert is triggered, it may indicate that one or more Argo CD services are not functioning as expected, which can lead to:

* Deployment failures or delays
* Inconsistent application states
* Increased latency or errors in the application
* Reduced overall system reliability

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the Argo CD dashboard for any errors or warnings related to the affected service.
2. Verify the service configuration and deployment parameters to ensure they are correct.
3. Review the Argo CD logs to identify any errors or exceptions related to the service.
4. Check the application logs to see if there are any errors or issues related to the service.
5. Verify that the necessary dependencies and resources are available for the service.

## Mitigation
To mitigate the issue, follow these steps:

1. Investigate and resolve any configuration or deployment issues with the affected service.
2. Restart the Argo CD service or redeploy the application as needed.
3. Verify that the service is healthy and functioning as expected.
4. Update the Argo CD configuration and deployment parameters as needed to prevent future issues.
5. Consider implementing additional monitoring and alerting to detect service health issues earlier.

Note: This runbook provides general guidance for troubleshooting and mitigating the `ArgocdServiceUnhealthy` alert. For more detailed and specific instructions, refer to the Argo CD documentation and the `argocd_app_info` metric documentation.