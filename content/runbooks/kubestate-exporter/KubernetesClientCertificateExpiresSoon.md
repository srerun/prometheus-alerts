---
title: KubernetesClientCertificateExpiresSoon
description: Troubleshooting for alert KubernetesClientCertificateExpiresSoon
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesClientCertificateExpiresSoon

A client certificate used to authenticate to the apiserver is expiring in less than 24.0 hours.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesClientCertificateExpiresSoon" %}}

{{% comment %}}

```yaml
alert: KubernetesClientCertificateExpiresSoon
expr: apiserver_client_certificate_expiration_seconds_count{job="apiserver"} > 0 and histogram_quantile(0.01, sum by (job, le) (rate(apiserver_client_certificate_expiration_seconds_bucket{job="apiserver"}[5m]))) < 24*60*60
for: 0m
labels:
    severity: critical
annotations:
    summary: Kubernetes client certificate expires soon (instance {{ $labels.instance }})
    description: |-
        A client certificate used to authenticate to the apiserver is expiring in less than 24.0 hours.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesClientCertificateExpiresSoon.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
