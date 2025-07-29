---
title: ProxmoxCertificateExpiringWarning
description: Troubleshooting for alert ProxmoxCertificateExpiringWarning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxCertificateExpiringWarning

The certificate with subject {{ printf "{{ $labels.subject }}" }} on that node is expiring in {{ printf "{{ $value }}" }} days

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxCertificateExpiringWarning" %}}

{{% comment %}}

```yaml
alert: ProxmoxCertificateExpiringWarning
expr: |
    proxmox_node_days_until_cert_expiration < {{ .Values.prometheusRule.threshold_ProxmoxCertificateExpiringWarning | default 14 }}
for: 5m
labels:
    severity: warning
annotations:
    summary: Proxmox certificate on node {{ printf "{{ $labels.node }}" }} is expiring soon
    description: The certificate with subject {{ printf "{{ $labels.subject }}" }} on that node is expiring in {{ printf "{{ $value }}" }} days
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxcertificateexpiringwarning/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "ProxmoxCertificateExpiringWarning":

## Meaning

The ProxmoxCertificateExpiringWarning alert is triggered when a Proxmox node's certificate is close to expiring. This alert is raised when the number of days until the certificate expires is less than the configured threshold (default is 14 days). This alert is meant to warn about an impending certificate expiration, which could cause connectivity issues and disruptions to Proxmox services.

## Impact

If this alert is not addressed, the Proxmox node's certificate will expire, causing:

* Disruptions to Proxmox services
* Connection issues to the node
* Potential security risks due to an invalid certificate

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Proxmox node's certificate information using the Proxmox web interface or command-line tools.
2. Verify the certificate's expiration date and subject using the `proxmox_node_days_until_cert_expiration` metric in Prometheus.
3. Check the Proxmox node's system logs for any certificate-related errors or warnings.

## Mitigation

To mitigate this issue, follow these steps:

1. Renew the Proxmox node's certificate using the Proxmox web interface or command-line tools.
2. Verify that the new certificate is properly installed and configured.
3. Test the Proxmox services to ensure they are functioning as expected.
4. Clear the alert in Prometheus once the certificate has been renewed and the issue is resolved.