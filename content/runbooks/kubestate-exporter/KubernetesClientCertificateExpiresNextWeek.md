---
title: KubernetesClientCertificateExpiresNextWeek
description: Troubleshooting for alert KubernetesClientCertificateExpiresNextWeek
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesClientCertificateExpiresNextWeek

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A client certificate used to authenticate to the apiserver is expiring next week.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesClientCertificateExpiresNextWeek" %}}

<!-- Rule when generated

```yaml
alert: KubernetesClientCertificateExpiresNextWeek
expr: apiserver_client_certificate_expiration_seconds_count{job="apiserver"} > 0 and histogram_quantile(0.01, sum by (job, le) (rate(apiserver_client_certificate_expiration_seconds_bucket{job="apiserver"}[5m]))) < 7*24*60*60
for: 0m
labels:
    severity: warning
annotations:
    summary: Kubernetes client certificate expires next week (instance {{ $labels.instance }})
    description: |-
        A client certificate used to authenticate to the apiserver is expiring next week.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesClientCertificateExpiresNextWeek.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
