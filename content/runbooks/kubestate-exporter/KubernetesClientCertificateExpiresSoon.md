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

The KubernetesClientCertificateExpiresSoon alert is triggered when a client certificate used to authenticate to the Kubernetes API server (apiserver) is approaching its expiration date. This certificate is used by clients, such as kubectl, to authenticate with the API server.

## Impact

If a client certificate expires, it can cause disruptions to the cluster's functionality, leading to:

* Loss of access to the cluster for clients using the expiring certificate
* Disruption of automation tools and scripts that rely on the certificate for authentication
* Potential security risks if the expired certificate is not renewed or replaced in a timely manner

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the alert notifications for the instance(s) affected by the expiring certificate.
2. Verify the certificate expiration date and time using the `kubectl` command: `kubectl get csr/<csr-name> -o yaml`
3. Check the Kubernetes API server logs for any authentication errors related to the expiring certificate.

## Mitigation

To mitigate the issue, follow these steps:

1. Renew or replace the expiring client certificate as soon as possible.
2. Update the certificate on all clients, such as kubectl, that use it for authentication.
3. Verify that the new certificate is properly configured and functional.
4. Restart any automation tools or scripts that rely on the certificate for authentication.

Additionally, consider implementing a certificate rotation process to ensure that certificates are renewed or replaced regularly, avoiding last-minute expirations.